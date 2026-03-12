package command

type Command struct {
	// Run runs the command. The args are the arguments after the command name.
	Run func(cmd *Command, args []string) error

	// UsageLine is the one-line usage message.
	UsageLine string

	// Short is the short description shown in the 'help' output.
	Short string

	// Long is the long message shown in the 'go help <this-command>' output.
	Long string
}
