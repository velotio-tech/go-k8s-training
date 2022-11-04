package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "journal [command]",
	Short: "An application that helps manage accounts of users",
	Long: `
  This is a CLI that enables users to manage their accounts.
  You would be able to add credit transactions and debit transactions to various users.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
