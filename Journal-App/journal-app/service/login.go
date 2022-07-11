package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func LoginUser(username, password string) {
	if username == "" && password == "" {
		fmt.Print("Enter Username:")
		fmt.Scanln(&username)
		fmt.Print("Enter Password:")
		bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
		password = string(bytePassword)
	}
	// check for file exists or not
	if _, err := os.Stat("users.txt"); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("User with username: [%s] does not exist, Please Sign up first!\n", username)
		return
	}
	// Get all users from file
	var ciphertext []byte
	var err error
	if ciphertext, err = ioutil.ReadFile("users.txt"); err != nil {
		fmt.Println("Error while reading users file, err: ", err)
		return
	}
	if ciphertext == nil || len(ciphertext) < 1 {
		fmt.Printf("User with username: [%s] does not exist, Please Sign up first!\n", username)
		return
	}
	lines := strings.Split(string(ciphertext), "\n")
	// Check if user exist or not
	var found bool
	for _, v := range lines {
		if v == "" {
			break
		}
		line := Decrypt(v)
		user := strings.Split(line, " ")
		if user[0] == username {
			if user[1] == password {
				found = true
				break
			}
		}
	}
	if !found {
		fmt.Printf("User with username: [%s] does not exist, Please Sign up first!\n", username)
		return
	}
	fmt.Printf("User with username: [%s] logged in successfully!\n", username)
	GetChoice(username)
}
