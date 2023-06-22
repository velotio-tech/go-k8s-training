package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var journalData = make(map[string]string)
var authData = make(map[string]string)
var totalUsers = 0

const (
	MAX_USERS    = 10
	AUTH_DATA    = "../authData.txt"
	JOURNAL_DATA = "../journalData.txt"
)

func AddNewUser(username, password string) error {

	err := checkUserExist(username)

	if err != nil {
		return err
	}

	if totalUsers > MAX_USERS {
		return errors.New("maximum user count exceeded! maximum 10 accounts are allowed")
	}

	err = validateEntry(username, "username")
	if err != nil {
		return err
	}
	err = validateEntry(password, "password")
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(AUTH_DATA)
	if err != nil {
		return err
	}

	decryptedData, err := decryptFile(data)
	if err != nil {
		return nil
	}

	authData := string(decryptedData) + username + "#" + password + "~"

	encryptedData, err := encryptFile([]byte(authData))
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(AUTH_DATA, encryptedData, 0644)
	if err != nil {
		return err
	}

	LoadAuthData()

	return nil

}

func StartUserPrompt(username string) error {

	for {

		fmt.Println("1.Create New Entry in the Journal\n2.List all Entries in Journal\n3.Logout")

		userInput := 0

		fmt.Scan(&userInput)

		switch userInput {
		case 1:
			if err := addNewEntry(username); err != nil {
				return err
			}
		case 2:
			listAllEntries(username)
		case 3:
			fmt.Println("Logging out ...")
			return nil
		default:
			fmt.Println("Entered wrong option!")
		}

	}

}

func addNewEntry(username string) error {
	var message string
	fmt.Print("Enter the new journal entry: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message = scanner.Text()
	if err := AddNewEntry(username, message); err != nil {
		return err
	}
	return nil
}

func AddNewEntry(username, message string) error {

	if strings.Contains(message, "#") || strings.Contains(message, "~") {
		return errors.New("message should not contain characters `#` and `~`")
	}
	if message == "" {
		return errors.New("message should not be blank")
	}
	currTime := time.Now().Format(time.RFC1123)

	data, err := ioutil.ReadFile(JOURNAL_DATA)
	if err != nil {
		return err
	}

	decryptedData, err := decryptFile(data)
	if err != nil {
		return nil
	}

	entry := string(decryptedData) + username + "#" + string(currTime) + "#" + message + "~"

	encryptedData, err := encryptFile([]byte(entry))
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(JOURNAL_DATA, encryptedData, 0644)
	if err != nil {
		return err
	}

	LoadJournalData()

	return nil
}

func listAllEntries(username string) {

	if journalData[username] == "" {
		fmt.Println("You don't have any entry added")
	}

	userData := strings.Split(journalData[username], "#")

	type entry struct {
		time    string
		message string
	}

	var journalEntry []entry

	for index, data := range userData {
		if data == "" {
			continue
		}
		if index%2 == 0 {
			journalEntry = append(journalEntry, entry{data, ""})
		} else {
			journalEntry[len(journalEntry)-1].message = data
		}
	}

	for _, each := range journalEntry {
		fmt.Println("\n\t\t--------------------")
		fmt.Println("\tTime:", each.time)
		fmt.Println("\tEntry:", each.message)
	}

}
