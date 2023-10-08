package assistant_test

import (
	"github.com/auvitly/lab/tools/assistant"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssistant(t *testing.T) {
	data := map[string]any{
		"title": t.Name(),
		"dir":   t.TempDir(),
	}

	ctx := assistant.New().
		WithValues(data).
		NewContext()

	ctx = assistant.New().
		WithValue("title", "test").
		Context(ctx)

	assert.Equal(t, "test", assistant.GetValue[string](ctx, "title"))
	assert.Equal(t, data["dir"], assistant.GetValue[string](ctx, "dir"))
}
