package kit_test

import (
	"embed"
	"github.com/auvitly/lab/tools/kit"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:embed test
var run embed.FS

func TestRun_Success(t *testing.T) {
	require.NoError(t, kit.Run(t, run,
		func(
			t *testing.T,
			test kit.Test[kit.Empty, *kit.Out[kit.Empty, error]],
		) {
			t.Log("Success")
		},
	),
	)
}

func TestRun_ErrFileConflictName(t *testing.T) {
	require.ErrorIs(t, kit.Run(t, run,
		func(
			t *testing.T,
			test kit.Test[kit.Empty, kit.Empty],
		) {
			t.Log("Success")
		},
	), kit.ErrFileConflictName)
}

func TestRun_ErrNotFoundTestData(t *testing.T) {
	require.ErrorIs(t, kit.Run(t, run,
		func(
			t *testing.T,
			tests *kit.Test[kit.Empty, kit.Empty],
		) {
			t.Log("Success")
		},
	), kit.ErrNotFoundTestData)
}
