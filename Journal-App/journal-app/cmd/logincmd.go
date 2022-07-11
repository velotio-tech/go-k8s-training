package cmd

import (
	"journal-app/journal-app/service"

	"github.com/spf13/cobra"
)

var logincmd = &cobra.Command{
	Use:   "Login",
	Short: "Enter username and password to login to the application",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		service.LoginUser(username, password)
	},
}

func init() {
	rootCmd.AddCommand(logincmd)
	logincmd.Flags().StringP("username", "u", "", "Username")
	logincmd.Flags().StringP("password", "p", "", "Password")
	logincmd.MarkFlagRequired("usernane")
	logincmd.MarkFlagRequired("password")
}
