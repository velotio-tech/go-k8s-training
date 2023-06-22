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

var entry string

// entryCmd represents the entry command
var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "entry command",
	Long: `this is entry command. This is a wildcard command if you want to make an entry to journal and login at the same time.
	
	
	Command: journal entry --add <your_entry> --username <user_name> --password <password>`,
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
		if err := utils.AddNewEntry(username, entry); err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Successfully created the entry ... ")
	},
}

func init() {
	rootCmd.AddCommand(entryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// entryCmd.PersistentFlags().String("foo", "", "A help for foo")

	entryCmd.PersistentFlags().StringVarP(&entry, "add", "a", "", "Entry (non-empty and should not contain `#` & `~`)")
	entryCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Username (having length between 4-10 and not containing `#` & `~`)")
	entryCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password (having length between 4-10 and not containing `#` & `~`)")

	entryCmd.MarkFlagRequired(entry)
	entryCmd.MarkFlagRequired(username)
	entryCmd.MarkFlagRequired(password)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// entryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
