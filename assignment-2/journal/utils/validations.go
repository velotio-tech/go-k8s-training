package utils

import (
	"errors"
	"strings"
)

func ValidateUser(username, password string) error {
	if authData[username] == "" {
		return errors.New("user with this username doesn't exists, please check the username or try signup")
	} else if authData[username] != password {
		return errors.New("authentication failed! password didn't matched")
	}
	return nil
}

func checkUserExist(username string) error {
	if authData[username] != "" {
		return errors.New("user with this username already exists please try other username or login if you are already signed up")
	}
	return nil
}

func validateEntry(username, entryType string) error {
	if len(username) < 4 {
		return errors.New(entryType + " should be atleast 4 characters long")
	}
	if len(username) > 10 {
		return errors.New(entryType + " should be atmost 10 characters long")
	}
	if strings.Contains(username, "#") || strings.Contains(username, "~") {
		return errors.New(entryType + " should not contain `#` and `~` characters")
	}
	return nil
}
