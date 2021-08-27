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
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/thisisprasad/go-k8s-training/Prasad/golang/assignment2/app"
	"github.com/thisisprasad/go-k8s-training/Prasad/golang/assignment2/users"

	"golang.org/x/term"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A command to login into the personal journal management aap.",
	Long: `A command to login into the personal journal management app. For example: assignment2 login -u prasad

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		uname, _ := cmd.Flags().GetString("uname")
		var passw string = ""
		if uname != "" {
			fmt.Println("User:", uname)
			passw = getUserPassword(uname)
			fmt.Println("password:", passw)
		} else {
			uname = getUserName()
			passw = getUserPassword(uname)
			fmt.Println(passw)
		}

		if validateLoginCredentials(uname, passw) {
			app.LoginUser(uname)
			app.LaunchCMDclient(uname)
		} else {
			fmt.Println("Error: invalid username or password.")
		}
	},
}

//	gets username from console
func getUserName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter user-name: ")
	uname, _ := reader.ReadString('\n')
	uname = strings.Replace(uname, "\n", "", -1)
	return uname
}

//	gets password from console
func getUserPassword(user string) string {
	fmt.Printf("password for user [" + user + "]: ")
	bytePassw, _ := term.ReadPassword(int(syscall.Stdin))
	return strings.Replace(string(bytePassw), "\n", "", -1)
}

//	Validates the login credentials
func validateLoginCredentials(uname string, passw string) bool {
	return users.ValidateUser(uname, passw)
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("uname", "u", "", "username for login.")
}
