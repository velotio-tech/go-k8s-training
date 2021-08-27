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
	"time"

	"github.com/spf13/cobra"
	"github.com/thisisprasad/go-k8s-training/Prasad/golang/assignment2/app"
	"github.com/thisisprasad/go-k8s-training/Prasad/golang/assignment2/journal"
	"github.com/thisisprasad/go-k8s-training/Prasad/golang/assignment2/users"
)

// entryCmd represents the entry command
var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "Add an entry in journal",
	Long:  `Adds an entry in journal of the logged-in user.`,
	Run: func(cmd *cobra.Command, args []string) {
		entry, _ := cmd.Flags().GetString("add")
		uname, _ := cmd.Flags().GetString("user")
		passwd, _ := cmd.Flags().GetString("passwd")
		if users.ValidateUser(uname, passwd) {
			app.LoginUser(uname)
			je := journal.JournalEntry{
				Timestamp: time.Now().Format("02 Jan 2006 15:04"),
				Text:      entry,
			}
			journal.AddEntry(je)
			app.CloseJournApp()
		}
	},
}

func init() {
	rootCmd.AddCommand(entryCmd)
	entryCmd.Flags().StringP("add", "a", "", "Add an entry")
	entryCmd.Flags().StringP("user", "u", "", "username")
	entryCmd.Flags().StringP("passwd", "p", "", "password")
	entryCmd.MarkFlagRequired("add")
	entryCmd.MarkFlagRequired("user")
	entryCmd.MarkFlagRequired("passwd")
}
