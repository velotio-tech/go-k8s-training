package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func ListAllEntries(username string) {

	path := "./local_db/" + username + "_journal.txt"
	data, err := ReadFromFile(path)

	if err != nil {
		log.Fatal(err)
		fmt.Println("Sorry some error occured!!")
		return
	}

	if len(data) == 0 {
			fmt.Println("\nYou dont have any entries")
			return
	}

	fmt.Println()

	for _, entry :=  range data {
		fmt.Println(entry)
	}
	
}

func AddNewEntry(username, dataEntry string) {
	var inputData string

	if dataEntry == "" {
		fmt.Println("\nPlease write your entry here.")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		inputData = input
	} else {
		inputData = dataEntry
	}

	entry := strings.TrimSuffix(inputData, "\n")

	entryFilePath := "./local_db/"+ username +"_journal.txt"

	entries, err := ReadFromFile(entryFilePath) 

	entry = getFormattedDate() + " " + entry

	if err != nil {
		fmt.Println("Some error occured!!")
		return
	}

	if(len(entries) == 5) {
		newEntries := entries[1:]
		newEntries = append(newEntries, entry)
		err := WriteToFile(newEntries, entryFilePath)
		if err != nil {
			log.Fatal(err)
			fmt.Println("Some error occured!!")
			return
		}
	} else {
		err := AppendToFile(entry, entryFilePath)
				
		if err != nil {
			fmt.Println("Some error occured!!")
			return
		}
	}

	fmt.Println("\nData added successfully!!")

}

func getFormattedDate() string {
	currentTime := time.Now()
	return fmt.Sprint(currentTime.Format("06-Jan-02 03:04pm"))
}
