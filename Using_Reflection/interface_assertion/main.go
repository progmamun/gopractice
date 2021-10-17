package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
)

/*
type Fooer interface {
Foo()
}
type Barer interface {
Bar()
}

*/

/*type A int
func (A) Foo() {}
type B int
func (B) Bar() {}
func (B) Foo() {}
*/
func Closer(r io.Reader) io.ReadCloser {
	if rc, ok := r.(io.ReadCloser); ok {
		return rc
	}
	return ioutil.NopCloser(r)
}

func main() {
	log.Printf("%T", Closer(&bytes.Buffer{}))
	log.Printf("%T", Closer(&os.File{}))
}

/*
type Fooer interface {
	Foo()
}

type Barer interface {
	Bar()
}

type A int

func (A) Foo() {}

type B int

func (B) Bar() {}
func (B) Foo() {}

func main() {
	var a Fooer

	a = A(0)
	v, ok := a.(Barer)
	fmt.Println(v, ok)

	a = B(0)
	v, ok = a.(Barer)
	fmt.Println(v, ok)
}

*/
