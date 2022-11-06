package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login registered user to the journal",
	Long: `
  This command helps user login to the journal.
  Usage: journal login --email=<user email> --passwd=<user password> `,
	Run: func(cmd *cobra.Command, args []string) {
		err := manager.Login(email, password)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Login  Successful ", email)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&email, "email", "e", "", "Email of the user (required)")
	loginCmd.Flags().StringVarP(&password, "passwd", "p", "", "Password for the user (required)")
	loginCmd.MarkFlagRequired("email")
	loginCmd.MarkFlagRequired("passwd")
}
