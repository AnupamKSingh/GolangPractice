package log

import (
	"fmt"
	"math/rand"
	"net/http"
	"context"
)

const requestIdKey = 42

func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIdKey).(int64)
	if !ok {
		fmt.Println("Not able to set the context Value")
		return
	}
	fmt.Println(id, msg)
}

func Decorator(f http.HandlerFunc) (http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIdKey, id)
		f(w, r.WithContext(ctx))
	}
}
