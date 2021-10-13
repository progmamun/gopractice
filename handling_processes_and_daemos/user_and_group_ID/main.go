package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("User ID:", os.Getuid())  // 1
	fmt.Println("Group ID:", os.Getgid()) // 1
	groups, err := os.Getgroups()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Group IDs:", groups) // [1]
}
