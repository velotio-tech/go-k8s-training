package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Journal App",
	Long:  `All software has versions. This is Journal's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Journal Application v0.9 -- HEAD")
	},
}
