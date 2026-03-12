package sample

import (
	"fmt"

	"github.com/StevenACoffman/command-pattern-go/pkg/pattern/command"
)

func SampleCommand() *command.Command {
	cmd := &command.Command{
		UsageLine: "c1",
		Short:     "command1",
		Long:      "Sample command, the first one",
		Run:       sampleCommand,
	}

	return cmd
}

func sampleCommand(cmd *command.Command, args []string) error {
	fmt.Println("Hello from command1")
	return nil
}
