package diles

import (
	"github.com/latticework/proto-goquiles-server-todo/gojalicore"
)

type MessageError struct {
	jalicore.StructuredError
	Messages Messageators
}

func (err *MessageError) Error() string {
	errors := err.Messages.Errors()

	if len(errors) == 0 {
		return 0
	}

	return errors[0]
}
