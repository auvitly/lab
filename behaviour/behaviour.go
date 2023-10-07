package behaviour

// Set - set behavior for mock by supported type.
func Set(behaviour any) error {
	var support S

	switch any(support).(type) {
	case *GoMock[]
	}

	return nil
}
