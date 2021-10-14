package main

import (
	"fmt"
	"time"
)

// A timer exposes a read-only channel, so it's not possible to close it.
func main() {
	ch1, ch2 := make(chan int), make(chan int)
	a, b := 2, 10
	go func() { <-ch1 }()
	go func() { ch2 <- a }()
	t := time.NewTimer(time.Nanosecond)
	select {
	case ch1 <- b:
		fmt.Println("ch1 got a", b)
	case v := <-ch2:
		fmt.Println("ch2 got a", v)
	case <-t.C:
		fmt.Println("too slow")
	}
}

/*func main() {
	t := time.NewTimer(time.Millisecond)
	time.Sleep(time.Millisecond / 2)
	if !t.Stop() {
		panic("it should not fire")
	}
	select {
	case <-t.C:
		panic("not fired")
	default:
		fmt.Println("not fired")
	}
	if t.Reset(time.Millisecond) {
		panic("timer should not be active")
	}
	time.Sleep(time.Millisecond)
	if t.Stop() {
		panic("it should fire")
	}
	select {
	case <-t.C:
		fmt.Println("fired")
	default:
		panic("not fired")
	}
}
*/
