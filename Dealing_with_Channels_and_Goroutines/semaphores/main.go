package main

import (
	"fmt"
	"time"
)

type sem chan struct{}

func (s sem) Acquire() {
	s <- struct{}{}
}

func (s sem) Relase() {
	<-s
}

func main() {
	s := make(sem, 5)
	for i := 0; i < 10; i++ {
		go func(i int) {
			s.Acquire()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			s.Relase()
		}(i)
	}
	time.Sleep(time.Second * 3)
}
