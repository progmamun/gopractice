package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = int64(23)
	fmt.Println(reflect.TypeOf(a).String())
	// int64
	fmt.Println(reflect.ValueOf(a).String())
	// <int64 Value>
	fmt.Println(reflect.ValueOf(a).Type().String())
	// int64
}
