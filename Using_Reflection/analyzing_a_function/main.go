package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func Foo() {}

func Bar(a int, b string) {}

func Baz(a int, b string) (int, error) { return 0, nil }

func Qux(a int, b ...string) (int, error) { return 0, nil }

func main() {
	for _, f := range []interface{}{Foo, Bar, Baz, Qux} {
		t := reflect.TypeOf(f)
		name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		in := make([]reflect.Type, t.NumIn())
		for i := range in {
			in[i] = t.In(i)
		}
		out := make([]reflect.Type, t.NumOut())
		for i := range out {
			out[i] = t.Out(i)
		}
		fmt.Printf("%q %v %v %v\n", name, in, out, t.IsVariadic())
	}
}
