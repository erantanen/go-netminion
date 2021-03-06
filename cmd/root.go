package cmd

import "github.com/spf13/cobra"

const (
	// AppVersion is the application version
	AppVersion = "0.0.1"
)

var (
	// portExpressions represents ranges or inclusive ports
	//
	// For example:
	//  0-2000     is from port 0 through 2000
	//  22         is just port 22
	//  1-100,8080 is ports 1 through 100 and port 8080
	portExpressions string

	// timeout per network operation.
	//
	// If an integer then it is assumed to be in milliseconds,
	// but if value includes a known unit then it will be
	// used instead.
	timeout string
)

func init() {
	scanCmd.Flags().StringVarP(&portExpressions, "port", "p", "", "Inclusive ports to scan")
	scanCmd.Flags().StringVarP(&timeout, "timeout", "t", "50ms", "Timeout in milliseconds")

	RootCmd.AddCommand(scanCmd)
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
