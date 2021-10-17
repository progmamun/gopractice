package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = int64(12)
	v := reflect.ValueOf(&a)
	fmt.Println(v.String(), v.CanSet())
	e := v.Elem()
	fmt.Println(e.String(), e.CanSet())
	e.SetInt(24)
	fmt.Println(a)
}

/*
func main() {
	var a = int64(12)
	v := reflect.ValueOf(a)
	fmt.Println(v.String(), v.CanSet())
	v.SetInt(24)
}
*/
