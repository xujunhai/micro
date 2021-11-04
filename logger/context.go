package logger

import "context"

type Key struct{}

func FromContext(ctx context.Context) (Logger, bool) {
	l, ok := ctx.Value(Key{}).(Logger)
	return l, ok
}

func NewContext(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, Key{}, l)
}
