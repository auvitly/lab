package assistant

import (
	"context"
	"fmt"
)

// Inc - increases the counter by 1 unit. Available for options: WithCounter, WithLimit.
func Inc(ctx context.Context, key string) bool {
	a, err := fromContext(ctx)
	if err != nil {
		return false
	}

	value, ok := a.data[fmt.Sprintf("%s_counter", key)].(int)
	if !ok {
		return false
	}

	a.data[fmt.Sprintf("%s_counter", key)] = value + 1

	return true
}

// IsLimit - returns true if the counter is greater than or equal to the limit value.
// If there is no limit or the limit is zero, returns false.
func IsLimit(ctx context.Context, key string) bool {
	a, err := fromContext(ctx)
	if err != nil {
		return false
	}

	counter, ok := a.data[fmt.Sprintf("%s_counter", key)].(int)
	if !ok {
		return false
	}

	limit, ok := a.data[fmt.Sprintf("%s_limit", key)].(int)
	if !ok {
		return false
	}

	if limit == 0 {
		return false
	}

	return counter >= limit
}

// GetCounter - returns counter by key from Assistant.
func GetCounter(ctx context.Context, key string) (int, bool) {
	a, err := fromContext(ctx)
	if err != nil {
		return 0, false
	}

	counter, ok := a.data[fmt.Sprintf("%s_counter", key)].(int)
	if !ok {
		return 0, false
	}

	return counter, ok
}

// GetFlag - returns flag by key from Assistant.
func GetFlag(ctx context.Context, key string) (bool, bool) {
	a, err := fromContext(ctx)
	if err != nil {
		return false, false
	}

	flag, ok := a.data[fmt.Sprintf("%s_flag", key)].(bool)
	if !ok {
		return false, false
	}

	return flag, ok
}

// GetValue - returns the value by key from Assistant according to the set type.
func GetValue[T any](ctx context.Context, key string) (result T, ok bool) {
	a, err := fromContext(ctx)
	if err != nil {
		return result, false
	}

	value, ok := a.data[fmt.Sprintf("%s_value", key)]
	if !ok {
		return result, ok
	}

	result, ok = value.(T)

	return result, ok
}
