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
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/thisisprasad/go-k8s-training/Prasad/golang/assignment2/app"
	"github.com/thisisprasad/go-k8s-training/Prasad/golang/assignment2/users"
	"golang.org/x/term"
)

// signupCmd represents the signup command
var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		uname := getNewUserName()
		passw, err := getNewUserPassw()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if users.CreateNewUser(uname, passw) {
			fmt.Println("User created successfully...")
			app.LoginUser(uname)
			app.LaunchCMDclient(uname)
		} else {
			fmt.Println("Error occured while creating user.")
		}
	},
}

func checkUserLimit() {

}

func getNewUserName() string {
	reader := bufio.NewReader(os.Stdin)
	var uname string = ""
	for {
		fmt.Printf("Enter username: ")
		uname, _ = reader.ReadString('\n')
		uname = strings.Replace(uname, "\n", "", -1)
		if users.UserAlreadyExists(uname) {
			fmt.Printf("username '" + uname + "' already exists. PLease try again.\n")
		} else {
			break
		}
	}

	return uname
}

//	gets the password for new user from console
func getNewUserPassw() (string, error) {
	var passw string = ""
	var recheckPassw string = ""
	fmt.Printf("Enter password: ")
	bytePassw, _ := term.ReadPassword(int(syscall.Stdin))
	passw = strings.Replace(string(bytePassw), "\n", "", -1)
	fmt.Println()
	fmt.Printf("Enter password again: ")
	recheckBytePassw, _ := term.ReadPassword(int(syscall.Stdin))
	recheckPassw = strings.Replace(string(recheckBytePassw), "\n", "", -1)
	fmt.Println()

	if passw != recheckPassw {
		fmt.Printf("Passwords do not match.\nPlease enter password again: ")
		recheckBytePassw, _ = term.ReadPassword(int(syscall.Stdin))
		recheckPassw = strings.Replace(string(recheckBytePassw), "\n", "", -1)
		fmt.Println()
		if passw != recheckPassw {
			return "", errors.New("Password do not match.")
		}
	}

	return passw, nil
}

func init() {
	rootCmd.AddCommand(signupCmd)
}
