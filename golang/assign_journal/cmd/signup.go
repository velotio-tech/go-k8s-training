/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/6akshita/assign_journal/user"
	"github.com/spf13/cobra"
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
		fmt.Println("signup called")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		user.UserSignUp(username, password)
	},
}

func init() {
	rootCmd.AddCommand(signupCmd)
	signupCmd.Flags().StringP("username", "u", "", "Username")
	signupCmd.Flags().StringP("password", "p", "", "Password")
	signupCmd.MarkFlagRequired("usernane")
	signupCmd.MarkFlagRequired("password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
