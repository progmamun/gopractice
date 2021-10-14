package main

import (
	"fmt"
)

func main() {
	var (
		a = make(chan int, 5)
	)
	for i := 0; i < 5; i++ {
		a <- i
		fmt.Println("a is", len(a), "/", cap(a))
	}

}

/*func main() {
	var (
		a = make(chan int, 0)
		b = make(chan int, 5)
	)

	fmt.Println("a is", cap(a))
	fmt.Println("b is", cap(b))
}
*/
