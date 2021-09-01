// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-journal/users"
)


// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ = cmd.Flags().GetString("username")
		pass, _ = cmd.Flags().GetString("password")
		fmt.Printf("Creating a journal for %s", username)
		loggedInUser := AuthUser(username, pass)
		if loggedInUser != nil{
			GetStarted(loggedInUser)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringP("username", "u", "", "username of the users")
	loginCmd.Flags().StringP("password", "p", "", "Password from the users")
	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("password")
}

func AuthUser(username, pass string) *users.User {
	if ! users.AlreadyExists(username, pass, "login") {
		fmt.Println("User with this username doesn't exists, please signup")
		return new(users.User)
	} else {
		newUser := users.GetData(username)
		fmt.Printf("User with Username: %s logged In Successfully \n", username)
		return newUser
	}
}