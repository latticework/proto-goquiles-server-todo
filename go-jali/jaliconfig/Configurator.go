package jaliconfig

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
	"time"

	"github.com/latticework/proto-goquiles-server-todo/go-jali/jalicore"
)

type Configurator interface {
	GetJsonValue(key string, defaultValue map[string]interface{}) (map[string]interface{}, error)
	GetValue(key string, defaultValue interface{}) (interface{}, error)
	//	GetStringValue(key string, defaultValue string) (string, error)
	//	GetIntegerValue(key string, defaultValue int) (int, error)
	//	GetDecimalValue(key string, defaultValue decimal.Decimal) (decimal.Decimal, error)
	//	GetDate(key string, defaultValue time.Time) (time.Time, error)
	//	GetDuration(key string, defaultValue time.Duration) (time.Duration, error)
	//	GetTimestamp(key string, defaultValue time.Time) (time.Time, error)
}

// https://github.com/luciotato/golang-notes/blob/master/second-level%20methods.md
func (config Configurator) GetStringValue(key string, defaultValue string) (string, error) {
	v, err := config.GetValue(key, defaultValue)

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

func (config Configurator) GetIntegerValue(key string, defaultValue int) (int, error) {
	v, err := config.GetValue(key, defaultValue)

	if err != nil {
		return _, err
	}

	return strconv.Atoi(v)
}

func (config Configurator) GetDecimalValue(key string, defaultValue decimal.Decimal) (decimal.Decimal, error) {
	return _, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config Configurator) GetDate(key string, defaultValue time.Time) (time.Time, error) {
	return _, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config Configurator) GetDuration(key string, defaultValue time.Duration) (time.Duration, error) {
	return _, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config Configurator) GetTimestamp(key string, defaultValue time.Time) (time.Time, error) {
	return _, jalicore.NotImplementedError{}.Init(nil, nil)
}
