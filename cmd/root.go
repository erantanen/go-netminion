package cmd

import "github.com/spf13/cobra"

const (
	// AppVersion is the application version
	AppVersion = "0.0.1"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

// RootCmd is the entry point for the application from which all actions are subcommands.
var RootCmd = &cobra.Command{
	Use:   "go-netminion",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
