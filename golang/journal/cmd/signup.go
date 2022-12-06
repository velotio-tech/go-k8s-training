package cmd

import (
	"fmt"

	"journal/helper"
	"journal/manager"

	"github.com/spf13/cobra"
)

var signCmd = &cobra.Command{
	Use:   "signup",
	Short: "Register user to the journal",
	Long:  `This command helps user signup to the journal.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := manager.CaptureNewUser(Name, Email, Password)
		helper.Check(err)
		fmt.Println("User Captured successfully")
	},
}

func init() {
	rootCmd.AddCommand(signCmd)
	signCmd.Flags().StringVarP(&Name, "name", "n", "", "Name of the user (required)")
	signCmd.Flags().StringVarP(&Email, "email", "e", "", "Email of the user (required)")
	signCmd.Flags().StringVarP(&Password, "password", "p", "", "Password of the user (required)")

	signCmd.MarkFlagRequired("name")
	signCmd.MarkFlagRequired("email")
	signCmd.MarkFlagRequired("password")

}
