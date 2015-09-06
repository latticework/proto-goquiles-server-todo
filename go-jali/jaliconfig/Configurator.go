package jaliconfig

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
	"time"

	"github.com/latticework/proto-goquiles-server-todo/go-jali/jalicore"
)

type Configurator interface {
	GetJsonValue(key string) (interface{}, error)
	GetValue(key string) (interface{}, error)
	GetStringValue(key string) (string, error)
	GetIntegerValue(key string) (int, error)
	GetDecimalValue(key string) (decimal.Decimal, error)
	GetDate(key string) (time.Time, error)
	GetDuration(key string) (time.Duration, error)
	GetTimestamp(key string) (time.Time, error)
}

// https://github.com/luciotato/golang-notes/blob/master/second-level%20methods.md
func (config *Configurator) GetStringValue(key string) (string, error) {
	v, err := config.GetValue(key)

	if v == nil || err != nil {
		return _, err
	}

	strValue, ok := v.(string)

	if ok {
		return strValue, _
	}

	msg := fmt.Sprintf("Value '%v' for key '%s' is not a valid string.", v, key)

	return _, jalicore.InvalidOperationError{}.Init(msg, _)
}

func (config *Configurator) GetIntegerValue(key string) (int, error) {
	v, err := config.GetValue(key)

	if err != nil {
		return _, err
	}

	return strconv.Atoi(v)
}

func (config *Configurator) GetDecimalValue(key string) (decimal.Decimal, error) {
	return _, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *Configurator) GetDate(key string) (time.Time, error) {
	return _, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *Configurator) GetDuration(key string) (time.Duration, error) {
	return _, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *Configurator) GetTimestamp(key string) (time.Time, error) {
	return _, jalicore.NotImplementedError{}.Init(nil, nil)
}
