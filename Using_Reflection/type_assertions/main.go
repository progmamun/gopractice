package main

import (
	"fmt"
)

/*
# unsafe assertion
v := SomeVar.(SomeType)

# safe assertion
v, ok := SomeVar.(SomeType)

func main() {
var a interface{} = "hello"
fmt.Println(a.(string)) // ok
fmt.Println(a.(int)) // panics!
}
*/
func main() {
	var a interface{} = "hello"
	s, ok := a.(string) // true
	fmt.Println(s, ok)
	i, ok := a.(int) // false
	fmt.Println(i, ok)
}
