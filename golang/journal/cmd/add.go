package cmd

import (
	"fmt"

	"journal/manager"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add user entries to the journal",
	Long:  `This command helps user to add entries to the journal.`,
	Run: func(cmd *cobra.Command, args []string) {
		manager.AddEntry(Message)
		fmt.Println("Entry added successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&Message, "message", "m", "", "Entry of the user (required)")
	addCmd.MarkFlagRequired("message")
}
