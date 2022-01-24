package journal

import (
	"assignment2/account"
	"assignment2/cryptography"
	"assignment2/utils"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func CreateJournalEntry(u *account.User, jEntry ...string) {
	reader := bufio.NewReader(os.Stdin)

	// Get timestamp of current time
	nowTimestamp := time.Now()
	var line string
	if len(jEntry) != 0 {
		line = jEntry[0]
	} else {
		fmt.Println("Enter your journal entry: ")
		line, _ = reader.ReadString('\n')
		line = strings.ReplaceAll(line, "\n", "")
	}
	entry := nowTimestamp.String() + ":" + line

	// Get file object
	userFile := GetUserJournalFile(u)
	fileObj := utils.GetFileObj(userFile, false)

	fileContents, _ := ioutil.ReadAll(fileObj)
	fileSlice := make([]string, 1)
	if string(fileContents) != "" {
		fileString := cryptography.Decrypt(string(fileContents), u.Password)
		fileSlice = strings.Split(fileString, "\n")
	}
	fileSlice = append(fileSlice, entry)
	fileObj.Close()

	// Truncate existing journal file and write new data to it
	writeObj, _ := os.OpenFile(userFile, os.O_WRONLY|os.O_TRUNC, 0644)
	defer writeObj.Close()
	writeObj.Truncate(0)
	writeObj.Seek(0, 0)

	if len(fileSlice) < 50 {
		newFileContents := strings.Join(fileSlice, "\n")
		encryptedFileContents := cryptography.Encrypt(newFileContents, u.Password)
		writeObj.WriteString(encryptedFileContents)
	} else {
		fileSlice := fileSlice[len(fileSlice)-50:]
		newFileContents := strings.Join(fileSlice, "\n")
		encryptedFileContents := cryptography.Encrypt(newFileContents, u.Password)
		writeObj.WriteString(encryptedFileContents)
	}
}

func ReadJournalEntries(u *account.User) {
	fileName := GetUserJournalFile(u)
	fileObj := utils.GetFileObj(fileName, false)
	encryptedContents, _ := ioutil.ReadAll(fileObj)
	if string(encryptedContents) == "" {
		fmt.Println("Journal empty, please provide an entry before trying to read!")
		return
	}
	fileContents := cryptography.Decrypt(string(encryptedContents), u.Password)
	fmt.Print(fileContents, "\n\n")
}

func GetUserJournalFile(user *account.User) string {
	userJournalPath := utils.APPLICATION_PATH + "/" + user.Username
	return userJournalPath
}
