package appctx

import (
	"context"
	"strconv"
)

type key int

func (k key) String() string {
	return strconv.Itoa(int(k))
}

const (
	MetaContextKey key = iota + 1
	DataContextKey
	ErrorContextKey
	TransactionContextKey
)

// SetValue wrapped the context.WithValue with appctx keys
func SetValue(ctx context.Context, key key, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

// GetValue from app context with key
func GetValue(ctx context.Context, key key) interface{} {
	return ctx.Value(key)
}
