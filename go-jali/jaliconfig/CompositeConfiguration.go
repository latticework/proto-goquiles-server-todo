package jaliconfig

import (
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jalicore"
	"sort"
)

type CompositeConfiguration struct {
	configurators []Configurator
}

func NewCompositeConfiguration(config *CompositeConfig) (CompositeConfiguration, error) {
	if config == nil {
		config = DefaultCompositeConfig
	}

	if config.Configurators == nil {
		config.Configurators = DefaultCompositeConfig.Configurators
	}

	c := CompositeConfiguration{
		configurators: sort.Reverse(config.Configurators),
	}

	return c
}

func (config *CompositeConfiguration) GetJsonValue(
	key string, defaultValue map[string]interface{}) (map[string]interface{}, error) {
	for _, cfg := range config.configurators {
		json, err := cfg.GetJsonValue(key, _)
		if json != nil {
			return json, _
		}

		if !err.(*KeyNotFoundError) {
			return _, err
		}
	}

	if defaultValue != nil {
		return defaultValue, _
	}

	return _, KeyNotFoundError{}.Init(key)
}

func (config *CompositeConfiguration) GetValue(key string, defaultValue interface{}) (interface{}, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func buildCompositeConfig(config *CompositeConfig, defaultConfig *CompositeConfig) (*CompositeConfig, error) {
	if config == nil {
		newConfig := CompositeConfig{
			Configurators: defaultConfig.Configurators,
		}

		return newConfig, _
	}

	newConfig := CompositeConfig{
		Configurators: config.Configurators,
	}

	if newConfig.Configurators == nil {
		newConfig.Configurators = defaultConfig.Configurators
	}

	return newConfig
}
