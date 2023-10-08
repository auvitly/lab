package inventory_test

import (
	"embed"
	"github.com/auvitly/lab/tools/inventory"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:embed test/
var fs embed.FS

func TestMustParseTestsFromFS(t *testing.T) {
	var tests = inventory.MustLoadTestsFromFS[inventory.Test[
		inventory.In[any],
		inventory.Out[
			any,
			*inventory.Error,
		],
	]](fs, "test/test/parse/TestMustParseTestsFromFS.json")

	require.GreaterOrEqual(t, len(tests), 1)
}
