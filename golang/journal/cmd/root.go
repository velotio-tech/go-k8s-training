package cmd

import (
	"journal/manager"

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
	manager.PrintStruct()
	return rootCmd.Execute()
}
