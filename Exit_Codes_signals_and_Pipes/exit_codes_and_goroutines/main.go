package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("go end (deferred)")
		fmt.Println("go start")
		os.Exit(1)
	}()
	defer fmt.Println("main end (deferred)")
	fmt.Println("main start")
	time.Sleep(time.Second)
	fmt.Println("main end")
}
