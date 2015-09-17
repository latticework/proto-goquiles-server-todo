package quilesmodel

import "github.com/latticework/proto-goquiles-server-todo/go-jali/jalicore"

type ServiceParseError struct {
	jalicore.StructuredError
}

func (err *ServiceParseError) Init(message string, inner error) *ServiceParseError {
	err.StructuredError.Init(message, inner)

	return &err
}
