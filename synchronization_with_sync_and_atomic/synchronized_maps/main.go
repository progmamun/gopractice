package main

import (
	"fmt"
	"sync"
)

func main() {
	var m = map[int]int{}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			m[i%5]++
			fmt.Println(m)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/*
func main() {
var m = map[int]int{}
var done = make(chan struct{})
go func() {
for i := 0; i < 100; i++ {
time.Sleep(time.Nanosecond)
m[i]++
}
close(done)
}()
for {
time.Sleep(time.Nanosecond)
fmt.Println(len(m), m)
select {
case <-done:
return
default:
}
}
}
*/
