package jalicore

import "fmt"

type ArgumentError struct {
	StructuredError
	Name string
}

func (err *ArgumentError) Init(name string, message string, inner error) {
	msg := message

	switch {
	case msg != nil:
	case name != nil || len(name) > 0:
		msg = fmt.Sprintf("Argument '%v' is in error.", name)
	default:
		msg = "Argument is in error."
	}

	err.StructuredError.Init(msg, inner)
	err.Name = name
}