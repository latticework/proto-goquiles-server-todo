package jali

import (
	"github.com/shopspring/decimal"
	"time"
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
