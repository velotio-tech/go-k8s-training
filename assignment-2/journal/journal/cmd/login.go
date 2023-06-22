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

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login command",
	Long: `this is the login command, can be used if you already have an account.
	
	Command: journal login --username <user_name> --password <password>`,

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
		err := utils.ValidateUser(username, password)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Successfully Logged in ...")
		if err := utils.StartUserPrompt(username); err != nil {
			log.Fatal(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	loginCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Username (having length between 4-10 and not containing `#` & `~`)")
	loginCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password (having length between 4-10 and not containing `#` & `~`)")

	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("password")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
