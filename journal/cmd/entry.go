/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

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
	"journal/utils"

	"github.com/spf13/cobra"
)

var username, password, entryData *string

// entryCmd represents the entry command
var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		isAuthenticated := utils.Authenticate(*username, *password)

		if isAuthenticated {
			utils.AddNewEntry(*username, *entryData)
		} else {
			fmt.Println("Error : Invalid Username or Password")
		}
	},
}

func init() {
	rootCmd.AddCommand(entryCmd)

	entryData = entryCmd.PersistentFlags().String("entryData", "", "Enter your journal entry")
	username = entryCmd.PersistentFlags().String("username", "", "Username")
	password = entryCmd.PersistentFlags().String("password", "", "Password")

}
