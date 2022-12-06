package cmd

import (
	"fmt"
	"journal/helper"
	"journal/manager"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add user entries to the journal",
	Long:  `This command helps user to add entries to the journal.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := manager.AddEntry(Message, Email)
		helper.Check(err)
		fmt.Println("Entry added successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&Message, "message", "m", "", "Entry of the user (required)")
	addCmd.Flags().StringVarP(&Email, "email", "e", "", "Email of the user (required)")
	addCmd.MarkFlagRequired("email")
	addCmd.MarkFlagRequired("message")
}
