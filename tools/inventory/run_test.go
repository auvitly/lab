package inventory_test

import (
	"embed"
	"fmt"
	"github.com/auvitly/lab/tools/inventory"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:embed test/data/run
var run embed.FS

func TestRun_Success(t *testing.T) {
	err := inventory.Run(t, run, func(
		t *testing.T,
		test *inventory.Test[any, any],
	) {
		fmt.Println(test)
	},
	)
	require.NoError(t, err, t.Name())
}

func TestRun_ErrNotFoundTests(t *testing.T) {
	err := inventory.Run(t, run, func(
		t *testing.T,
		test *inventory.Test[any, any],
	) {
	},
	)
	require.Error(t, err, inventory.ErrNotFoundTests)
}

func TestRun_ErrNotFoundTestData(t *testing.T) {
	err := inventory.Run(t, run,
		func(
			t *testing.T,
			test *inventory.Test[any, any],
		) {
		},
	)
	require.Error(t, err, inventory.ErrNotFoundTestData)
}

func TestRun_ErrFileNotFound(t *testing.T) {
	err := inventory.Run(t, run, func(
		t *testing.T,
		test *inventory.Test[any, any],
	) {
	},
	)
	require.Error(t, err, inventory.ErrFileNotFound)
}

func TestMustRun_Success(t *testing.T) {
	defer func() {
		require.Nil(t, recover(), t.Name())
	}()

	inventory.MustRun(t, run, func(
		t *testing.T,
		test *inventory.Test[any, any],
	) {
	},
	)
}

func TestMustRun_ErrNotFoundTests(t *testing.T) {
	defer func() {
		require.ErrorIs(t, recover().(error), inventory.ErrNotFoundTests, t.Name())
	}()

	inventory.MustRun(t, run, func(
		t *testing.T,
		test *inventory.Test[any, any],
	) {
	},
	)
}

func TestMustRun_ErrNotFoundTestData(t *testing.T) {
	defer func() {
		require.ErrorIs(t, recover().(error), inventory.ErrNotFoundTestData, t.Name())
	}()

	inventory.MustRun(t, run, func(
		t *testing.T,
		test *inventory.Test[any, any],
	) {
	},
	)
}

func TestMustRun_ErrFileNotFound(t *testing.T) {
	defer func() {
		require.ErrorIs(t, recover().(error), inventory.ErrFileNotFound, t.Name())
	}()

	inventory.MustRun(t, run, func(
		t *testing.T,
		test *inventory.Test[any, any],
	) {
	},
	)
}
