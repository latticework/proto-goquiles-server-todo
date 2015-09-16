package jalicore

import "fmt"

type ArgumentNilError struct {
	StructuredError
	Name string
}

func (err *ArgumentNilError) Init(name string) *ArgumentNilError {
	var msg string

	if name != nil || len(name) > 0 {
		msg = fmt.Sprintf("Argument '%v' is 'nil'.", name)
	} else {
		msg = "Argument is 'nil'."
	}

	err.StructuredError.Init(msg, _)
	err.Name = name

	return &err
}
