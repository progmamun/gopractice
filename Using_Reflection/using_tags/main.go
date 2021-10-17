package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string `json:"name,omitempty" xml:"-"`
	Surname string `json:"surname,omitempty" xml:"-"`
}

func main() {
	v := reflect.ValueOf(Person{"Micheal", "Scott"})
	t := v.Type()
	fmt.Println("Type:", t)
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("%v: %v\n", t.Field(i).Name, v.Field(i))
	}
}

/*
type A struct {
	Name    string `json:"name,omitempty" xml:"-"`
	Surname string `json:"surname,omitempty" xml:"-"`
}

func main() {
	t := reflect.TypeOf(A{})
	fmt.Println(t)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s JSON=%s  XML=%s\n", f.Name, f.Tag.Get("json"), f.Tag.Get("xml"))
	}
}

*/
