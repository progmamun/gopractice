package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	a, b := 2, 10
	for i := 0; i < 10; i++ {
		go func() { <-ch1 }()
		go func() { ch2 <- a }()
		time.Sleep(time.Nanosecond)
		select {
		case ch1 <- b:
			fmt.Println("ch1 got a", b)
		case v := <-ch2:
			fmt.Println("ch2 got a", v)
		default:
			fmt.Println("too slow")
		}
	}
}
