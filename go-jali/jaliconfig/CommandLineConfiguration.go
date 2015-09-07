package jaliconfig

import (
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jalicore"
)

type CommandLineConfiguration struct {
}

func (config *CommandLineConfiguration) GetJsonValue(
	key string, defaultValue map[string]interface{}) (map[string]interface{}, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *CommandLineConfiguration) GetValue(key string, defaultValue interface{}) (interface{}, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}
