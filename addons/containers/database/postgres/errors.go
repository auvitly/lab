package postgres

import "errors"

var (
	// ErrInvalidOption - invalid option.
	ErrInvalidOption = errors.New("invalid option")
)

var (
	// ErrPortOutOfRange - port value must be in the range from 0 to 65535.
	ErrPortOutOfRange = errors.New("port value must be in the range from 0 to 65535")
	// ErrInvalidAuthParameters - invalid authorization parameters.
	ErrInvalidAuthParameters = errors.New("invalid authorization parameters")
	// ErrNotFoundNetwork - network is nil.
	ErrNotFoundNetwork = errors.New("network is nil")
)
