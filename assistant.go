package assistant

import (
	"context"
)

// Assistant - stores data about test settings and is able to place them in context.
type Assistant struct {
	data map[string]any
}

// New - returns an initialized Assistant.
func New() *Assistant {
	return &Assistant{
		data: make(map[string]any),
	}
}

// OnContext - applies the configuration to the parent context.
func (a *Assistant) OnContext(parent context.Context) context.Context {
	return context.WithValue(parent, key, a)
}

// fromContext - returns the Assistant, which must be in context, otherwise panic.
func fromContext(ctx context.Context) (*Assistant, error) {
	var value = ctx.Value(key)

	if value == nil {
		return nil, ErrNotFoundAssistant
	}

	result, ok := value.(*Assistant)
	if !ok {
		return nil, ErrNotValidAssistantKey
	}

	return result, nil
}
