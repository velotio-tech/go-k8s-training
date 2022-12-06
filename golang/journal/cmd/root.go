package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "journal",
		Short: "CLI application of journal",
		Long:  `A CLI application to store personal journal log with user management.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
