package jubeo_consul

import (
	"github.com/latticework/proto-goquiles-server-todo/gojalicore"
	"github.com/shopspring/decimal"
	"time"
)

type ConsulConfiguration struct {
}

func (config *ConsulConfiguration) GetConsulValue(key string) (interface{}, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *ConsulConfiguration) GetValue(key string) (interface{}, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *ConsulConfiguration) GetStringValue(key string) (string, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *ConsulConfiguration) GetIntegerValue(key string) (int, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *ConsulConfiguration) GetDecimalValue(key string) (decimal.Decimal, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *ConsulConfiguration) GetDate(key string) (time.Time, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *ConsulConfiguration) GetDuration(key string) (time.Duration, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *ConsulConfiguration) GetTimestamp(key string) (time.Time, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}
