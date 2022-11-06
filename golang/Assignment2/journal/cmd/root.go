package cmd

import (
	"github.com/jshiwam/journal/pkg"
	"github.com/spf13/cobra"
)

var manager = pkg.Manager{}

var rootCmd = &cobra.Command{
	Use:   "journal [command]",
	Short: "Journal CLI is an app that helps users to manage their journals",
	Long: `
  This is a CLI that enables users to manage their journals.
  You would be able to add add and list journal entries to your account.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	InitConfig()
	cobra.CheckErr(rootCmd.Execute())
}
