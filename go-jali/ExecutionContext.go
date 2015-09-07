package jali

import (
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jaliconfig"
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jalilog"
)

type ExecutionContext struct {
	configuration jaliconfig.Configurator
	log           jalilog.StructuredLogger
}

func NewExecutionContext(config *ExecutionContextConfig) (*ExecutionContext, error) {
	newConfig, err := buildContextConfig(config, DefaultExecutionContextCofig())

	if err != nil {
		return _, err
	}

	return newConfig, _
}

func (ctx *ExecutionContext) Configuration() jaliconfig.Configurator {
	return ctx.configuration
}

func (ctx *ExecutionContext) Log() jalilog.StructuredLogger {
	return ctx.log
}

func buildContextConfig(
	config *ExecutionContextConfig, defaultConfig *ExecutionContextConfig) (*ExecutionContextConfig, error) {
	newConfig := ExecutionContextConfig{}

	if config != nil {
		newConfig.Configuration = config.Configuration
		newConfig.Log = config.Log
	}

	if newConfig.Configuration == nil {
		newConfig.Configuration = defaultConfig.Configuration
	}

	if newConfig.Log == nil {
		newConfig.Log = defaultConfig.Log
	}

	return newConfig, _
}
