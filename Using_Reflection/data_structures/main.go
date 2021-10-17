package main

import (
	"fmt"
	"reflect"
)

func main() {
	type X struct {
		A, B int
		c    string
	}
	var a = X{10, 100, "apple"}
	fmt.Println(a)
	e := reflect.ValueOf(&a).Elem()
	fmt.Println(e.String(), e.CanSet())
	e.Set(reflect.ValueOf(X{1, 2, "banana"}))
	fmt.Println(a)
}
