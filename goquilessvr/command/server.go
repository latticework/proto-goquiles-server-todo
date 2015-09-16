package command

import (
	"fmt"
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jalicore"
	"path/filepath"
	"strings"
)

type ServerCommand struct {
	Meta
}

func (c *ServerCommand) Run(args []string) int {
	var requestUrl, requestJson string

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

	pwdAbsolute, err := filepath.Abs("./")

	if err != nil {
		msg := "Error getting current working directory: '%s'", err.Error()
		c.Ui.Error(msg)
		return 1
	}

	serviceDefinitionFile, err := jalicore.FindSettingsFile("servicequiles.json", pwdAbsolute)

	if err != nil {
		var msg string

		switch err := err.(type) {
		case jalicore.StructuredError:
			msg = fmt.Sprintf(
				"Unable to find service definition file from location '%s'.\nError Type: '%s'\nMessage:\n'%s'\nStack:\n%s",
				pwdAbsolute, err.TypeName(), err.Error(), err.ErrorStack())
		default:
			msg = fmt.Sprintf(
				"Unable to find service definition file from location '%s'. Error: %s",
				pwdAbsolute, err.Error())

		}

		c.Ui.Error(msg)

		return 1
	}

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
