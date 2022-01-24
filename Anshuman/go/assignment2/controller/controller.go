package controller

import (
	"assignment2/account"
	"assignment2/journal"
	"assignment2/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunApplication(u *account.User) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter 1 for reading Journal entries.")
		fmt.Println("Enter 2 for writing a new Journal entry.")
		fmt.Println("Enter 3 for logging out.")
		inp, _ := reader.ReadString('\n')
		inp = strings.ReplaceAll(inp, " ", "")
		inp = strings.ReplaceAll(inp, "\n", "")

		switch inp {
		case "1":
			journal.ReadJournalEntries(u)
		case "2":
			journal.CreateJournalEntry(u)
		case "3":
			os.Exit(0)
		default:
			fmt.Println("Invalid input entered, please type in your choice using the options provided.")
		}
	}
}

func CreateNewAccount() {
	if account.GetUserCount() > 10 {
		fmt.Println("Sorry but no new accounts can be created right now, please try again later.")
		os.Exit(0)
	}
	reader := bufio.NewReader(os.Stdin)
	var username, password string
	for {
		fmt.Println("Please enter your username. Do note that the usernames need to be unique.")
		fmt.Println("Do note that usernames cannot contain the following characters: ", `:+-=,./\|'" `)
		username, _ = reader.ReadString('\n')
		username = strings.ReplaceAll(username, " ", "")
		username = strings.ReplaceAll(username, "\n", "")
		username = strings.ToLower(username)
		err := account.ValidateUsername(username)
		if err != nil {
			fmt.Println("Invalid username entered, please try again.")
			continue
		} else {
			break
		}
	}
	for {
		fmt.Println("Please enter the password. Password needs to be atleast 6 characters.")
		password, _ = reader.ReadString('\n')
		password = strings.ReplaceAll(password, " ", "")
		password = strings.ReplaceAll(password, "\n", "")
		if len(password) < 6 {
			fmt.Print("Password too short, try again.\n\n")
		} else {
			break
		}
	}
	newUser := account.User{Username: username, Password: password}
	newUser.CreateUser()
	userJournalFile := journal.GetUserJournalFile(&newUser)
	utils.GetFileObj(userJournalFile, true)
	RunApplication(&newUser)
}

func Login() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.ReplaceAll(username, "\n", "")
	username = strings.ReplaceAll(username, " ", "")
	username = strings.ToLower(username)
	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.ReplaceAll(password, " ", "")
	password = strings.ReplaceAll(password, "\n", "")
	u := account.User{Username: username, Password: password}
	if u.AuthenticateUser() != nil {
		fmt.Println("Invalid credentials entered. Please try again.")
	} else {
		RunApplication(&u)
	}

}

func SaveSingleEntry(username, password, entry *string) {
	u := account.User{Username: *username, Password: *password}
	if u.AuthenticateUser() == nil {
		journal.CreateJournalEntry(&u, *entry)
	} else {
		fmt.Println("Invalid options provided, please try again.")
	}
}
