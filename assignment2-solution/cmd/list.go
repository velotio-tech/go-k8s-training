package cmd

import (
	"journal/app"

	"github.com/spf13/cobra"
)

var uname string

var listCmd = &cobra.Command{
	Use:   "list command",
	Short: "List all entries",
	Long:  "List all existing entries to user",

	// command: journal list --username piyu
	Run: func(cmd *cobra.Command, args []string) {
		//list the entries
		app.DecryptFile()
		app.GetDataFromFile()
	},
}

func init() {
	cmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringVarP(&uname, "uname", "n", "", "Enter username")
	listCmd.MarkPersistentFlagRequired(uname)
}
