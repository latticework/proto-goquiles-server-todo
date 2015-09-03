package jali

import (
	"github.com/latticework/proto-goquiles-server-todo/gojubeo"
	"github.com/latticework/proto-goquiles-server-todo/gonotado"
)

type ExecutionContextator interface {
	Configuration() jubeo.Configurator
	Log() notado.StructuredLogger
}
