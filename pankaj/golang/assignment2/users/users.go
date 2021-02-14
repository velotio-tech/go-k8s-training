package users

import (
	"bufio"
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pankaj9310/go-k8s-training/pankaj/golang/assignment2/journals"
)

var userPath string = "data/users.csv"

var journalPath string = "data/journal/"
var userJournalPath string

func LoginOrRegisterUser() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("User Managment select below option.\n1. Sign Up \n2. Login\n")
	scanner.Scan()
	option, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid option")
	}

	jouranls := &journals.Journal{}
	jouranls.Init(50)

	switch option {
	case 1:
		password, status := registration(scanner)
		if status {
			journalData := journals.ReadFile(userJournalPath, password)
			GetJournalData(jouranls, journalData)
			journals.JournalInput(scanner, userJournalPath, jouranls, password)
		}
	case 2:
		password, status := login(scanner)
		if status {
			journalData := journals.ReadFile(userJournalPath, password)
			GetJournalData(jouranls, journalData)
			journals.JournalInput(scanner, userJournalPath, jouranls, password)
		}
	default:
		fmt.Println("Invalid option. Please try again!")
	}
}

func GetJournalData(journal *journals.Journal, journalData string) {
	for _, ele := range strings.Split(journalData, "\n\n") {
		journal.Capture(ele)
	}
}

func registration(scanner *bufio.Scanner) (string, bool) {
	username, password := takeUserInput(scanner)
	return password, CreateUser(username, password)
}

func login(scanner *bufio.Scanner) (string, bool) {
	username, password := takeUserInput(scanner)
	status := authenticateUser(username, password)
	if !status {
		fmt.Println("Invalid username or password. Please try again!")
	} else {
		fmt.Println("Authernication successfull!")
		return password, true
	}
	return password, false

}

func takeUserInput(scanner *bufio.Scanner) (string, string) {
	var name, password string
	fmt.Printf("Enter Your Name: ")
	for scanner.Scan() {
		name = strings.TrimSpace(scanner.Text())
		if strings.Compare(name, "") != 0 {
			break
		} else {
			fmt.Println("Name can't be empty. Please enter valid name")
		}
	}
	fmt.Printf("Enter Password: ")
	for scanner.Scan() {
		password = strings.TrimSpace(scanner.Text())
		if strings.Compare(password, "") != 0 {
			break
		} else {
			fmt.Println("Password can't be empty. Please enter valid name")
		}
	}
	userJornalFile := GetPassword(name + password)
	userJournalPath = filepath.Join(journalPath, userJornalFile+".txt")
	return name, GetPassword(password)
}

func authenticateUser(username string, password string) bool {
	file, err := os.Open(userPath)
	if err != nil {
		fmt.Println("Error to access user data", err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	users, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error to parse csv file", err)
	}
	for _, user := range users {
		if strings.Compare(username, user[0]) == 0 && strings.Compare(password, user[1]) == 0 {
			return true
		}
	}
	return false
}

func CreateUser(user string, password string) bool {
	if !authenticateUser(user, password) {
		file, err := os.OpenFile(userPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("Error to access user data", err)
		}

		defer file.Close()

		var record [][]string
		record = append(record, []string{user, password})
		csvWriter := csv.NewWriter(file)
		err = csvWriter.WriteAll(record)

		if err != nil {
			fmt.Println("Error to update user data", err)
		}
		fmt.Println("User created successfully!")
	}
	return true

}

func GetPassword(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}
