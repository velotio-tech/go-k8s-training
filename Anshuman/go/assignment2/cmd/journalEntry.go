/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"assignment2/controller"
	"assignment2/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var username, password, entry *string

// journalEntryCmd represents the journalEntry command
var journalEntryCmd = &cobra.Command{
	Use:   "journalEntry",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if *username != "" && *password != "" && *entry != "" {
			if !utils.CheckApplicationExists() {
				utils.SetupApplication()
			}
			controller.SaveSingleEntry(username, password, entry)
		} else {
			fmt.Println("Please enter valid credentials to post your entry.")
		}
	},
}

func init() {
	rootCmd.AddCommand(journalEntryCmd)
	username = journalEntryCmd.Flags().StringP("username", "u", "", "your account username")
	password = journalEntryCmd.Flags().StringP("password", "p", "", "your account password")
	entry = journalEntryCmd.Flags().StringP("entry", "e", "", "your journal entry")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// journalEntryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// journalEntryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
