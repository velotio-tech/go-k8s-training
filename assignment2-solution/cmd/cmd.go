package cmd

import (
	"fmt"
	"journal/app"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "entry command",
	Short: "Command to create a new entry",
	Long:  "This command is used to create a new entry",

	Run: func(cmd *cobra.Command, args []string) {
		// check if the user is authorized or not
		fmt.Println(entry, " ", username, " ", password)

		app.GetDataFromFile()
		fmt.Println("Valid user!!!")
		app.AddNewEntry(username, entry)
		app.AddNewEntryToFile(username)
		app.EncryptFile()
	},
}

var entry, username, password, email string
var flag bool

func init() {
	rootCmd.AddCommand(cmd)
	cmd.PersistentFlags().StringVarP(&entry, "add", "a", "", "add a new entry")
	cmd.PersistentFlags().StringVarP(&username, "uname", "u", "", "enter username")
	cmd.PersistentFlags().StringVarP(&password, "passwd", "p", "", "enter password")

	// required flags
	cmd.MarkPersistentFlagRequired("username")
	cmd.MarkPersistentFlagRequired("password")
}
