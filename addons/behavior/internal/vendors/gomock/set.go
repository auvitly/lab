package gomock

import "github.com/golang/mock/gomock"

// Set - sets the behavior for the passed mock.
func Set(mock any, behaviour any) error {
	sig, err := parseMock(mock)
	if err != nil {
		return err
	}

	_ = sig

	data, err := parseBehaviour(behaviour)
	if err != nil {
		return err
	}

	_ = data

	gomock.Any()

	return err
}
