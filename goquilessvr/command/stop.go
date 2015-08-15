package command

import (
	"strings"
)


type StopCommand struct {
	Meta
}

func (c* StopCommand) Run(args []string) int {
	return 0
}

func (c* StopCommand) Synopsis() string {
	return "Start a Quiles server"
}

func (c* StopCommand) Help() string {
	helpText := `
Usage: quiles start [options]

	Stops a Quiles server.

	This command connects to a Quiles server and stops it.

General Options:
  ` + generalOptionsUsage() + `
Stop Options:
  -stop-option=value      The stop option being explained that has a
  						  default value.
`
	return strings.TrimSpace(helpText)
}