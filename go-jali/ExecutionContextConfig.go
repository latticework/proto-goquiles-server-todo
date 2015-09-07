package jali

import (
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jaliconfig"
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jalilog"
)

type ExecutionContextConfig struct {
	Configuration jaliconfig.Configurator
	Log           jalilog.StructuredLogger
}
