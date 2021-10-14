package main

import (
	"fmt"
)

func main() {
	const max = 10
	var a = make(chan int)

	go func() {
		for i := 0; i < max; i++ {
			a <- i
		}
	}()
	for i := 0; i < max; i++ {
		fmt.Println(<-a)
	}
}

/*func main() {
	var (
		a = make(chan int, 5)
	)
	for i := 0; i < 5; i++ {
		a <- i
		fmt.Println("a is", len(a), "/", cap(a))
	}
	a <- 0
}
*/
