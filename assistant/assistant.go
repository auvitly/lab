package assistant

import (
	"context"
	"errors"
)

// General assistant errors.
var (
	// ErrNotFoundAssistant - couldn't find assistant in context.
	ErrNotFoundAssistant = errors.New("couldn't find assistant in context")
	// ErrNotValidAssistantKey - the error occurs if the assistant's technical key has been compromised.
	ErrNotValidAssistantKey = errors.New("not valid assistant key")
)

// ContextAssistant - stores data_assistant about test settings and is able to place them in context.
type ContextAssistant struct {
	data map[string]any
}

// NewContextAssistant - returns an initialized Assistant.
func NewContextAssistant() *ContextAssistant {
	return &ContextAssistant{
		data: make(map[string]any),
	}
}

// OnContext - applies the configuration to the parent context.
func (a *ContextAssistant) OnContext(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(context.WithValue(parent, key, a))
}

// fromContext - returns the Assistant, which must be in context, otherwise panic.
func fromContext(ctx context.Context) (*ContextAssistant, error) {
	var value = ctx.Value(key)

	if value == nil {
		return nil, ErrNotFoundAssistant
	}

	result, ok := value.(*ContextAssistant)
	if !ok {
		return nil, ErrNotValidAssistantKey
	}

	return result, nil
}
