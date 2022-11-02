package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Entry - Journal app",
	Short: "Journal application",
	Long:  "Journal application",

	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		fmt.Println("Error occured", err)
		os.Exit(1)
	}
}

func start() {
	fmt.Println("Welcome to Personal Journal app")
	fmt.Println("-------------------------------")
	fmt.Println("1. Register 	 : register --uname <username> --passwd <password>")
	fmt.Println("2. Login         : login --uname <username> --passwd <password>")
	fmt.Println("3. Add new entry : entry --add <data> --uname <username> --passwd <password>")
	fmt.Println("4. List entries  : entry list --uname <username>")
}
