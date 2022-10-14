package journal

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/velotio-tech/go-k8s-training/pkg/user"
)

var rootCmd = &cobra.Command{
	Use:   "journal",
	Short: "journal - a simple CLI to manage journal entries",
	Long:  "journal is a super fancy multi-user CLI which manages user's journal entries",
	PreRun: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		if cmd.Flag("username").Changed {
			user := user.InitializeUser(username, password)
			err := user.Login(true)
			if err != nil {
				fmt.Fprintln(os.Stderr, "login failed :", err.Error())
				os.Exit(1)
			}
			fmt.Fprintln(os.Stdout, "Login Successful")
			return
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	InitConfig()
	rootCmd.PersistentFlags().StringP("username", "u", "", "user's name")
	rootCmd.PersistentFlags().StringP("password", "p", "", "user's password")
	rootCmd.MarkFlagsRequiredTogether("username", "password")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI\n")
		fmt.Fprintf(os.Stderr, "Description : %s\n", err.Error())
		os.Exit(1)
	}

}
