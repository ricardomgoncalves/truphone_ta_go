package requestid

import (
	"context"
)

type key int

var requestIdKey key

func WithRequestId(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIdKey, id)
}

func GetRequestId(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(requestIdKey).(string)
	return val, ok
}
