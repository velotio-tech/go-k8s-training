package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/farkaskid/go-k8s-training/assignment2/entities"
	"github.com/spf13/cobra"
)

func Signup() func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		app := entities.App{Masterfile: "app.gob", Passphrase: "hello"}
		app.Load()

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter username: ")
		username, _ := reader.ReadString('\n')

		fmt.Print("Enter password: ")
		password, _ := reader.ReadString('\n')

		user := app.AddUser(username, password)

		fmt.Println("User created successfully")

		for {
			fmt.Print("Enter\n1 - To add a new journal entry\n2 - To view existing journal entry\n3 - To exit\n")

			option, _ := reader.ReadString('\n')
			option = strings.TrimSpace(option)

			if option == "3" {
				fmt.Println("Bye")

				user.Dump()
				app.Dump()

				break
			}

			switch option {
			case "1":
				AddNewJournalEntry(reader, user)
			case "2":
				user.ReadJournal()
			default:
				fmt.Println("Invalid option")
				continue
			}
		}
	}
}

func AddNewJournalEntry(reader *bufio.Reader, user entities.User) {
	fmt.Print("Add text: ")

	text, _ := reader.ReadString('\n')
	user.AddEntry(text, time.Now().Unix())

	fmt.Println("Entry added successfully")
}
