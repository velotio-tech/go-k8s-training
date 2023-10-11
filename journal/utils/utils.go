package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const USERS_PATH  = "users.txt"

func AddUser(username, password string) {

	userData := username + " " + password

	if err := AppendToFile( userData, USERS_PATH); err != nil {
		log.Fatal(err)
		fmt.Println("\nSome error occured while adding a new users.")
	} 
	
}

func AppendToFile (data, path string) error {
	file, err := os.OpenFile(path, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	encryptedData, encryptErr := Encrypt(data)

	if encryptErr != nil {
		return encryptErr
	}

	if _, e := file.WriteString(encryptedData + "\n"); e != nil {
		return err
	}

	return nil
}

func WriteToFile (data []string, path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer file.Close()

	for _, line := range data {
		encryptedLine, encryptErr := Encrypt(line)
		if encryptErr != nil {
			return encryptErr
		}
		_, err := file.WriteString(encryptedLine + "\n")

		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}



func ReadFromFile(path string) ([]string, error) {

	
	file, err := os.Open(path)
	fileExists := os.IsExist(err)

	if err != nil {
		if !fileExists {
			return []string{} , nil
		}
		return nil, err
	}

	defer file.Close()

	var fileData []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text, _ := Decrypt(scanner.Text())
		fileData = append(fileData, text)
	}

	return fileData, scanner.Err()

}

func UserInput(username string) {
	var input int

	for {
		fmt.Println("\n1. List all entries.\n2. Create New entry.\n3. Exit")
		fmt.Scanln(&input)

		switch input { 
			case 1:
				ListAllEntries(username)
				
			case 2:
				AddNewEntry(username, "")

			case 3:
				os.Exit(0)

			default: 
			fmt.Println("Please enter a valid input.")
		}

	}
}
