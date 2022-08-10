package cmd

import (
	"journal/auth"
	"github.com/spf13/cobra"
)

var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		add, _ := cmd.Flags().GetString("add")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		auth.Login(false, add, username, password)
	},
}

func init() {
	rootCmd.AddCommand(entryCmd)
	entryCmd.Flags().StringP("add", "a", "", "Add")
	entryCmd.Flags().StringP("username", "u", "", "Username")
	entryCmd.Flags().StringP("password", "p", "", "Password")
}