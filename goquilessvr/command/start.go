package command

import (
	"strings"
)


type StartCommand struct {
	Meta
}

func (c* StartCommand) Run(args []string) int {
	return 0
}

func (c* StartCommand) Synopsis() string {
	return "Start a Quiles server"
}

func (c* StartCommand) Help() string {
	helpText := `
Usage: quiles start [options]

	Starts a Quiles server.

	This command connects to a Quiles server and starts it.

General Options:
  ` + generalOptionsUsage() + `
Start Options:
  -start-option=value     The start option being explained that has a
  						  default value.
`
	return strings.TrimSpace(helpText)
}