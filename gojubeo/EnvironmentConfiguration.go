package jubeo

import (
	"github.com/latticework/proto-goquiles-server-todo/gojalicore"
	"github.com/shopspring/decimal"
	"time"
)

type EnvironmentConfiguration struct {
}

func (config *EnvironmentConfiguration) GetJsonValue(key string) (interface{}, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *EnvironmentConfiguration) GetValue(key string) (interface{}, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *EnvironmentConfiguration) GetStringValue(key string) (string, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *EnvironmentConfiguration) GetIntegerValue(key string) (int, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *EnvironmentConfiguration) GetDecimalValue(key string) (decimal.Decimal, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *EnvironmentConfiguration) GetDate(key string) (time.Time, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *EnvironmentConfiguration) GetDuration(key string) (time.Duration, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *EnvironmentConfiguration) GetTimestamp(key string) (time.Time, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}
