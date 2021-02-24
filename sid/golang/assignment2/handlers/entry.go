package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/farkaskid/go-k8s-training/assignment2/entities"
	"github.com/spf13/cobra"
)

func Entry() func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		journalEntry, _ := cmd.Flags().GetString("add")

		username = strings.TrimSpace(username)
		password = strings.TrimSpace(password)
		journalEntry = strings.TrimSpace(journalEntry)

		if len(username) == 0 || len(password) == 0 {
			fmt.Println("Please supply a username or password")
		}

		app := entities.App{Masterfile: "app.gob", Passphrase: "hello"}
		app.Load()

		user, ok := app.AuthenticateUser(username, password)

		if !ok {
			fmt.Println("Bad username/password")

			return
		}

		if len(journalEntry) != 0 {
			user.AddEntry(journalEntry, time.Now().Unix())
		} else {
			user.ReadJournal()
		}
	}
}
