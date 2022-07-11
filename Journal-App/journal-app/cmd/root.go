/*
Copyright © 2022 Mugdha Watve <watve.mugdha98@gmail.com>
*/
package cmd

import (
	"fmt"
	"journal-app/journal-app/service"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "journal-app",
	Short: "A brief description of your application",
	Long:  "A Journal App",
	Run: func(cmd *cobra.Command, args []string) {
		StartApp()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func StartApp() {
	for {
		fmt.Println("Welcome to Journal App\n1.Login\n2.Sign up\n3.Exit\nEnter Your Choice: ")
		var choice string
		fmt.Scanln(&choice)
		if choice == "1" {
			service.LoginUser("", "")
		} else if choice == "2" {
			service.SignUpUser()
		} else if choice == "3" {
			break
		} else {
			fmt.Println("Please enter correct option")
		}
	}
}
