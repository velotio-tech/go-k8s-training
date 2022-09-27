package cmd

import (
	"journal/database"
	"fmt"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register new user",
	Long: `This command register a new user to the journal CLI app.
For example:
./journal register --username=<username> --name=<name>
./journal register -u=<username> -n=<name>
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Registering you in the system. Should be done really quick....")
		_, err := database.CreateUser(username, name)
		if err == nil {
			fmt.Println("Regsitered successfully!")
		} else {
			fmt.Println("Failed to regsitered '%s'", err)
		}
	},
}

var username string
var name string

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&username, "username", "u", "", "Username for the user")
	registerCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the user")

	registerCmd.MarkFlagRequired("username")	
	registerCmd.MarkFlagRequired("name")
}
