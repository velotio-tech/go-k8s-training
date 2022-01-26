/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"journal/users"

	"github.com/spf13/cobra"
)

// entryCmd represents the entry command
var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "Add and Remove entry from the journal.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		usr, _ := cmd.Flags().GetString("username")
		pswd, _ := cmd.Flags().GetString("password")
		add, _ := cmd.Flags().GetString("add")
		show, _ := cmd.Flags().GetBool("show")
		if add != "" {
			fmt.Println("Validating user & adding journal entry for the user: ", usr)
			users.AddEntry(usr, pswd, add)
		}
		if show {
			fmt.Println("Showing the journal entry for the user: ", usr)
			users.ShowEntry(usr, pswd)
		}
	},
}

func init() {
	rootCmd.AddCommand(entryCmd)
	entryCmd.Flags().StringP("username", "u", "", "enter username for the user")
	entryCmd.Flags().StringP("password", "p", "", "enter password for the user")
	entryCmd.Flags().StringP("add", "a", "", "enter a journal entry for the user")
	entryCmd.Flags().BoolP("show", "s", false, "shows the journal entry for the user")
	entryCmd.MarkFlagRequired("username")
	entryCmd.MarkFlagRequired("password")

	//entryCmd.MarkFlagRequired("add")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// entryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// entryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
