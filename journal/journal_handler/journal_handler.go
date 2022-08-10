package journal_handler

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"journal/secure"
)

const (
	JOURNALS_LIMIT = 50
	JOURNALS_COLLECTION = "database"
)

func TaskSelection(username string) {
	for {
		fmt.Print("1.Add new journal\n2.List all journals\n3.Exit\nEnter choice: ")
		var choice string
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			AddJournal("",username)
		case "2":
			ListJournals(username)
		case "3":
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func ListJournals(username string) {
	var err error
	homeDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	journalDir := homeDir + "/" + JOURNALS_COLLECTION + "/"
	fileName := username + ".txt"
	userDir := journalDir + fileName
	if _, err := os.Stat(userDir); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("%s don't have any journals :(\n\n\n", username)
		return
	}
	var ciphertext []byte
	ciphertext, err = ioutil.ReadFile(userDir)
	if err != nil {
		fmt.Println("Error while reading journals, err: ", err)
	}
	lines := strings.Split(string(ciphertext), "\n")
	fmt.Println("\nJournals:")
	for _, v := range lines {
		if v == "" {
			break
		}
		line := secure.Decode(v)
		fmt.Println(line)
	}
	fmt.Println()
}

func AddJournal(add, username string) {
	var err error
	homeDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	journalDir := homeDir + "/" + JOURNALS_COLLECTION + "/"
	fileName := username + ".txt"
	userDir := journalDir + fileName
	if !VerifyLimit(userDir, JOURNALS_LIMIT) {
		return
	}
	if err != nil {
		panic(err)
	}
	journal := ""
	if(len(add) > 0) {
		journal = add
	} else {
		fmt.Println("Enter text for the new journal: ")
		reader := bufio.NewReader(os.Stdin)
		journal, _ = reader.ReadString('\n')
		journal = strings.TrimSuffix(journal, "\n")
	}
	timeStamp := time.Now().Format("02-Jul-06 03:04pm")
	newJournal := secure.Encode(timeStamp + "\t" + journal)
	file, err := os.OpenFile(userDir, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Error while opening journals :(, err: ", err)
		return
	}
	defer file.Close()
	if _, err = file.WriteString(newJournal + "\n"); err != nil {
		fmt.Println("Error while adding journal :(, err: ", err)
		return
	}
	fmt.Println("Journal Added :)\n")
}

func VerifyLimit(filename string, limit int) bool {
	_, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return true
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error while opening file :(")
		return false
	}
	defer file.Close()
	var count int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
	}
	if count >= limit {
		fmt.Println("Limit exceeded :(")
		return false
	}
	return true
}