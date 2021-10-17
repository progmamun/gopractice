package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = int(12)
	v := reflect.ValueOf(a)
	fmt.Println(v.String())
	fmt.Printf("%v", v.Interface())
}
