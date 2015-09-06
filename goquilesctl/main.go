package main

import (
	"github.com/latticework/proto-goquiles-server-todo/goquilesctl/cli"
	"os"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
