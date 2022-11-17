package cmd

import (
	"github.com/spf13/cobra"
	"github.com/velotio-ajaykumbhar/journal/app/service"
)

var username string
var password string

var signupCmd = &cobra.Command{
	Use:                   "signup",
	Short:                 "signup in journal app",
	Long:                  `provide your credentails to create your account in system. for usages see example and help`,
	Example:               `signup --username=yourUsername --password=yourPassword`,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		service.SignUp(username, password)
	},
}

func init() {
	signupCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "provide username of your account")
	signupCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "provide password of your account")

	rootCmd.AddCommand(signupCmd)
}
