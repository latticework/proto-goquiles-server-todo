package notado

import (
	diles "github.com/latticework/proto-goquiles-server-todo/godiles"
)

type LogInformer interface {
	Context() LogContext
	Categories() []string
	MessageCode() uint32
	Priority() diles.MessagePriority
	Severity() diles.MessageSeverity
	ServiceCenterId() string
	Message() string
	Args() []interface{}
	// ServiceName() string?
	// ResourceName() string?
	// RoutineName() string?
	// FunctionName() string?
	FileName() string // TODO: LogInformer.FileName: Are these subsumed by CallStack?
	LineNumber() uint16
	CallStack() string
}
