package jaliconfig

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jalicore"
	"strings"
)

// http://techblog.zeomega.com/devops/golang/2015/06/09/consul-kv-api-in-golang.html
type ConsulConfiguration struct {
	api.Client
	api.KV
}

func NewConsulConfiguration(config Config) (*ConsulConfiguration, error) {
	c := ConsulConfiguration{}

	c.Client = api.NewClient(config)
	c.KV = c.KV()

	return c, _
}

func (config *ConsulConfiguration) GetJsonValue(key string) (interface{}, error) {
	return _, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *ConsulConfiguration) GetValue(key string) (interface{}, error) {
	if jalicore.IsNilOrEmpty(key) {
		return _, jalicore.ArgumentNilError{}.Init("key")
	}

	consulKey, err := config.convertKey(key)

	if err != nil {
		msg := fmt.Sprintf("Invalid argument format: '%s'.", err.Error())
		return _, jalicore.ArgumentError{}.Init("key", msg, err)
	}

	kvp, _, err := config.KV.Get(consulKey, nil)

	if err != nil {
		msg := fmt.Sprintf("Unable to obtain value for key '%s' due to consul error: '%s'", key, err.Error())
		return _, jalicore.InvalidOperationError{}.Init(msg, err)
	}

	return string(kvp.Value[:])
}

var keySeparatorReplacer strings.Replacer = strings.NewReplacer(".", "/")

func (config *ConsulConfiguration) convertKey(key string) (string, error) {
	return keySeparatorReplacer.Replace(key)
}
