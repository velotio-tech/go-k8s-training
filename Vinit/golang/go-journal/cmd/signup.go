// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-journal/users"
)
var username, pass string

// signupCmd represents the signup command
var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "Command will help users sign up to the go-journal",
	Long: `example command:
go-journal signup -u <Desired Username> -p <Desired Password>

This command will help you create an account to access all the features of the go-journal application`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ = cmd.Flags().GetString("username")
		pass, _ = cmd.Flags().GetString("password")
		fmt.Printf("Creating a journal for %s", username)
		_ = createJournal(username, pass)
	},
}

func init() {
	rootCmd.AddCommand(signupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//signupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	signupCmd.Flags().StringP("username", "u", "", "username of the users")
	signupCmd.Flags().StringP("password", "p", "", "Password from the users")
	signupCmd.MarkFlagRequired("username")
	signupCmd.MarkFlagRequired("password")
}

func createJournal(username, pass string) *users.User {
	if users.AlreadyExists(username, pass, "signup") {
		fmt.Println("User with this username already exists, please login or create a user with different username")
		return new(users.User)
	} else {
		newUser := users.CreateNew(username, pass)
		fmt.Printf("User with Username: %s Created Successfully \n", username)
		return newUser
	}
}