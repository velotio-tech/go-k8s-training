package cmd

import (
	"fmt"
	"journal/helper"
	"journal/manager"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login registered user to the journal",
	Long:  `This command helps user login to the journal.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := manager.LogIn(Email, Password)
		helper.Check(err)
		fmt.Println("Login successful")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&Email, "email", "e", "", "Email of the user (required)")
	loginCmd.Flags().StringVarP(&Password, "password", "p", "", "Password of the user (required)")
	loginCmd.MarkFlagRequired("email")
	loginCmd.MarkFlagRequired("password")
}
