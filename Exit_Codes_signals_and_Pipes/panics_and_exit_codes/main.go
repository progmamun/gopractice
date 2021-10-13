package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Hello, playground")
	panic("panic")
}
