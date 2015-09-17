package command

import (
	"fmt"
	"github.com/latticework/proto-goquiles-server-todo/go-quiles/quilesmodel"
	"strings"
)

type ServerCommand struct {
	Meta
}

func (c *ServerCommand) Run(args []string) int {
	// TODO: Ref https://github.com/hashicorp/vault/blob/master/command/server.go

	flags := c.Meta.FlagSet("server", FlagSetDefault)
	flags.Usage = func() { c.Ui.Error(c.Help()) }

	// @kenbrubaker: How to do validation...
	// Validation
	//	if !dev && len(configPath) == 0 {
	//		c.Ui.Error("At least one config path must be specified with -config")
	//		flags.Usage()
	//		return 1
	//	}

	service, err := quilesmodel.DecodeService(nil)

	if err != nil {
		c.Ui.Error(err.Error())
	}

	c.Ui.Output(fmt.Sprintf("Service Name: '%s'", service.Name))
	c.Ui.Output(fmt.Sprintf("Url: '%s'", service.Url))

	return 0
}

func (c *ServerCommand) Synopsis() string {
	return "Start quilessvr listing on the configured port."
}

func (c *ServerCommand) Help() string {
	helpText := `
Usage: quilessvr server [options]

	Starts the quilessvr service.

	This command starts quilessvr listing on the configured port.

General Options:
  ` + generalOptionsUsage() + `
Send Options:
  -send-option=value      The send option being explained that has a
  						  default value.
`
	return strings.TrimSpace(helpText)
}
