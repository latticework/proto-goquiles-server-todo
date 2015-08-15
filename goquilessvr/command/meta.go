package command

import  (
	"strings"

	"github.com/mitchellh/cli"
)

// Meta contains the meta-options and functionality that nearly every
// Quiles server command inherits.
type Meta struct {
	Ui cli.Ui
}

func (m *Meta) Client() ()

// generalOptionsUsage returns the usage documenation for commonly
// available options
func generalOptionsUsage() string {
	general := `
  -option-name=value      The Option Name that is explained and has
  					      a value.
	`
	return strings.TrimSpace(general)
}