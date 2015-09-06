package quiles

import "github.com/latticework/proto-goquiles-server-todo/go-jali"

type RoutineContextator interface {
	ServiceContext() ServiceContextator
	ExecutionContext() jali.ExecutionContextator
	Input() chan ServiceMessageator
	Output() chan ServiceMessageator
}
