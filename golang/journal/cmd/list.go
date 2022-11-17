package cmd

import (
	"github.com/spf13/cobra"
	"github.com/velotio-ajaykumbhar/journal/app/service"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list of journal entry",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		service.GetAllEntry()
	},
}

func init() {
	entryCmd.AddCommand(listCmd)
}
