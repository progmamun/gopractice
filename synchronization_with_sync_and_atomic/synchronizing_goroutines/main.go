package main

import (
	"fmt"
	"sync"
)

type WaitGroup struct {
	noCopy noCopy
	state1 [3]uint32
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func(a int) {
			for i := 1; i <= 10; i++ {
				fmt.Printf("%dx%d=%d\n", a, i, a*i)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/*func main() {
wg := sync.WaitGroup{}
for i := 1; rand.Intn(10) != 0; i++ {
wg.Add(1)
go func(a int) {
for i := 1; i <= 10; i++ {
fmt.Printf("%dx%d=%d\n", a, i, a*i)
}
wg.Done()
}(i)
}
wg.Wait()
}

func main() {
wg := sync.WaitGroup{}
for i := 1; i < 10; i++ {
go func(a int) {
wg.Add(1)
for i := 1; i <= 10; i++ {
fmt.Printf("%dx%d=%d\n", a, i, a*i)
}
wg.Done()
}(i)
}
wg.Wait()
}
*/
