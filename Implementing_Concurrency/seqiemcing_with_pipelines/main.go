package main

import "context"

func stage(in <-chan interface{}) <-chan interface{} {
	var out = make(chan interface{})
	go func() {
		for v := range in {
			v = v.(int) + 1 // some operation
			out <- v
		}
		close(out)
	}()
	return out
}

func stage(ctx context.Context, in <-chan interface{}) <-chan interface{} {
	var out = make(chan interface{})
	go func() {
		defer close(out)
		for v := range in {
			v = v.(int) + 1 // some operation
			select {
			case out <- v:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func SourceLine(ctx context.Context, r io.ReadCloser) <-chan string {
	ch := make(chan string)
	go func() {
	defer func() { r.Close(); close(ch) }()
	s := bufio.NewScanner(r)
	for s.Scan() {
	select {
	case <-ctx.Done():
	return
	case ch <- s.Text():
	}
	}
	}()
	return ch
	}

	func TextFilter(ctx context.Context, src <-chan string, filter string) <-chan string {
		ch := make(chan string)
		go func() {
		defer close(ch)
		for v := range src {
		if !strings.Contains(v, filter) {
		continue
		}
		select {
		case <-ctx.Done():
		return
		case ch <- v:
		}
		}
		}()
		return ch
		}

		func Printer(ctx context.Context, src <-chan string, color int, highlight string, w io.Writer) {
			const close = "\x1b[39m"
			open := fmt.Sprintf("\x1b[%dm", color)
			for {
			select {
			case <-ctx.Done():
			return
			case v, ok := <-src:
			if !ok {
			return
			}
			i := strings.Index(v, highlight)
			if i == -1 {
			panic(v)
			}
			fmt.Fprint(w, v[:i], open, highlight, close, v[i+len(highlight):], "\n")
			}
			}
			}
			func main() {
				var search string
				...
				ctx := context.Background()
				src := SourceLine(ctx, ioutil.NopCloser(strings.NewReader(sometext)))
				filter := TextFilter(ctx, src, search)
				Printer(ctx, filter, 31, search, os.Stdout)
				}
							
