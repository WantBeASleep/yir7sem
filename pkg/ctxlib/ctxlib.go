package ctxlib

import "context"

type ctxData map[string]any

type cxtKey int

const (
	public cxtKey = iota
	private
)

func PublicSet(ctx context.Context, key string, value any) context.Context {
	data, ok := ctx.Value(public).(ctxData)
	if !ok {
		data = ctxData{}
	}

	data[key] = value
	return context.WithValue(ctx, public, data)
}

func PublicGet(ctx context.Context, key string) any {
	data, ok := ctx.Value(public).(ctxData)
	if !ok {
		return nil
	}

	return data[key]
}

func PublicGetAll(ctx context.Context) map[string]any {
	data, ok := ctx.Value(public).(ctxData)
	if !ok {
		return nil
	}

	return data
}

func PrivateSet(ctx context.Context, key string, value any) context.Context {
	data, ok := ctx.Value(private).(ctxData)
	if !ok {
		data = ctxData{}
	}

	data[key] = value
	return context.WithValue(ctx, private, data)
}

func PrivateGet(ctx context.Context, key string) any {
	data, ok := ctx.Value(private).(ctxData)
	if !ok {
		return nil
	}

	return data[key]
}
