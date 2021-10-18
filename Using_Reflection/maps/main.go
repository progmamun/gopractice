package main

import (
	"fmt"
	"reflect"
)

func main() {
	maps := []interface{}{
		make(map[string]struct{}),
		make(map[int]rune),
		make(map[float64][]byte),
		make(map[int32]chan bool),
		make(map[[2]string]interface{}),
	}
	for _, m := range maps {
		t := reflect.TypeOf(m)
		fmt.Printf("%s - k:%-10s v:%-10s\n", m, t.Key(), t.Elem())
	}
}
