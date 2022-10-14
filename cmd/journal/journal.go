package journal

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/velotio-tech/go-k8s-training/pkg/journal"
	"github.com/velotio-tech/go-k8s-training/pkg/user"
)

var entryCommand = &cobra.Command{
	Use:   "entry",
	Short: "entry - allows to manage your entries",
	Long:  "entry - allows to manage your entries",
	PreRun: func(cmd *cobra.Command, args []string) {
		parent := cmd.Parent()
		if parent != nil {
			parent.PreRun(parent, args)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		data, _ := cmd.Flags().GetString("add")
		e := journal.NewEntry(data)
		if err := e.Validate(); err != nil {
			fmt.Fprintln(os.Stderr, "entry validation failed :", err.Error())
			os.Exit(1)
		}
		u := user.GetCurrentUser()
		if u == nil {
			fmt.Fprintln(os.Stderr, "please log in")
			os.Exit(1)
		}
		j, err := u.GetJournal()
		if err != nil {
			fmt.Fprintln(os.Stderr, "unable to get journal for user :", err.Error())
			os.Exit(1)
		}
		err = j.AddEntry(e)
		if err != nil {
			fmt.Fprintln(os.Stderr, "unable to add entry to journal for user :", err.Error())
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, "Entry added successfully")
	},
}

var listEntryCommand = &cobra.Command{
	Use:   "list",
	Short: "list - displays the complete journal",
	Long:  "list - displays the complete list of journal entries",
	PreRun: func(cmd *cobra.Command, args []string) {
		parent := cmd.Parent()
		if parent != nil {
			parent.PreRun(parent, args)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		u := user.GetCurrentUser()
		if u == nil {
			fmt.Fprintln(os.Stderr, "please log in")
			os.Exit(1)
		}
		j, err := u.GetJournal()
		if err != nil {
			fmt.Fprintln(os.Stderr, "unable to get journal for user :", err.Error())
			os.Exit(1)
		}
		j.PrintList()
	},
}

func init() {
	entryCommand.Flags().StringP("add", "a", "", "<JOURNAL-ENTRY-TEXT>")
	entryCommand.AddCommand(listEntryCommand)
	rootCmd.AddCommand(entryCommand)
}
