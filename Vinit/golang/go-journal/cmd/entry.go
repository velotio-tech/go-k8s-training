// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// entryCmd represents the entry command
var entry string
var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ = cmd.Flags().GetString("username")
		pass, _ = cmd.Flags().GetString("password")
		entry, _ = cmd.Flags().GetString("entry")
		fmt.Printf("Authenticating User and Adding Entry")
		addEntry(username,pass, entry)
	},
}

func init() {
	rootCmd.AddCommand(entryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// entryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// entryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	entryCmd.Flags().StringP("username", "u", "", "username of the users")
	entryCmd.Flags().StringP("password", "p", "", "Password from the users")
	entryCmd.Flags().StringP("entry", "e", "", "Entry for the journal")
	entryCmd.MarkFlagRequired("username")
	entryCmd.MarkFlagRequired("password")
	entryCmd.MarkFlagRequired("entry")
}

func addEntry(usr, pwd, entry string) {
	var loggedInUser = AuthUser(usr, pwd)
	if loggedInUser != nil {
		loggedInUser.WriteJournal(entry)
	} else {
		fmt.Println("User Creds Incorrect \nCannot add the entry")
	}
}
