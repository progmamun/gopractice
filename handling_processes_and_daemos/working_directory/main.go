package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Working Directory:", wd)
	fmt.Println("Application:", filepath.Join(wd, os.Args[0]))
	d := filepath.Join(wd, "test")
	if err := os.Mkdir(d, 0755); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Created", d)
	if err := os.Chdir(d); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Working Directory:", d)

}
