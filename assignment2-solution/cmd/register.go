package cmd

import (
	"journal/app"

	"github.com/spf13/cobra"
)

var user, pass string

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register into the journal app",
	Long:  "To register, please provide username and password",

	Run: func(cmd *cobra.Command, args []string) {
		// register
		// add entry to the map
		app.AddNewUserToMap(user)
		app.AddNewEntryToFile(user)
		app.EncryptFile()
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.PersistentFlags().StringVarP(&user, "uname", "u", "", "Enter username")
	registerCmd.PersistentFlags().StringVarP(&pass, "passwd", "p", "", "Enter password")

	registerCmd.MarkFlagRequired(user)
	registerCmd.MarkFlagRequired(pass)
}
