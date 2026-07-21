package core

import (
	"strings"
)

type Errors []FieldError

func (e Errors) Error() string {

	var messages []string

	for _, err := range e {

		messages = append(
			messages,
			err.Error(),
		)
	}

	return strings.Join(messages, ", ")
}
