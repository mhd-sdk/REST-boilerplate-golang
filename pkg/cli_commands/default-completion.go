package cli_commands

import "github.com/spf13/cobra"

func completionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Generate the autocompletion script for the specified shell",
	}
}

func init() {
	completion := completionCommand()

	// mark completion hidden
	completion.Hidden = true
	rootCmd.AddCommand(completion)

	// other subcommands
	rootCmd.AddCommand(completion)
}
