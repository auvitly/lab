package vendors

// Method - behavior for method.
type Method[I, O any] []*Call[I, O]

// Call - function mock call description.
type Call[I, O any] struct {
	// In - method arguments.
	In I `json:"in"`
	// Out - return fields.
	Out O `json:"out"`
	// Times - number of times the result is called. If -1, then set for any number of results.
	Times int `json:"times"`
}
