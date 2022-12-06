package cmd

import (
	"fmt"
	"journal/helper"
	"journal/manager"

	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout registered user to the journal",
	Long:  `This command helps user logout to the journal.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := manager.LogOut(Email)
		helper.Check(err)
		fmt.Println("Logout successful")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
	logoutCmd.Flags().StringVarP(&Email, "email", "e", "", "Email of the user (required)")
	logoutCmd.MarkFlagRequired("email")
}
