package inventory

import "errors"

// Parse errors.
var (
	// ErrFileConflictName - filename conflict with test data.
	ErrFileConflictName = errors.New("filename conflict with test data")
	// ErrFileNotFound - must be at least one test.
	ErrFileNotFound = errors.New("file not found")
	// ErrNotFoundTestData - not found test data.
	ErrNotFoundTestData = errors.New("not found test data")
	// ErrParsing - json parsing error.
	ErrParsing = errors.New("json parsing error")
	// ErrNotSupportingImplementation - not supporting implementation.
	ErrNotSupportingImplementation = errors.New("not supporting implementation")
)
