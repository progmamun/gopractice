package main

import (
	"fmt"
	"sync"
)

func main() {
	// three producer
	var ch = make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(n int) {
			for i := 0; i < 100; i++ {
				ch <- fmt.Sprintln(n, i)
			}
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	// one consumer
	var done = make(chan struct{})
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
		close(done)
	}()
	<-done
}
