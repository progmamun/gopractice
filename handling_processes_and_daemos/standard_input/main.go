package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Current PID:", os.Getpid())         //Current PID: 3
	fmt.Println("Current Parent PID:", os.Getppid()) // Current Parent PID: 2
}
