package assistant

func (a *Assistant) WithValue(key string, value any) *Assistant {
	a.data[key] = value

	return a
}

func (a *Assistant) WithValues(data map[string]any) *Assistant {
	for key, value := range data {
		a.data[key] = value
	}

	return a
}
