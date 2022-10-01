package cmd

import (
	"fmt"
	"journal/database"
	"github.com/spf13/cobra"
)

// entryCmd represents the entry command
var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
For Example: ./journal entry --add <journal-text>  --username <username>`,
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := database.FindUser(username)

		if user != nil {
			fmt.Println("Saving the entry....")
			_, err := database.CreateEntry(username, add)

			if err == nil {
				fmt.Println("Entry added successfully!")
			} else {
				fmt.Println("Failed to entry '%s'", err)
			}
		} else {
			fmt.Println("Looks like the user is not registered in the system!")
		}
	},
}

var add string

func init() {
	rootCmd.AddCommand(entryCmd)

	entryCmd.Flags().StringVarP(&username, "username", "u", "", "Username for the user")
	entryCmd.Flags().StringVarP(&add, "add", "a", "", "Add new journal entry")

	entryCmd.MarkFlagRequired("username")
	entryCmd.MarkFlagRequired("add")
}
