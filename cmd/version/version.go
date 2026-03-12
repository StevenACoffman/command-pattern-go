package version

import (
	"github.com/StevenACoffman/command-pattern-go/pkg/pattern/command"
	"github.com/StevenACoffman/command-pattern-go/pkg/version"
)

func VersionCommand() *command.Command {
	cmd := &command.Command{
		UsageLine: "version",
		Short:     "Shows version",
		Long:      "Shows the version of this binary",
		Run:       versionCmd,
	}

	return cmd
}

func versionCmd(cmd *command.Command, args []string) error {
	version.GetVersionInfo()
	return nil
}
