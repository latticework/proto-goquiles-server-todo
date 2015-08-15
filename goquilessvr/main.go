package main

import (
	"log"
	"os"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("goquilescvr", "0.0.1")

	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"start": startCommandFactory,
		"stop": stopCommandFactory,
		"send": sendCommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nul {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
