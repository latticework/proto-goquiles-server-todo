package command

import (
	"strings"
)

type SendCommand struct {
	Meta
}

func (c* SendCommand) Run(args []string) int {
	return 0
}

func (c* SendCommand) Synopsis() string {
	return "Send a message to a Quiles server"
}

func (c* SendCommand) Help() string {
	helpText := `
Usage: quiles send [options]

	Sends a message to a Quiles server.

	This command connects to a Quiles server and sends a message.

General Options:
  ` + generalOptionsUsage() + `
Send Options:
  -send-option=value      The send option being explained that has a
  						  default value.
`
	return strings.TrimSpace(helpText)
}