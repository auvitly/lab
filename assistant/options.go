package assistant

import (
	"fmt"
)

// WithLimit - allows to add a counter with a limit that can be increased using the Inc() command.
func (a *ContextAssistant) WithLimit(key string, limit int) *ContextAssistant {
	a.data[fmt.Sprintf("%s_counter", key)] = 0
	a.data[fmt.Sprintf("%s_limit", key)] = limit

	return a
}

// WithCounter - allows to add a counter that can be increased using the Inc() command.
func (a *ContextAssistant) WithCounter(key string) *ContextAssistant {
	a.data[fmt.Sprintf("%s_counter", key)] = 0

	return a
}

func (a *ContextAssistant) WithFlag(key string, flag bool) *ContextAssistant {
	a.data[fmt.Sprintf("%s_flag", key)] = flag

	return a
}

func (a *ContextAssistant) WithValue(key string, value any) *ContextAssistant {
	a.data[fmt.Sprintf("%s_value", key)] = value

	return a
}
