/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"journal/users"
	"log"

	"github.com/spf13/cobra"
)

// signupCmd represents the signup command
var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "Create a new account.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example: journal signup -u <USERNAME> -p <PASSWORD>

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("signup called")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		signup(username, password)
	},
}

func init() {
	rootCmd.AddCommand(signupCmd)
	signupCmd.Flags().StringP("username", "u", "", "enter a username for user")
	signupCmd.Flags().StringP("password", "p", "", "enter a password for user")
	signupCmd.MarkFlagRequired("username")
	signupCmd.MarkFlagRequired("password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func signup(username, password string) *users.User {
	if users.Exists(username) {
		fmt.Println("User already exists. Login or Use different username")
		return new(users.User)
	} else {
		newUser, err := users.Create(username, password)
		if err == nil {
			fmt.Printf("User account for username : %s created! \n", username)
			return newUser
		} else {
			log.Fatalf("Error : %s", err)
			return newUser
		}
	}
}
