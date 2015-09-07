package command

import (
	"fmt"
	"github.com/latticework/proto-goquiles-server-todo/goquilesctl/api"
	"strings"
)

type ServerCommand struct {
	Meta
}

func (c *ServerCommand) Run(args []string) int {
	var requestUrl, requestJson string

	flags := c.Meta.FlagSet("server", FlagSetDefault)
	flags.Usage = func() { c.Ui.Error(c.Help()) }
	//	flags.StringVar(&requestUrl, "message", "", "")
	//	flags.StringVar(&requestJson, "message", "", "")

	if err := flags.Parse(args); err != nil {
		return 1
	}

	client, err := c.Client()
	if err != nil {
		c.Ui.Error(fmt.Sprintf(
			"Error initializing client: %s", err))
		return 1
	}

	request := api.SendRequest{
		RequestURL:  requestUrl,
		RequestJSON: requestJson,
	}

	response, err := client.Send(&request)
	if err != nil {
		c.Ui.Error(fmt.Sprintf(
			"Error sending message: %s", err))

		return 1
	}

	c.Ui.Output(fmt.Sprintf("Response for message to '%s':\n%s",
		request.RequestURL, response.ResponseJSON))

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
