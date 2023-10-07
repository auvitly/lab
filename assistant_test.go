package assistant_test

import (
	"context"
	"embed"
	"github.com/auvitly/assistant"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:embed test
var fs embed.FS

func TestAssistant_OnContext(t *testing.T) {
	ctx := assistant.New().
		WithValue("key", "test").
		OnContext(context.Background())

	result, _ := assistant.Value[string](ctx, "key")

	assert.Equal(t, result, "test")
}

func TestMustParseTestsFromFS(t *testing.T) {
	var tests = assistant.MustLoadTestsFromFS[assistant.Test[
		assistant.ActualWithBehaviour[
			any,
			any,
		],
		assistant.Expect[
			any,
			*assistant.Error,
		],
	]](fs, "test/TestMustParseTestsFromFS.json")

	require.GreaterOrEqual(t, len(tests), 1)
}
