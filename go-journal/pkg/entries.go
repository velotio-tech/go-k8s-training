package pkg

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper" // used for importing env variables
)

func GetUserInput(username string) {
	var input string
	for {
		fmt.Println("\n1. List all entries")
		fmt.Println("2. Create a new entry")
		fmt.Println("3. Exit")
		fmt.Scanln(&input)

		switch input {
		case "1":
			ListAllEntries(username)
		case "2":
			emptyStr := ""
			CreateNewEntry(username, emptyStr)
		case "3":
			os.Exit(0)
		default:
			fmt.Println("Please enter a valid input.")
		}
	}
}

func ListAllEntries(username string) {

	// https://stackoverflow.com/questions/17071286/how-can-i-open-files-relative-to-my-gopath
	relativePath := fmt.Sprintf("./local_db/%s_journal.txt", username)
	absPath, _ := filepath.Abs(relativePath)
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Println("You don't have any journals yet!")
	} else {
		fmt.Println("Your Journal is:")
	}
	key := viperEnvVariable("KEY")
	val, _ := strconv.ParseUint(key, 10, 32)
	stringData := CaesarCipherDecrypt(string(data), uint(val))
	fmt.Println(stringData)
}

func CreateNewEntry(username, entry string) {
	
	var userEntry string
	file, _ := os.Open("./local_db/" + username + "_journal.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	if lineCount >= 50 {
		fmt.Println("Entries Limit Exceed!!")
		return
	}
	var line string
	if entry == "" {
		fmt.Println("Please write your entry:")
		reader := bufio.NewReader(os.Stdin)
		line, _ = reader.ReadString('\n')
		// https://stackoverflow.com/questions/44448384/how-remove-n-from-lines
		line = strings.TrimSuffix(line, "\n")
	} else {
		line = entry
	}
	// fmt.Scanln(&userEntry)
	date := getFormattedDate()
	userEntry = "\n" + date + " - " + line
	filePath := "./local_db/" + username + "_journal.txt"

	key := viperEnvVariable("KEY")
	val, _ := strconv.ParseUint(key, 10, 32)
	userEntry = CaesarCipherEncrypt(userEntry, uint(val))
	if err := WriteToFile(userEntry, filePath); err != nil {
		fmt.Println("err: could not write to file", err)
	} else {
		fmt.Println("Your data inserted successfully. ")
		if entry == "" {
			GetUserInput(username)
		}
	}
}

// golangprograms.com/get-current-date-and-time-in-various-format-in-golang.html
func getFormattedDate() string {
	currentTime := time.Now()
	return fmt.Sprint(currentTime.Format("06-Jan-02 03:04pm"))
}

// https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66
func viperEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}
