package journal

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/thisisprasad/go-k8s-training/Prasad/golang/assignment2/constants"
	"github.com/thisisprasad/go-k8s-training/Prasad/golang/assignment2/encdec"
)

type JournalEntry struct {
	Timestamp string
	Text      string
}

var userJournData []JournalEntry
var loggedInUser string

//	fetches the user journal data.
func GetUserJournData() []string {
	journData := make([]string, 0, constants.MaxJournalEntry)
	for _, entry := range userJournData {
		journData = append(journData, entry.Timestamp+" - "+entry.Text)
	}

	return journData
}

//	adds an entry into the journal
func AddEntry(je JournalEntry) {
	journLen := len(userJournData)
	if journLen == constants.MaxJournalEntry {
		userJournData = userJournData[1:journLen]
	}
	userJournData = append(userJournData, je)
}

//	Writes user journal data to file.
func WriteJournDataToFile() {
	if len(userJournData) == 0 {
		return
	}
	file, err := os.OpenFile(constants.UserRegistryPath+loggedInUser+constants.UserJournSuffix, os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("Error while opening file ::", err)
		return
	}
	defer file.Close()

	var sb strings.Builder
	for i := 0; i < len(userJournData); i++ {
		sb.WriteString(userJournData[i].Timestamp + "$#v#$" + userJournData[i].Text)
		if i != len(userJournData)-1 {
			sb.WriteString("$@e@$")
		}
	}
	encBuffer := encdec.Encrypt([]byte(sb.String()))
	writer := bufio.NewWriter(file)
	if _, err := writer.Write(encBuffer[:]); err != nil {
		fmt.Println("Error writing journal data to file.")
		return
	}
	writer.Flush()
}

//	Load use journal data in memory.
func LoadUserJournalData(uname string) {
	loggedInUser = uname
	journfp := constants.UserRegistryPath + uname + constants.UserJournSuffix
	_, err := os.Stat(journfp)
	if err != nil && os.IsNotExist(err) {
		//	create user journal file
		os.Create(journfp)
		userJournData = make([]JournalEntry, 0, constants.MaxJournalEntry)
		return
	}

	file, err := os.OpenFile(journfp, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println("Error loading user data ::", err)
		return
	}
	defer file.Close()

	rawData, err := ioutil.ReadFile(journfp)
	if len(rawData) == 0 {
		return
	}
	if err != nil {
		fmt.Println("Error reading journal data from file.")
		return
	}
	decBuffer := encdec.Decrypt(rawData)
	// fmt.Println("rawData:", string(decBuffer))
	strData := string(decBuffer)
	journEntry := strings.Split(strData, constants.JournEntryDelim)
	for _, record := range journEntry {
		entry := strings.Split(record, constants.JournValDelim)
		je := JournalEntry{
			Timestamp: entry[0],
			Text:      entry[1],
		}
		userJournData = append(userJournData, je)
	}
}
