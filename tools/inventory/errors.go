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
	// ErrNotFoundTests - must be at least one test.
	ErrNotFoundTests = errors.New("must be at least one test")
	// ErrParsing - json parsing error.
	ErrParsing = errors.New("json parsing error")
	// ErrInitAddon - addon initialization error occurs.
	ErrInitAddon = errors.New("init addon error")
	// ErrCloseAddon - addon close error occurs.
	ErrCloseAddon = errors.New("close addon error")
)
