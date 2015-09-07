package jaliconfig

import "github.com/hashicorp/consul/api"

const DefaultConfig = ConsulConfig{
	Config: api.DefaultConfig(),
}
