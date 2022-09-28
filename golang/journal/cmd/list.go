package cmd

import (
	"fmt"
	"journal/database"
	"github.com/spf13/cobra"
	"time"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		entries := database.FindEntriesForUser(username)

		for index, entry := range(entries) {
			fmt.Println("--------------------")
			fmt.Println("Entry ", index+1, ": ")
			fmt.Println(entry.InsertedAt.Format(time.RFC822), " - ", entry.Content)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&username, "username", "u", "", "Username for the user")
}
