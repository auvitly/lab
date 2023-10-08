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

// Assistant - stores data_assistant about test settings and is able to place them in context.
type Assistant struct {
	data map[string]any
}

// New - returns an initialized Assistant.
func New() *Assistant {
	return &Assistant{
		data: make(map[string]any),
	}
}

// Context - applies the configuration to the parent context.
// * The parent context can be nil, then a new one will be created.
// * If the assistant already exists in the context, then the data will be taken from the
// * parent assistant and overwritten by the child.
func (a *Assistant) Context(parent context.Context) context.Context {
	// Check nil pointer.
	if parent == nil {
		return a.NewContext()
	}
	// Try to find assistant from parent context.
	founded, ok := parent.Value(assistantKey).(*Assistant)
	if !ok {
		return context.WithValue(parent, assistantKey, a)
	}
	// If founder, when building new values.
	var values = make(map[string]any)
	// Copy data from parent assistant.
	for key, value := range founded.data {
		values[key] = value
	}
	// Overwrite child assistant data.
	for key, value := range a.data {
		values[key] = value
	}
	// Set new values.
	a.data = values
	// Return updated parent context.
	return context.WithValue(parent, assistantKey, a)
}

// NewContext - return context with assistant.
func (a *Assistant) NewContext() context.Context {
	return context.WithValue(context.Background(), assistantKey, a)
}

// fromContext - returns the Assistant, which must be in context, otherwise panic.
func fromContext(ctx context.Context) (*Assistant, error) {
	var value = ctx.Value(assistantKey)

	if value == nil {
		return nil, ErrNotFoundAssistant
	}

	result, ok := value.(*Assistant)
	if !ok {
		return nil, ErrNotValidAssistantKey
	}

	return result, nil
}
