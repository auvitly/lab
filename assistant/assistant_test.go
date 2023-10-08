package assistant_test

import (
	"context"
	"github.com/auvitly/assistant/assistant"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssistant(t *testing.T) {
	ctx, cancel := assistant.NewContextAssistant().
		WithValue("key", "test").
		OnContext(context.Background())

	defer cancel()

	result, _ := assistant.GetValue[string](ctx, "key")

	assert.Equal(t, result, "test")
}
