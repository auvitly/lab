package inventory

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Run - running spreadsheet tests based on data from the embed.FS and file path.
func Run[T Tester](t *testing.T, fs embed.FS, use func(t *testing.T, test T)) error {
	t.Helper()

	path, err := obtainPath(fs, t.Name())
	if err != nil {
		return err
	}

	tests, err := LoadTests[T](fs, strings.ReplaceAll(path, "\\", "/"))
	if err != nil {
		return err
	}

	for i := range tests {
		t.Run(
			tests[i].name(t),
			func(t *testing.T) {
				t.Parallel()

				use(t, tests[i])
			},
		)
	}

	return nil
}

func obtainPath(fs embed.FS, name string) (result string, err error) {
	entities, err := fs.ReadDir(".")
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrFileNotFound, err.Error())
	}

	var founded []string

	for _, entity := range entities {
		if entity.IsDir() {
			err = filepath.Walk(entity.Name(), func(path string, info os.FileInfo, err error) error {
				if strings.Contains(info.Name(), fmt.Sprintf("%s.json", name)) {
					founded = append(founded, path)
				}

				return nil
			})
		}
	}

	switch {
	case len(founded) == 1:
		return founded[0], nil
	case len(founded) > 1:
		return "", ErrFileConflictName
	default:
		return "", ErrFileNotFound
	}
}

// MustRun - running spreadsheet tests based on data from the embed.FS and file path.
func MustRun[T Tester](t *testing.T, fs embed.FS, use func(t *testing.T, test T)) {
	t.Helper()

	err := Run[T](t, fs, use)
	if err != nil {
		panic(err)
	}
}
