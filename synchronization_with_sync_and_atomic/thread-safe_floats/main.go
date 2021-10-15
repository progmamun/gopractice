package main

import (
	"math"
	"sync/atomic"
)

type f64 uint64

func uf(u uint64) (f float64) { return math.Float64frombits(u) }
func fu(f float64) (u uint64) { return math.Float64bits(f) }
func newF64(f float64) *f64 {
	v := f64(fu(f))
	return &v
}
func (f *f64) Load() float64 {
	return uf(atomic.LoadUint64((*uint64)(f)))
}
func (f *f64) Store(s float64) {
	atomic.StoreUint64((*uint64)(f), fu(s))
}
