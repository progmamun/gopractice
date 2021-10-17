package main

import (
	"bufio"
	"bytes"
	"context"
	"io"
)

func SourceLineWords(ctx context.Context, r io.ReadCloser) <-chan []string {
	ch := make(chan []string)
	go func() {
		defer func() { r.Close(); close(ch) }()
		b := bytes.Buffer{}
		s := bufio.NewScanner(r)
		for s.Scan() {
			b.Reset()
			b.Write(s.Bytes())
			words := []string{}
			w := bufio.NewScanner(&b)
			w.Split(bufio.ScanWords)
			for w.Scan() {
				words = append(words, w.Text())
			}
			select {
			case <-ctx.Done():
				return
			case ch <- words:
			}
		}
	}()
	return ch
}

func WordOccurrence(ctx context.Context, src <-chan []string) <-chan map[string]int {
	ch := make(chan map[string]int)
	go func() {
		defer close(ch)
		for v := range src {
			count := make(map[string]int)
			for _, s := range v {
				count[s]++
			}
			select {
			case <-ctx.Done():
				return
			case ch <- count:
			}
		}
	}()
	return ch
}

/*ctx, canc := context.WithCancel(context.Background())
defer canc()
src := SourceLineWords(ctx,
ioutil.NopCloser(strings.NewReader(cantoUno)))
count1, count2 := WordOccurrence(ctx, src), WordOccurrence(ctx, src)
*/
