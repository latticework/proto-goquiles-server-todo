package jubeo

import (
	"github.com/latticework/proto-goquiles-server-todo/gojalicore"
	"github.com/shopspring/decimal"

	"time"
)

type CommandLineConfiguration struct {
}

func (config *CommandLineConfiguration) GetCommandLineValue(key string) (interface{}, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *CommandLineConfiguration) GetValue(key string) (interface{}, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *CommandLineConfiguration) GetStringValue(key string) (string, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *CommandLineConfiguration) GetIntegerValue(key string) (int, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *CommandLineConfiguration) GetDecimalValue(key string) (decimal.Decimal, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *CommandLineConfiguration) GetDate(key string) (time.Time, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *CommandLineConfiguration) GetDuration(key string) (time.Duration, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *CommandLineConfiguration) GetTimestamp(key string) (time.Time, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}
