package cmd

import (
	"fmt"
	"journal/app"

	"github.com/spf13/cobra"
)

//var user, password string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to journal app",
	Long:  "Please provide username and password to login into the journal app",

	// command looks like -> journal login --username piyu --password piyu123
	Run: func(cmd *cobra.Command, args []string) {
		// check authetication of user
		// if authorized then load file data if user perform list
		app.DecryptFile()
		app.GetDataFromFile()

		isValid := app.ValidateUser(user)
		if isValid {
			fmt.Println("Login Successful!")
		} else {
			fmt.Println("Sorry, Please enter correct username and password!")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.PersistentFlags().StringVarP(&user, "uname", "u", "", "Enter username")
	loginCmd.PersistentFlags().StringVarP(&pass, "passwd", "p", "", "Enter password")

	loginCmd.MarkFlagRequired(user)
	loginCmd.MarkFlagRequired(pass)
}
