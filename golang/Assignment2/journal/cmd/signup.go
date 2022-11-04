package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var userName, password, email string
var manager = Manager{}

var signUpCmd = &cobra.Command{
	Use:   "signup",
	Short: "Register a new user to the journal",
	Long: `
  This command registers a new user to the journal.
  Usage: journal signup --user=<username> --passwd=<user password> --email=<user email>.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Signing Up: ", userName, password, email)
		user := NewUser(userName, email, password)
		manager.register(user)
		manager.commit()
	},
}

func init() {
	rootCmd.AddCommand(signUpCmd)
	signUpCmd.Flags().StringVarP(&userName, "user", "u", "", "Username (required)")
	signUpCmd.Flags().StringVarP(&password, "passwd", "p", "", "Password for the user (required)")
	signUpCmd.Flags().StringVarP(&email, "email", "e", "", "Email of the user (required)")
	signUpCmd.MarkFlagRequired("user")
	signUpCmd.MarkFlagRequired("passwd")
	signUpCmd.MarkFlagRequired("email")
}
