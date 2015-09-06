package jali

import (
	"./jaliconfig"
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jalilog"
)

type ExecutionContextator interface {
	Configuration() jaliconfig.Configurator
	Log() jalilog.StructuredLogger
}
