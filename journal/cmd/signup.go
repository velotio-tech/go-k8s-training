package cmd

import (
	"journal/auth"
	"github.com/spf13/cobra"
)

var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		auth.Register(username, password)
	},
}

func init() {
	rootCmd.AddCommand(signupCmd)
	signupCmd.Flags().StringP("username", "u", "", "Username")
	signupCmd.Flags().StringP("password", "p", "", "Password")
}