/*
Copyright © 2022 Parav Kaushal <paravkaushal.kv1@gmail.com>

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

	"github.com/paravkaushal/go-journal/pkg"
	"github.com/spf13/cobra"
)

var username, entryData, passwd *string

// entryCmd represents the entry command
var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "You can insert your entry using flags",
	Long:  `PS: You must be a registered user to perform this action`,
	Run: func(cmd *cobra.Command, args []string) {

		isAuthenticated := pkg.AuthenticateUser(*username, *passwd)
		if isAuthenticated {
			pkg.CreateNewEntry(*username, *entryData)
		} else {
			fmt.Println("Err!! Either your username/password is incorrect or you are not a regisered user.")
		}
	},
}

func init() {
	rootCmd.AddCommand(entryCmd)
	entryData = entryCmd.PersistentFlags().String("add", "", "A help for foo")
	username = entryCmd.PersistentFlags().String("user", "", "A help for foo")
	passwd = entryCmd.PersistentFlags().String("passwd", "", "A help for foo")
}
