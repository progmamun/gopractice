package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.ChanOf(reflect.BothDir, reflect.TypeOf(""))
	v := reflect.MakeChan(t, 0)
	go func() {
		for i := 0; i < 10; i++ {
			v.Send(reflect.ValueOf(fmt.Sprintf("msg-%d", i)))
		}
		v.Close()
	}()
	for msg, ok := v.Recv(); ok; msg, ok = v.Recv() {
		fmt.Println(msg)
	}
}
