package main

import (
	"fmt"
	"sync/atomic"
)

type genInt64 int64

func (g *genInt64) Next() int64 {
	return atomic.AddInt64((*int64)(g), 1)
}

type genInt64 struct {
	ch chan int64
}

func (g genInt64) Next() int64 {
	return <-g.ch
}
func NewGenInt64() genInt64 {
	g := genInt64{ch: make(chan int64)}
	go func() {
		for i := int64(0); ; i++ {
			g.ch <- i
		}
	}()
	return g
}

func GenInt64() <-chan int64 {
	ch := make(chan int64)
	go func() {
		for i := int64(0); ; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	ch1, ch2 := GenInt64(), GenInt64()
	for i := 0; i < 20; i++ {
		select {
		case v := <-ch1:
			fmt.Println("ch 1", v)
		case v := <-ch2:
			fmt.Println("ch 2", v)
		}
	}
}
