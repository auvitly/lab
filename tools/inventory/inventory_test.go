package inventory_test

import (
	"embed"
	"github.com/auvitly/lab/tools/inventory"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:embed test/data/run
var run embed.FS

func TestMustRun_Success(t *testing.T) {
	defer func() {
		require.Nil(t, recover(), t.Name())
	}()

	inventory.MustRun(t, run, func(
		t *testing.T,
		test inventory.Test[
			inventory.Empty,
			*inventory.Out[inventory.Empty, error],
		],
	) {
		t.Log("Success")
	},
	)
}

func TestMustRun_ErrFileConflictName(t *testing.T) {
	defer func() {
		require.ErrorIs(t, recover().(error), inventory.ErrFileConflictName, t.Name())
	}()

	inventory.MustRun(t, run, func(
		t *testing.T,
		test inventory.Test[inventory.Empty, inventory.Empty],
	) {
		t.Log("Success")
	},
	)
}

func TestMustRun_ErrNotFoundTestData(t *testing.T) {
	defer func() {
		require.ErrorIs(t, recover().(error), inventory.ErrNotFoundTestData, t.Name())
	}()

	inventory.MustRun(t, run, func(
		t *testing.T,
		tests *inventory.Test[inventory.Empty, inventory.Empty],
	) {
		t.Log("Success")
	},
	)
}
