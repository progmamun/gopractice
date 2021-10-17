package main

import (
	"log"
	"reflect"
)

type A struct {
	B
	x int
	Y int
	Z int
}

type B struct {
	F string
	G string
}

func main() {
	var a A
	v := reflect.ValueOf(&a)
	func() {
		// trying to get fields from ptr panics
		defer func() {
			log.Println("panic:", recover())
		}()
		log.Printf("%s", v.Field(1).String())
	}()
	v = v.Elem()
	// changing fields by index
	for i := 0; i < 4; i++ {
		f := v.Field(i)
		if f.CanSet() && f.Type().Kind() == reflect.Int {
			f.SetInt(42)
		}
	}
	// changing nested fields by index
	v.FieldByIndex([]int{0, 1}).SetString("banana")

	// getting fields by name
	v.FieldByName("B").FieldByName("F").SetString("apple")

	log.Printf("%+v", a)
}
