package command

import (
	"fmt"
	"github.com/latticework/proto-goquiles-server-todo/goquilesctl/api"
	"strings"
)

type SendCommand struct {
	Meta
}

func (c *SendCommand) Run(args []string) int {
	var requestUrl, requestJson string

	flags := c.Meta.FlagSet("send", FlagSetDefault)
	flags.Usage = func() { c.Ui.Error(c.Help()) }
	flags.StringVar(&requestUrl, "message", "", "")
	flags.StringVar(&requestJson, "message", "", "")

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

func (c *SendCommand) Synopsis() string {
	return "Send a message to a Quiles server"
}

func (c *SendCommand) Help() string {
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
