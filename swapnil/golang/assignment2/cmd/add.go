/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"journalhelpers"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add entry to your journal",
	Long: `add entry to your journal. For example: $myjournal add --username <username> --password <password> --entry "<some long text>"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		journalhelpers.AddEntry(username1, password1, entry)
		os.Exit(0)
	},
}

var username1 string
var password1 string
var entry string

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVar(&username1, "username", "", "username")
	addCmd.MarkFlagRequired("username")
	addCmd.Flags().StringVar(&password1, "password", "", "password")
	addCmd.MarkFlagRequired("password")
	addCmd.Flags().StringVar(&entry, "entry", "", "entry text")
	addCmd.MarkFlagRequired("entry")
}
