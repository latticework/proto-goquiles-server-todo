package jaliconfig

import (
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jalicore"
)

type EnvironmentConfiguration struct {
}

func (config *EnvironmentConfiguration) GetJsonValue(
	key string, defaultValue map[string]interface{}) (map[string]interface{}, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *EnvironmentConfiguration) GetValue(key string, defaultValue interface{}) (interface{}, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}
