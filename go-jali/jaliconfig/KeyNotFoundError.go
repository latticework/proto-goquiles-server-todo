package jaliconfig

import (
	"fmt"
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jalicore"
)

type KeyNotFoundError struct {
	jalicore.StructuredError
	Path string
}

func (err *KeyNotFoundError) Init(path string) {
	msg := fmt.Sprintf("Could not find path '%s' when resolving for key", path)
	err.StructuredError.Init(msg, _)
	err.Path = path
}
