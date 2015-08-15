package command

import (
	"strings"
)

type InitCommand struct {
	Meta
}

func (c* InitCommand) Run(args []string) int {
	return 0
}

func (c* InitCommand) Synopsis() string {
	return "Initialize a new Quiles server"
}

func (c* InitCommand) Help() string {
	helpText := `
Usage: quiles init [options]

	Initialize a new Quiles server.

	This command connects to a Quiles server and initializes it for the
	first the first time. This sets up ....

General Options:
  ` + generalOptionsUsage() + `
Init Options:
  -init-option=value      The init option being explained that has a
  						  default value.
`
	return strings.TrimSpace(helpText)
}