package assistant

import "errors"

// General assistant errors.
var (
	// ErrNotFoundAssistant - couldn't find assistant in context.
	ErrNotFoundAssistant = errors.New("couldn't find assistant in context")
	// ErrNotValidAssistantKey - the error occurs if the assistant's technical key has been compromised.
	ErrNotValidAssistantKey = errors.New("not valid assistant key")
)

// Parse errors.
var (
	// ErrNotFoundTestData - no date found for downloading tests.
	ErrNotFoundTestData = errors.New("no date found for downloading tests")
	// ErrNotFoundTests - tests could not be found.
	ErrNotFoundTests = errors.New("tests could not be found")
)
