package assistant

import (
	"fmt"
)

// WithLimit - allows to add a counter with a limit that can be increased using the Inc() command.
func (a *Assistant) WithLimit(key string, limit int) *Assistant {
	a.data[fmt.Sprintf("%s_counter", key)] = 0
	a.data[fmt.Sprintf("%s_limit", key)] = limit

	return a
}

// WithCounter - allows to add a counter that can be increased using the Inc() command.
func (a *Assistant) WithCounter(key string) *Assistant {
	a.data[fmt.Sprintf("%s_counter", key)] = 0

	return a
}

func (a *Assistant) WithFlag(key string, flag bool) *Assistant {
	a.data[fmt.Sprintf("%s_flag", key)] = flag

	return a
}

func (a *Assistant) WithValue(key string, value any) *Assistant {
	a.data[fmt.Sprintf("%s_value", key)] = value

	return a
}

func (a *Assistant) WithValues(data map[string]any) *Assistant {
	for key, value := range data {
		a.data[fmt.Sprintf("%s_value", key)] = value
	}

	return a
}