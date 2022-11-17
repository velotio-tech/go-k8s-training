package cmd

import (
	"github.com/spf13/cobra"
	"github.com/velotio-ajaykumbhar/journal/app/service"
)

var add string

var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "add journal entry.",
	Long: `provide the task details to make entry in your journal.
	please note that only 50 entry will be stored in our system. old entry automatically deleted`,
	Example: `entry --add "this will be your entry"`,
	Run: func(cmd *cobra.Command, args []string) {
		service.CreateEntry(add)
	},
}

func init() {
	entryCmd.PersistentFlags().StringVarP(&add, "add", "a", "", "provide entry")

	rootCmd.AddCommand(entryCmd)
}
