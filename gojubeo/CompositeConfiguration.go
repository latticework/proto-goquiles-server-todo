package jubeo

import (
	"github.com/latticework/proto-goquiles-server-todo/gojalicore"
	"github.com/shopspring/decimal"

	"time"
)

type CompositeConfiguration struct {
	configurators []Configurator
}

func (config *CompositeConfiguration) GetJsonValue(key string) (interface{}, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *CompositeConfiguration) GetValue(key string) (interface{}, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *CompositeConfiguration) GetStringValue(key string) (string, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *CompositeConfiguration) GetIntegerValue(key string) (int, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *CompositeConfiguration) GetDecimalValue(key string) (decimal.Decimal, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *CompositeConfiguration) GetDate(key string) (time.Time, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *CompositeConfiguration) GetDuration(key string) (time.Duration, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *CompositeConfiguration) GetTimestamp(key string) (time.Time, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}
