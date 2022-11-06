package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout registered user from the journal",
	Long: `
  This command helps user logout from the journal.
  Usage: journal logout --email=<user email>`,
	Run: func(cmd *cobra.Command, args []string) {
		err := manager.Logout(email)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Logout  Successful ", email)

	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
	logoutCmd.Flags().StringVarP(&email, "email", "e", "", "Email of the user (required)")
	logoutCmd.MarkFlagRequired("email")

}
