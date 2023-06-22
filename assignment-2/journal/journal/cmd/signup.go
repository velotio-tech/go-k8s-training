/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/pratikpjain/journal/utils"
	"github.com/spf13/cobra"
)

var username, password string

// signupCmd represents the signup command
var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "signup command",
	Long: `This is the signup command used for creating a new account. Maximum of 10 accounts creation are allowed, if exceeded an error will be returned.
	
	Username should be of length 4 to 10 characters and should not contain '#' & '~'
	Password should be of length 4 to 10 characters and should not contain '#' & '~'
	
	Command: journal signup --username <username> --password <password>`,

	Run: func(cmd *cobra.Command, args []string) {

		if username == "" && password == "" {
			log.Fatal(errors.New("please enter username and password"))
		}
		if err := utils.LoadJournalData(); err != nil {
			log.Fatal(err.Error())
		}
		if err := utils.LoadAuthData(); err != nil {
			log.Fatal(err.Error())
		}
		err := utils.AddNewUser(username, password)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Added new user")
		if err := utils.StartUserPrompt(username); err != nil {
			log.Fatal(err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(signupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	signupCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Username (having length between 4-10 and not containing `#` & `~`)")
	signupCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password (having length between 4-10 and not containing `#` & `~`)")

	signupCmd.MarkFlagRequired("username")
	signupCmd.MarkFlagRequired("password")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
