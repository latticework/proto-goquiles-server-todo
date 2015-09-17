package goquilessvr

import (
	"github.com/latticework/proto-goquiles-server-todo/goquilessvr/cli"
	"os"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
