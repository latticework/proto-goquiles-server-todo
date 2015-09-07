package jali

import (
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jaliconfig"
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jalilog"
)

func DefaultExecutionContextCofig() *ExecutionContextConfig {
	return &ExecutionContextConfig{
		Configuration: jaliconfig.NewCompositeConfiguration(_),
		Log:           jalilog.NewConsoleLog(_),
	}
}
