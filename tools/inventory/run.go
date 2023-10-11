package inventory

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Addon - interface for addon matching.
// The addon allows you to enable add-ons that will finish their work along with test processing.
type Addon interface {
	Init() error
	io.Closer
}

// Run - running spreadsheet tests based on data from the embed.FS and file path.
func Run[T any](t *testing.T, fs embed.FS, use func(t *testing.T, test T), addons ...Addon) error {
	t.Helper()
	// * Finding file.
	path, err := obtainPath(fs, t.Name())
	if err != nil {
		return err
	}
	// * Load tests.
	tests, err := LoadTests[T](fs, strings.ReplaceAll(path, "\\", "/"))
	if err != nil {
		return err
	}
	// * Init addons.
	for _, addon := range addons {
		if err = addon.Init(); err != nil {
			return fmt.Errorf("%w: %s", ErrInitAddon, err.Error())
		}
	}
	// * Run tests.
	for i := range tests {
		var desc = t.Name()

		if test, ok := any(tests[i]).(testTitler); ok {
			desc = test.title(t)
		}

		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			use(t, tests[i])
		})
	}
	// * Closing on defer addons.
	for _, addon := range addons {
		if err = addon.Close(); err != nil {
			return fmt.Errorf("%w: %s", ErrCloseAddon, err.Error())
		}
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
func MustRun[T any](t *testing.T, fs embed.FS, use func(t *testing.T, test T), addons ...Addon) {
	t.Helper()

	err := Run[T](t, fs, use, addons...)
	if err != nil {
		panic(err)
	}
}
