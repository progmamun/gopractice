package main

import (
	"fmt"
	"sync"
)

func main() {
	// three consumers
	wg := sync.WaitGroup{}
	wg.Add(3)
	var ch = make(chan string)

	for i := 0; i < 3; i++ {
		go func(n int) {
			for i := range ch {
				fmt.Println(n, i)
			}
			wg.Done()
		}(i)
	}

	// one producer
	go func() {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintln("prod-", i)
		}
		close(ch)
	}()

	wg.Wait()
}
