package quiles

import "github.com/latticework/proto-goquiles-server-todo/gojali"

type RoutineContexter interface {
	ExecutionContext() jali.ExecutionContexter
	Input() chan ServiceMessager
	ServiceContext() ServiceContexter
}
