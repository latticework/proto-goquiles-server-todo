package quiles

import "github.com/latticework/proto-goquiles-server-todo/gojali"

type RoutineContextator interface {
	ServiceContext() ServiceContextator
	ExecutionContext() jali.ExecutionContextator
	Input() chan ServiceMessagator
	Output() chan ServiceMessagator
}
