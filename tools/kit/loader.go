package kit

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// loadTests - parser test data_assistant from JSON.
func loadTests[T any](raw []byte) (tests []T, err error) {
	if len(raw) == 0 {
		return nil, ErrNotFoundTestData
	}

	if err = json.Unmarshal(raw, &tests); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrParsing, err)
	}

	return tests, nil
}

// obtainPath - finding file in folder with passed filename.
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

// LoadTests - parser test data_assistant from JSON by path.
func LoadTests[D any](fs embed.FS, path string) (tests []D, err error) {
	file, err := fs.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%w: %s",
			ErrFileNotFound,
			fmt.Errorf("fs.ReadFile: %w", err).Error(),
		)
	}

	tests, err = loadTests[D](file)
	if err != nil {
		return nil, err
	}

	return tests, nil
}

// MustLoadTests - must parser test data_assistant from JSON by path.
func MustLoadTests[T any](fs embed.FS, path string) (tests []T) {
	var err error

	tests, err = LoadTests[T](fs, path)
	if err != nil {
		panic(err)
	}

	return tests
}
