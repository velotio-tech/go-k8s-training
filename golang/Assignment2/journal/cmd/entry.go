package cmd

import (
	"fmt"
	"os"

	"github.com/jshiwam/journal/pkg"
	"github.com/spf13/cobra"
)

var add string

var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "Manage(Add/List) user entries",
	Long: `
  This command helps user add/list their entries to the journal.
  Usage: journal entry --add <message> --email=<user email> 
		 journal entry list --email=<user email>
  `,
	Run: func(cmd *cobra.Command, args []string) {
		err := manager.LoadRegisteredUsers()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		user := manager.GetUser(email)
		if user == nil {
			fmt.Println("User doesnot exist, signup before adding new entry")
			os.Exit(1)
		}
		entry := pkg.NewEntry(add)
		if !entry.IsValid() {
			fmt.Println("Invalid entry, message can't be empty, please add again")
			os.Exit(1)
		}

		user.AddEntry(entry)
		err = manager.Commit()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("Entry  successfully added to user ", email)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List user entries",
	Long: `
  This sub command helps user list their entries to the journal.
  Usage: journal entry list --email=<user email>
  `,
	Run: func(cmd *cobra.Command, args []string) {
		err := manager.LoadRegisteredUsers()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		user := manager.GetUser(email)
		if user == nil {
			fmt.Println("User doesnot exist, signup before adding new entry")
			os.Exit(1)
		}
		for _, entry := range user.Journal {
			fmt.Println(entry.CreatedAt, entry.Message)
		}
	},
}

func init() {
	rootCmd.AddCommand(entryCmd)
	entryCmd.AddCommand(listCmd)
	entryCmd.Flags().StringVarP(&add, "add", "a", "", "Journal Message (required)")
	entryCmd.Flags().StringVarP(&email, "email", "e", "", "User Email (required)")
	listCmd.Flags().StringVarP(&email, "email", "e", "", "User Email (required)")
	listCmd.MarkFlagRequired("email")
	entryCmd.MarkFlagRequired("email")
	entryCmd.MarkFlagRequired("add")
}
