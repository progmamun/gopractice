package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "A", "sample", "command")
	fmt.Println(cmd.Path, cmd.Args[1:])
}
