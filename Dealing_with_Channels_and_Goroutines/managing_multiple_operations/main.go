package main

import (
	"fmt"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	a, b := 2, 10
	go func() { <-ch1 }()
	go func() { ch2 <- a }()
	select {
	case ch1 <- b:
		fmt.Println("ch1 got a", b)
	case v := <-ch2:
		fmt.Println("ch2 got a", v)
	}
}
