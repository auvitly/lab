package inventory

import (
	"errors"
	"fmt"
	"strings"
)

func closeAddons(activated []Addon) (err error) {
	var messages []string

	for _, addon := range activated {
		if err = addon.Close(); err != nil {
			messages = append(messages, fmt.Errorf("%T: %s", addon, err).Error())
		}
	}

	if len(messages) != 0 {
		return errors.New(strings.Join(messages, " ;"))
	}

	return nil
}

func initAddons(addons []Addon) (activated []Addon, err error) {
	for _, addon := range addons {
		if err = addon.Init(); err != nil {
			return activated, fmt.Errorf("%T: %s", addon, err)
		}

		activated = append(activated, addon)
	}

	return activated, nil
}
