package main

import (
	"bytes"
	"sync"
)

type A struct{}
type Pool chan *A

func (p Pool) Get() *A {
	select {
	case a := <-p:
		return a
	default:
		return new(A)
	}
}
func (p Pool) Put(a *A) {
	select {
	case p <- a:
	default:
	}
}

type Pool struct {
	// New optionally specifies a function to generate
	// a value when Get would otherwise return nil.
	// It may not be changed concurrently with calls to Get.
	New func() interface{}
	// contains filtered or unexported fields
}

var pool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 4096))
	},
}

func Get() *bytes.Buffer {
	return pool.Get().(*bytes.Buffer)
}
func Put(b *bytes.Buffer) {
	b.Reset()
	pool.Put(b)
}

start := time.Now()
wg := sync.WaitGroup{}
wg.Add(20)
for i := 0; i < 20; i++ {
go func(v int) {
time.Sleep(time.Second * time.Duration(1+v/4))
b := Get()
defer func() {
Put(b)
wg.Done()
}()
fmt.Fprintf(b, "Goroutine %2d using %p, after %.0fs\n", v, b, time.Since(start).Seconds())
fmt.Printf("%s", b.Bytes())
}(i)
}
wg.Wait()