package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(int64(100))
	// zero value
	fmt.Printf("%#v\n", reflect.Zero(t))
	// pointer to int
	fmt.Printf("%#v\n", reflect.New(t))
}
