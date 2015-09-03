package jubeo

import (
	"github.com/latticework/proto-goquiles-server-todo/gojalicore"
	"github.com/shopspring/decimal"
	"time"
)

type JsonConfiguration struct {
}

func (config *JsonConfiguration) GetJsonValue(key string) (interface{}, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *JsonConfiguration) GetValue(key string) (interface{}, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *JsonConfiguration) GetStringValue(key string) (string, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *JsonConfiguration) GetIntegerValue(key string) (int, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *JsonConfiguration) GetDecimalValue(key string) (decimal.Decimal, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *JsonConfiguration) GetDate(key string) (time.Time, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *JsonConfiguration) GetDuration(key string) (time.Duration, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *JsonConfiguration) GetTimestamp(key string) (time.Time, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}
