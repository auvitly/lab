package behaviour

// GoMock - behaviour for method.
type GoMock[A, R any] struct {
	// Data - list of behaviors for a specific method.
	Data []GoMockData[A, R] `json:"data"`
}

// GoMockData - case-specific behavior.
type GoMockData[A, R any] struct {
	// Times - number of times the result is called. If -1, then set for any number of results.
	Times int `json:"times"`
	// Argument - behavior for request.
	// Note: use assistant.Actual as a base solutions.
	Actual A `json:"arguments"`
	// Return - behavior for response.
	// Note: use assistant.Expect as a base solution.
	Return R `json:"returns"`
}
