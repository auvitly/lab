package assistant

import (
	"context"
)

// SetValue - returns counter by key from Assistant.
func SetValue(ctx context.Context, key string, value any) bool {
	a, err := fromContext(ctx)
	if err != nil {
		return false
	}

	a.data[key] = value

	return true
}

// GetValue - returns the value by key from Assistant according to the set type.
func GetValue[T any](ctx context.Context, key string) (result T) {
	a, err := fromContext(ctx)
	if err != nil {
		return result
	}

	value, ok := a.data[key]
	if !ok {
		return result
	}

	result, ok = value.(T)

	return result
}

// GetValueOK - returns the value by key from Assistant according to the set type.
func GetValueOK[T any](ctx context.Context, key string) (result T, ok bool) {
	a, err := fromContext(ctx)
	if err != nil {
		return result, false
	}

	value, ok := a.data[key]
	if !ok {
		return result, ok
	}

	result, ok = value.(T)

	return result, ok
}
