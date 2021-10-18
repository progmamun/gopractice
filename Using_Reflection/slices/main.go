package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := []int{10, 20, 100}
	v := reflect.ValueOf(m)

	for i := 0; i < v.Len(); i++ {
		fmt.Println(i, v.Index(i))
	}
}

/*
func main() {
	m := []int64{10, 20, 100}
	v := reflect.ValueOf(m)

	for i := 0; i < v.Len(); i++ {
		v.Index(i).SetInt(v.Index(i).Interface().(int64) * 2)
	}
	fmt.Println(m)
}

func main() {
	var s = []int{1, 2}
	fmt.Println(s)

	v := reflect.ValueOf(s)
	// same as append(s, 3)
	v2 := reflect.Append(v, reflect.ValueOf(3))
	// s can't and does not change
	fmt.Println(v.CanSet(), v, v2)

	// using the pointer allows change
	v = reflect.ValueOf(&s).Elem()
	v.Set(v2)
	fmt.Println(v.CanSet(), v, v2)
}
*/
