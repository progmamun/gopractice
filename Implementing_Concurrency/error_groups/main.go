package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	ctx, canc := context.WithTimeout(context.Background(), time.Second)
	defer canc()
	wg := sync.WaitGroup{}
	wg.Add(10)
	var ch = make(chan int)
	for i := 0; i < 10; i++ {
		go func(ctx context.Context, i int) {
			defer wg.Done()
			d := time.Duration(rand.Intn(2000)) * time.Millisecond
			time.Sleep(d)
			select {
			case <-ctx.Done():
				fmt.Println(i, "early exit after", d)
				return
			case ch <- i:
				fmt.Println(i, "normal exit after", d)
			}
		}(ctx, i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for range ch {
	}
}

func visitor(url string) func() error {
	return func() (err error) {
		s := time.Now()
		defer func() {
			log.Println(url, time.Since(s), err)
		}()
		var resp *http.Response
		if resp, err = http.Get(url); err != nil {
			return
		}
		return resp.Body.Close()
	}
}
func main() {
	eg := errgroup.Group{}
	var urlList = []string{
		"http://www.golang.org/",
		"http://invalidwebsite.hey/",
		"http://www.google.com/",
	}
	for _, url := range urlList {
		eg.Go(visitor(url))
	}
	if err := eg.Wait(); err != nil {
		log.Fatalln("Error:", err)
	}
}
func sender(ctx context.Context, ch chan<- string, n int) func() error {
	return func() (err error) {
		for i := 0; ; i++ {
			if rand.Intn(100) == 42 {
				return errors.New("the answer")
			}
			select {
			case ch <- fmt.Sprintf("[%d]%d", n, i):
			case <-ctx.Done():
				return nil
			}
		}
	}
}

/*func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		eg.Go(sender(ctx, ch, i))
	}
	go func() {
		for s := range ch {
			log.Println(s)
		}
	}()
	if err := eg.Wait(); err != nil {
		log.Println("Error:", err)
	}
	close(ch)
	log.Println("waiting...")
	time.Sleep(time.Second)
}
*/
