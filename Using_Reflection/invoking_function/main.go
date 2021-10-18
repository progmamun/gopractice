package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func Foo() {}

func Bar(a int, b string) {}

func Baz(a int, b string) (int, error) { return a * 2, nil }

func Qux(a int, b ...string) (int, error) { return a / len(b), nil }

func main() {
	for _, f := range []interface{}{Foo, Bar, Baz, Qux} {
		v, t := reflect.ValueOf(f), reflect.TypeOf(f)
		name := runtime.FuncForPC(v.Pointer()).Name()
		in := make([]reflect.Value, t.NumIn())
		for i := range in {
			switch a := t.In(i); a.Kind() {
			case reflect.Int:
				in[i] = reflect.ValueOf(42)
			case reflect.String:
				in[i] = reflect.ValueOf("42")
			case reflect.Slice:
				switch a.Elem().Kind() {
				case reflect.Int:
					in[i] = reflect.ValueOf(21)
				case reflect.String:
					in[i] = reflect.ValueOf("21")
				}
			}
		}
		out := v.Call(in)
		fmt.Printf("%q %v%v\n", name, in, out)
	}
}
