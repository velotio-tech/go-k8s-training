package cmd

import (
	"journal/auth"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		auth.Login(true, "", username, password)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("username", "u", "", "Username")
	listCmd.Flags().StringP("password", "p", "", "Password")
}