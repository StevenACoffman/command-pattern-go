package cmd

import (
	"errors"
	"fmt"

	"github.com/StevenACoffman/command-pattern-go/cmd/sample"
	"github.com/StevenACoffman/command-pattern-go/cmd/version"
	"github.com/StevenACoffman/command-pattern-go/pkg/pattern/command"
)

// Run takes Arguments *without* the executable name
// and runs the command
// This is where new commands are registered and dispatched
func Run(args []string) error {
	commands := []*command.Command{
		version.VersionCommand(),
		sample.SampleCommand(),
	}

	m := make(map[string]*command.Command)
	for i := range commands {
		cmd := commands[i]
		m[cmd.UsageLine] = cmd
	}

	// if they don't pass a command or only pass "help"
	if len(args) == 0 || (len(args) == 1 && args[0] == "help") {
		fmt.Println(
			"Available Commands:",
		)
		for i := range commands {
			cmd := *commands[i]
			fmt.Printf("%s - %s\n", cmd.UsageLine, cmd.Short)
		}
		return nil
	}

	arg := args[0]
	cmd := m[arg]
	if cmd == nil {
		return errors.New(arg + ": invalid command")
	}
	// pass arguments without the executable and without the command itself
	return cmd.Run(cmd, args[1:])
}
