package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(make(chan string, 1))
	fmt.Println("sending", v.TrySend(reflect.ValueOf("message"))) // true 1 1
	branches := []reflect.SelectCase{
		{Dir: reflect.SelectRecv, Chan: v, Send: reflect.Value{}},
		{Dir: reflect.SelectSend, Chan: v, Send: reflect.ValueOf("send")},
		{Dir: reflect.SelectDefault},
	}

	// send, receive and default
	i, recv, closed := reflect.Select(branches)
	fmt.Println("select", i, recv, closed)

	v.Close()
	// just default and receive
	i, _, closed = reflect.Select(branches[:2])
	fmt.Println("select", i, closed) // 1 false
}
