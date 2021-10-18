package main

import (
	"fmt"
	"reflect"
)

func main() {
	reflect.ChanOf(reflect.BothDir, reflect.TypeOf(""))
	v := reflect.MakeChan(t, 0)
	fmt.Printf("%T\n", v.Interface())
}
