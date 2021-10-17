package main

const (
	N = 3
	M = 5
	)
	wg1 := sync.WaitGroup{}
	wg1.Add(N)
	wg2 := sync.WaitGroup{}
	wg2.Add(M)
	var ch = make(chan string)


		for i := 0; i < N; i++ {
			go func(n int) {
			for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("src-%d[%d]", n, i)
			}
			wg1.Done()
			}(i)
			}
			for i := 0; i < M; i++ {
			go func(n int) {
			for i := range ch {
			fmt.Printf("cons-%d, msg %q\n", n, i)
			}
			wg2.Done()
			}(i)
			}

/*	wg1.Wait()
close(ch)
wg2.Wait() */