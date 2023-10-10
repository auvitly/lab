package inventory_test

import (
	"embed"
	"github.com/auvitly/lab/tools/inventory"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:embed test/data/parser
var parser embed.FS

func TestMustLoadTests_Success(t *testing.T) {
	defer func() {
		require.Nil(t, recover(), t.Name())
	}()

	tests := inventory.MustLoadTests[inventory.Test[
		inventory.Empty,
		*inventory.Out[inventory.Empty, error],
	]](parser, "test/data/parser/TestMustLoadTests_Success.json")

	for _, test := range tests {
		if test.Out.Error != nil {
			require.NotNil(t, test.Out.Error)
			require.Greater(t, len(test.Out.Error.Error()), 0)
		} else {
			require.Nil(t, test.Out.Error)
		}
	}
}

func TestMustLoadTests_ErrNotFoundTests(t *testing.T) {
	defer func() {
		require.ErrorIs(t, recover().(error), inventory.ErrNotFoundTests, t.Name())
	}()

	inventory.MustLoadTests[inventory.Test[
		inventory.Empty,
		*inventory.Out[inventory.Empty, error],
	]](parser, "test/data/parser/TestMustLoadTests_ErrNotFoundTests.json")
}

func TestMustLoadTests_ErrNotFoundTestData(t *testing.T) {
	defer func() {
		require.ErrorIs(t, recover().(error), inventory.ErrNotFoundTestData, t.Name())
	}()

	inventory.MustLoadTests[inventory.Test[
		inventory.Empty,
		*inventory.Out[inventory.Empty, error],
	]](parser, "test/data/parser/TestMustLoadTests_ErrNotFoundTestData.json")
}

func TestMustLoadTests_ErrParsing(t *testing.T) {
	defer func() {
		require.ErrorIs(t, recover().(error), inventory.ErrParsing, t.Name())
	}()

	inventory.MustLoadTests[inventory.Test[
		inventory.Empty,
		*inventory.Out[inventory.Empty, error],
	]](parser, "test/data/parser/TestMustLoadTests_ErrParsing.json")
}

func TestMustLoadTests_ErrFileNotFound(t *testing.T) {
	defer func() {
		require.ErrorIs(t, recover().(error), inventory.ErrFileNotFound, t.Name())
	}()

	inventory.MustLoadTests[inventory.Test[
		inventory.Empty,
		*inventory.Out[inventory.Empty, error],
	]](parser, "")
}
