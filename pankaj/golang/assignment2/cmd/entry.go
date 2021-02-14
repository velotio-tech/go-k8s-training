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
	"path/filepath"
	"strings"

	"github.com/pankaj9310/go-k8s-training/pankaj/golang/assignment2/journals"
	"github.com/pankaj9310/go-k8s-training/pankaj/golang/assignment2/users"
	"github.com/spf13/cobra"
)

var userPath string = "data/users.csv"

// entryCmd represents the user command
var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "Personal Journal App",
	Long: `Create a CLI application using cobra to store personal journal log with user management. For simplicity consider you could 
	enter only textual content in journal data.`,
	Run: func(cmd *cobra.Command, args []string) {
		journal, _ := cmd.PersistentFlags().GetString("add")
		user, _ := cmd.PersistentFlags().GetString("user")
		passwd, _ := cmd.PersistentFlags().GetString("passwd")
		journal = strings.TrimSpace(journal)
		passwd = strings.TrimSpace(passwd)
		user = strings.TrimSpace(user)
		if len(user) != 0 && len(passwd) != 0 {
			userJornalFile := users.GetPassword(user + passwd)
			if users.CreateUser(user, users.GetPassword(passwd)) {
				if len(journal) == 0 {
					fmt.Println("Journal entry can't be empty!")
				} else {
					jouranlsList := &journals.Journal{}
					jouranlsList.Init(50)
					userJournalPath := filepath.Join("data/journal/", userJornalFile+".txt")
					journalData := journals.ReadFile(userJournalPath, users.GetPassword(passwd))
					users.GetJournalData(jouranlsList, journalData)
					journal = journals.AddJournal(journal)
					jouranlsList.Capture(journal)
					journals.WriteFile(userJournalPath, journals.GetJournal(jouranlsList), users.GetPassword(passwd))
				}
			} else {
				fmt.Println("Invalid username or password")
			}

		} else {
			users.LoginOrRegisterUser()
		}

	},
}

func init() {
	entryCmd.PersistentFlags().String("add", "", "New journal entry")
	entryCmd.PersistentFlags().String("user", "", "Name of user")
	entryCmd.PersistentFlags().String("passwd", "", "Password")
	rootCmd.AddCommand(entryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// entryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// entryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
