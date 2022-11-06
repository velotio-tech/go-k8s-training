package cmd

import (
	"fmt"

	"github.com/jshiwam/journal/pkg"
	"github.com/spf13/cobra"
)

var userName, password, email string

var signUpCmd = &cobra.Command{
	Use:   "signup",
	Short: "Register a new user to the journal",
	Long: `
  This command registers a new user to the journal.
  Usage: journal signup --user=<username> --passwd=<user password> --email=<user email>.`,
	Run: func(cmd *cobra.Command, args []string) {
		user := pkg.NewUser(userName, email, password)
		err := manager.Register(user)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Sign up Successful", userName)
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
