package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Hello, playground")
	os.Exit(0)
}
