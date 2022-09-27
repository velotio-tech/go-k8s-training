package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register new user",
	Long: `This command register a new user to the journal CLI app.
For example: ./journal register --username=<username> --name=<name> --password=<password>`,
	Run: func(cmd *cobra.Command, args []string) {
		if(len(args) < 1) {
			fmt.Println("Required flags not present: ", args)
		}
	},
}

var username string
var name string
var password string

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&username, "username", "u", "", "Username for the user")
	registerCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the user")
	registerCmd.Flags().StringVarP(&password, "password", "p", "", "Password for the CLI app")

	registerCmd.MarkFlagRequired("username")	
	registerCmd.MarkFlagRequired("name")
	registerCmd.MarkFlagRequired("password")
}
