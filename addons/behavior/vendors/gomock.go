package vendors

type GoMockVendor interface {
	goMock()
}

// GoMock - behavior for method.
type GoMock[A, R any] struct {
	// Data - list of behaviors for a specific method.
	Data []GoMockData[A, R] `json:"data_assistant"`
}

type GoMockData[A, R any] struct {
	// Times - number of times the result is called. If -1, then set for any number of results.
	Times int `json:"times"`
	// Arguments - argument structure.
	// Note: use a sequential argument structure as specified in the signature.
	Arguments A `json:"arguments"`
	// Returns - return structure.
	// Note: use a sequential return structure as specified in the signature.
	Returns R `json:"returns"`
}

func (m *GoMock[A, R]) goMock() {}
