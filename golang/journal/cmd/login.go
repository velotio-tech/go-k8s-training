package cmd

import (
	"github.com/spf13/cobra"
	"github.com/velotio-ajaykumbhar/journal/app/service"
)

var loginCmd = &cobra.Command{
	Use:                   "login",
	Short:                 "login in journal app",
	Long:                  `provide your login credentails to enter into system. for usages see example and help`,
	Example:               `login --username=yourUsername --password=yourPassword`,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		service.Login(username, password)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "provide username of your account")
	loginCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "provide password of your account")
}
