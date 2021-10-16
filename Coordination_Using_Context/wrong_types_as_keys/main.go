package main

import (
	"context"
	"fmt"
)

func main() {
	var a interface{} = "request-id"
	var b interface{} = "request-id"
	fmt.Println(a == b)
	ctx := context.Background()
	ctx = context.WithValue(ctx, a, "a")
	ctx = context.WithValue(ctx, b, "b")
	fmt.Println(ctx.Value(a), ctx.Value(b))
}
