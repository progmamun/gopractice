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

/*
func main() {
	m := map[string]int64{}
	v := reflect.ValueOf(m)

	// setting one field
	v.SetMapIndex(reflect.ValueOf("key"), reflect.ValueOf(int64(1000)))

	fmt.Println(m)
}


func main() {
	m := map[string]int64{"a": 10}
	fmt.Println(m, len(m))

	v := reflect.ValueOf(m)

	// deleting field
	v.SetMapIndex(reflect.ValueOf("a"), reflect.Value{})

	fmt.Println(m, len(m))
}

*/
