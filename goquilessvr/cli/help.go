package cli

import (
	"bytes"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/mitchellh/cli"
)

// HelpFunc is a cli.HelpFunc that can is used to output the help for goquilessvr.
func HelpFunc(commands map[string]cli.CommandFactory) string {
	commonNames := map[string]struct {}{
		"start": struct {}{},
		"stop": struct {}{},
		"send": struct {}{},
	}

	// Determine the maximum key length, and classify based on type
	commonCommands := make(map[string]cli.CommandFactory)
	otherCommands := make(map[string]cli.CommandFactory)
	maxKeyLen := 0
	for key, f := range commands {
		if len(key) > maxKeyLen {
			maxKeyLen = len(key)
		}

		if _, ok := commonNames[key]; ok {
			commonCommands[key] = f
		} else {
			otherCommands[key] = f
		}
	}

	var buf bytes.Buffer
	buf.WriteString("usage: goquilessvr [-version] [-help] <command> [args]\n\n")
	buf.WriteString("Common commands:\n")
	buf.WriteString(listCommands(commonCommands, maxKeyLen))
	buf.WriteString("\nAll other commands:\n")
	buf.WriteString(listCommands(otherCommands, maxKeyLen))
	return buf.String()
	return nil;
}

// listCommands just lists the commands in the map with the
// given maximum key length.
func listCommands(commands map[string]cli.CommandFactory, maxKeyLen int) string {
	var buf bytes.Buffer

	// Get the list of keys so we can sort them, and also get the maximum
	// key length so they can be aligned properly.
	keys := make([]string, 0, len(commands))
	for key, _ := range commands {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		commandFunc, ok := commands[key]
		if !ok {
			// This should never happen since we JUST built the list of
			// keys.
			panic("command not found: " + key)
		}

		command, err := commandFunc()
		if err != nil {
			log.Printf("[ERR] cli: Command '%s' failed to load: %s",
				key, err)
			continue
		}

		key = fmt.Sprintf("%s%s", key, strings.Repeat(" ", maxKeyLen-len(key)))
		buf.WriteString(fmt.Sprintf("    %s    %s\n", key, command.Synopsis()))
	}

	return buf.String()
}