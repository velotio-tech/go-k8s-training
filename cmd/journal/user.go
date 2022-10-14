package journal

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/velotio-tech/go-k8s-training/pkg/user"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login - allows user to log in to his journal",
	Long:  "login - allows user to log in to his journal using his username and password",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		user := user.InitializeUser(username, password)
		err := user.Login(false)
		if err != nil {
			fmt.Fprintln(os.Stderr, "login failed :", err.Error())
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, "Login Successful")
	},
}

var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "signup - allows user to create a new account",
	Long:  "signup - allows user to create a new account to his journal using his username and password",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		user := user.InitializeUser(username, password)
		err := user.Create()
		if err != nil {
			fmt.Fprintln(os.Stderr, "signup failed :", err.Error())
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, "User created")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(signupCmd)
}
