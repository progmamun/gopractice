package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{}

	a = "" // bulltin string
	t := reflect.TypeOf(a)
	fmt.Println(t.String(), t.Kind())

	type A string // custom type
	a = A("")
	t = reflect.TypeOf(a)
	fmt.Println(t.String(), t.Kind())
}

/*
func main() {
	var a interface{}

	a = new(int) // int pointer
	t := reflect.TypeOf(a)
	fmt.Println(t.String(), t.Kind())

	a = new(struct{}) // struct pointer
	t = reflect.TypeOf(a)
	fmt.Println(t.String(), t.Kind())
}

*/
