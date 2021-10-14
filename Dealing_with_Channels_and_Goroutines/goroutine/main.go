package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello, playground")
	}()
	time.Sleep(time.Nanosecond)
}

/* func main() {
go fmt.Println("Hello, playground")
time.Sleep(time.Nanosecond)
}
*/

/*type a struct{}
func (a) Method() { fmt.Println("Hello, playground") }
func main() {
go a{}.Method()
time.Sleep(time.Nanosecond)
}*/
