package account

import (
	"assignment2/cryptography"
	"assignment2/utils"
	"bufio"
	"errors"
	"fmt"
	"strings"
)

type User struct {
	Username string
	Password string
}

func (u *User) CreateUser() {
	writeFileObj := utils.GetFileObj(utils.USER_DATA_FILE, true)
	defer writeFileObj.Close()
	writeFileObj.WriteString(fmt.Sprintf("%s:%s\n", u.Username, cryptography.EncryptPassword(u.Password)))
}

func (u *User) AuthenticateUser() error {
	readFileObj := utils.GetFileObj(utils.USER_DATA_FILE, false)
	defer readFileObj.Close()
	scanner := bufio.NewScanner(readFileObj)
	var dbUser User
	for scanner.Scan() {
		entry := scanner.Text()
		userDetails := strings.SplitN(entry, ":", 2)
		if u.Username == userDetails[0] {
			dbUser.Username = userDetails[0]
			dbUser.Password = userDetails[1]
			break
		}
	}
	if dbUser.Username == "" {
		return errors.New("invalid username")
	}
	if !cryptography.ComparePassword(u.Password, dbUser.Password) {
		return errors.New("invalid password")
	}
	return nil
}

func ValidateUsername(username string) error {
	if checkUserExists(username) {
		return errors.New("username already exists")
	}
	if strings.ContainsAny(username, `:+-=,./\|'" `) {
		return errors.New("username contains special characters")
	}
	return nil
}

func checkUserExists(username string) bool {
	fileObj := utils.GetFileObj(utils.USER_DATA_FILE, false)
	defer fileObj.Close()
	scanner := bufio.NewScanner(fileObj)
	for scanner.Scan() {
		entry := scanner.Text()
		userDetails := strings.SplitN(entry, ":", 2)
		if username == userDetails[0] {
			return true
		}
	}
	return false
}

func GetUserCount() int {
	readFileObj := utils.GetFileObj(utils.USER_DATA_FILE, false)
	defer readFileObj.Close()
	scanner := bufio.NewScanner(readFileObj)
	userCount := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		lineText = strings.ReplaceAll(lineText, " ", "")
		if lineText != "" {
			userCount++
		}
	}
	return userCount
}
