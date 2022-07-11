package service

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func SignUpUser() {
	// Scan username and password
	var username, password string
	fmt.Print("Enter Username:")
	fmt.Scanln(&username)
	fmt.Print("Enter Password:")
	bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
	password = string(bytePassword)
	text := username + " " + password
	// Check if User entries limit exceeded
	if !checkForLimit("users.txt", 10) {
		return
	}
	// Open User entries file to append new entry
	file, err := os.OpenFile("users.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Error while opening Users file, err: ", err)
		return
	}
	defer file.Close()
	// Add encrypted user entry to the file
	encryptedEntry := Encrypt(text) + "\n"
	if _, err = file.WriteString(encryptedEntry); err != nil {
		fmt.Println("Error while adding user entry to file, err: ", err)
		return
	}
	fmt.Printf("User: [%s] successfully signed up to the app!!\n", username)
	GetChoice(username)
}

func checkForLimit(filename string, limit int) bool {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return true
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error while opening users file!")
		return false
	}
	defer file.Close()
	var res int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res++
	}
	if res >= limit {
		fmt.Println("User entries limit exceeded!!")
		return false
	}
	return true
}
