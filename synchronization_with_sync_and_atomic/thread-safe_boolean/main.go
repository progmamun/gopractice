package main

import "sync/atomic"

type cond int32

func (c *cond) Set(v bool) {
	a := int32(0)
	if v {
		a++
	}
	atomic.StoreInt32((*int32)(c), a)
}
func (c *cond) Value() bool {
	return atomic.LoadInt32((*int32)(c)) != 0
}
