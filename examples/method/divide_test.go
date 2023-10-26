package method_test

import (
	"embed"
	"github.com/auvitly/lab/examples/method"
	"github.com/auvitly/lab/tools/kit"

	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed test
var fs embed.FS

func TestMethod(t *testing.T) {
	t.Parallel()

	kit.RequireRun(t, fs, func(
		t *testing.T,
		test kit.Test[
			*kit.In[struct {
				A float64 `json:"a"`
				B float64 `json:"b"`
			}],
			*kit.Out[float64, error],
		],
	) {
		result, err := method.Method(test.In.Args.A, test.In.Args.B)
		if err != nil {
			assert.EqualError(t, err, test.Out.Error.Error(), test.Title)

			return
		}

		assert.NoError(t, test.Out.Error, test.Title)
		assert.Equal(t, result, test.Out.Result, test.Title)
	})
}
