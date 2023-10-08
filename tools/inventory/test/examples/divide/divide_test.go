package divide_test

import (
	"embed"
	"fmt"
	"github.com/auvitly/lab/tools/inventory"
	"github.com/auvitly/lab/tools/inventory/test/examples/divide"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed test
var fs embed.FS

func TestDivide(t *testing.T) {
	t.Parallel()

	var tests = inventory.MustLoadTestsFromFS[inventory.Test[
		struct {
			A float64 `json:"a"`
			B float64 `json:"b"`
		},
		struct {
			C     float64          `json:"c"`
			Error *inventory.Error `json:"error"`
		},
	]](fs, fmt.Sprintf("test/%s.json", t.Name()))

	for i := range tests {
		var test = tests[i]

		t.Run(tests[i].Title, func(tt *testing.T) {
			tt.Parallel()

			c, err := divide.Divide(test.In.A, test.In.B)
			if err != nil {
				assert.EqualError(tt, err, test.Out.Error.Error(), test.Title)

				return
			}

			assert.NoError(tt, test.Out.Error.Extract(), test.Title)
			assert.Equal(tt, c, test.Out.C, test.Title)
		})
	}
}
