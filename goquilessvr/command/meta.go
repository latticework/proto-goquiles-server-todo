package command

import  (
	"strings"

	"github.com/mitchellh/cli"
	"github.com/latticework/proto-goquiles-server-todo/goquilessvr/api"
	"os"
	"net/http"
	"net"
	"time"
	"flag"
	"io"
	"bufio"
)

// Meta contains the meta-options and functionality that nearly every
// Quiles server command inherits.
type Meta struct {
	Ui cli.Ui

	flagAddress string
	flagServiceVersion string
}

func (m *Meta) Client() (*api.Client, error) {
	config := api.DefaultConfig()
	if address := os.Getenv(EnvQuilesServerAddress); address != "" {
		config.Address = address
	}

	if serviceVersion := os.Getenv(EnvQuilesServiceVersion); serviceVersion != "" {
		config.
	}

	if m.flagAddress != "" {
		config.Address = m.flagAddress
	}

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

type FlagSetFlags uint

const (
	FlagSetNone FlagSetFlags = 0
	FlagSetServer FlagSetFlags = 1 << iota
	FlagSetDefault = FlagSetServer
)

func (m *Meta) FlagSet(n string, fs FlagSetFlags) *flag.FlagSet {
	f := flag.NewFlagSet(n, flag.ContinueOnError)

	if fs & FlagSetServer != 0 {
		f.StringVar(&m.flagAddress, "address", "", "")
	}

	// Create an io.Writer that writes to our Ui properly for errors.
	// This is kind of a hack, but it does the job. Basically: create
	// a pipe, use a scanner to break it into lines, and output each line
	// to the UI. Do this forever.
	errR, errW := io.Pipe()
	errScanner := bufio.NewScanner(errR)
	go func() {
		for errScanner.Scan() {
			m.Ui.Error(errScanner.Text())
		}
	}()
	f.SetOutput(errW)

	return f
}

// generalOptionsUsage returns the usage documenation for commonly
// available options
func generalOptionsUsage() string {
	general := `
  -option-name=value      The Option Name that is explained and has
  					      a value.
	`
	return strings.TrimSpace(general)
}