package main

import "fmt"

func main() {
	// one producer
	var ch = make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
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
