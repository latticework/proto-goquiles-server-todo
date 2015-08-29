package main

import (
	"log"
	"os"
	"github.com/mitchellh/cli"
	"github.com/latticework/proto-goquiles-server-todo/goquilessvr/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))

// @kenbrubaker: not sure where I got this code...
//	c := cli.NewCLI("goquilescvr", "0.0.1")
//
//	c.Args = os.Args[1:]
//	c.Commands = map[string]cli.CommandFactory{
//		"start": startCommandFactory,
//		"stop": stopCommandFactory,
//		"send": sendCommandFactory,
//	}
//
//	exitStatus, err := c.Run()
//	if err != nul {
//		log.Println(err)
//	}
//
//	os.Exit(exitStatus)
}
