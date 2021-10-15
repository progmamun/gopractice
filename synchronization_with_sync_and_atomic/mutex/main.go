package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	done := make(chan struct{}, 10)
	for i := 0; i < cap(done); i++ {
		go func(i int, l sync.Locker) {
			l.Lock()
			defer l.Unlock()
			fmt.Println(i)
			time.Sleep(time.Millisecond * 10)
			done <- struct{}{}
		}(i, &m)
	}
	for i := 0; i < cap(done); i++ {
		<-done
	}
}
