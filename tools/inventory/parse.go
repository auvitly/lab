package inventory

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
)

// Parse errors.
var (
	// ErrNotFoundTestData - no date found for downloading tests.
	ErrNotFoundTestData = errors.New("no date found for downloading tests")
	// ErrNotFoundTests - tests could not be found.
	ErrNotFoundTests = errors.New("tests could not be found")
)

// ParseTests - parse test data_assistant from JSON.
func ParseTests[T Tester](raw []byte) (tests []T, err error) {
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

// MustLoadTests - must parse test data_assistant from JSON.
func MustLoadTests[T Tester](raw []byte) (tests []T) {
	var err error

	if tests, err = ParseTests[T](raw); err != nil {
		panic(err)
	}

	return tests
}

// ParseTestsFromFS - parse test data_assistant from JSON by path.
func ParseTestsFromFS[T Tester](fs embed.FS, path string) (tests []T, err error) {
	file, err := fs.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("fs.ReadFile: %w", err)
	}

	tests, err = ParseTests[T](file)
	if err != nil {
		return nil, err
	}

	return tests, nil
}

// MustLoadTestsFromFS - must parse test data_assistant from JSON by path.
func MustLoadTestsFromFS[T Tester](fs embed.FS, path string) (tests []T) {
	var err error

	tests, err = ParseTestsFromFS[T](fs, path)
	if err != nil {
		panic(err)
	}

	return tests
}
