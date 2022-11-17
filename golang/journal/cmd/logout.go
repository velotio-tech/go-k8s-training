package cmd

import (
	"github.com/spf13/cobra"
	"github.com/velotio-ajaykumbhar/journal/app/service"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		service.Logout()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

}
