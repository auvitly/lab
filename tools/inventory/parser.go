package inventory

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
)

// Parse errors.
var (
	// ErrFileConflictName - filename conflict with test data.
	ErrFileConflictName = errors.New("filename conflict with test data")
	// ErrFileNotFound - must be at least one test.
	ErrFileNotFound = errors.New("file not found")
	// ErrNotFoundTestData - not found test data.
	ErrNotFoundTestData = errors.New("not found test data")
	// ErrNotFoundTests - must be at least one test.
	ErrNotFoundTests = errors.New("must be at least one test")
)

// loadTests - parser test data_assistant from JSON.
func loadTests[T Tester](raw []byte) (tests []T, err error) {
	if len(raw) == 0 {
		return nil, ErrNotFoundTestData
	}

	if err = json.Unmarshal(raw, &tests); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	if len(tests) == 0 {
		return nil, ErrNotFoundTests
	}

	return tests, nil
}

// LoadTests - parser test data_assistant from JSON by path.
func LoadTests[T Tester](fs embed.FS, path string) (tests []T, err error) {
	file, err := fs.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%w: %s",
			ErrFileNotFound,
			fmt.Errorf("fs.ReadFile: %w", err).Error(),
		)
	}

	tests, err = loadTests[T](file)
	if err != nil {
		return nil, err
	}

	return tests, nil
}

// MustLoadTests - must parser test data_assistant from JSON by path.
func MustLoadTests[T Tester](fs embed.FS, path string) (tests []T) {
	var err error

	tests, err = LoadTests[T](fs, path)
	if err != nil {
		panic(err)
	}

	return tests
}
