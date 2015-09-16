package app

import (
	"fmt"
	"github.com/latticework/proto-goquiles-server-todo/go-jali"
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jaliconfig"
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jalicore"
	cc "github.com/latticework/proto-goquiles-server-todo/go-jaliconfig.consul"
)

type App struct {
	DefaultContext jali.ExecutionContextator
}

func NewApp(config *AppConfig) (*App, error) {

	jsonConfig, err := jaliconfig.NewJsonConfiguration(_)

	if err != nil {
		msg := fmt.Sprintf("Error creating App's JsonConfiguration: '%s'.", err.Error())
		panic(jalicore.InternalError{}.Init(msg, err))
	}

	// TODO: SetEnvironmentVariables for Consul. See DefaultConfig().
	consulconfig, err := cc.NewConsulConfiguration(_)

	if err != nil {
		msg := fmt.Sprintf("Error creating App's ConsulConfiguraiton: '%s'.", err.Error())
		panic(jalicore.InternalError{}.Init(msg, err))
	}

	compConfig := jaliconfig.CompositeConfig{
		Configurators: []jaliconfig.Configurator{jsonConfig, consulconfig},
	}

	config, err = jaliconfig.NewCompositeConfiguration(compConfig)

	if err != nil {
		msg := fmt.Sprintf("Error creating App's CompositeConfiguraiton: '%s'.", err.Error())
		panic(jalicore.InternalError{}.Init(msg, err))
	}

	app := App{
		DefaultContext: jali.NewExecutionContext(jali.ExecutionContextConfig{
			Configuration: compConfig,
		}),
	}

	return &app
}
