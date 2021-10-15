package main

import (
	"context"
	"fmt"
)

type keyType struct{}

var key = &keyType{}

func WithKey(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

func GetKey(ctx context.Context) (string, bool) {
	v := ctx.Value(key)
	if v == nil {
		return "", false
	}
	return v.(string), true
}

func main() {
	ctx := WithKey(context.Background(), "123")
	fmt.Println(GetKey(ctx))
}
