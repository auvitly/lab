package divide_test

import (
	"embed"
	"github.com/auvitly/lab/tools/inventory"
	"github.com/auvitly/lab/tools/inventory/test/examples/divide"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed test
var fs embed.FS

func TestDivide(t *testing.T) {
	t.Parallel()

	inventory.MustRun(t, fs, func(
		t *testing.T,
		test inventory.Test[
			*inventory.In[struct {
				A float64 `json:"a"`
				B float64 `json:"b"`
			}],
			*inventory.Out[float64, error],
		],
	) {
		result, err := divide.Divide(test.In.Arguments.A, test.In.Arguments.B)
		if err != nil {
			assert.EqualError(t, err, test.Out.Error.Error(), test.Title)

			return
		}

		assert.NoError(t, test.Out.Error, test.Title)
		assert.Equal(t, result, test.Out.Result, test.Title)
	})
}
