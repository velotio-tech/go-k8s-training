package utils

import (
	"io/ioutil"
	"strings"
)

func LoadJournalData() error {

	for k := range journalData {
		delete(journalData, k)
	}

	encryptedData, err := ioutil.ReadFile(JOURNAL_DATA)
	if err != nil {
		return err
	}

	byteData, err := decryptFile(encryptedData)
	if err != nil {
		return err
	}

	d := string(byteData)

	if d == "" {
		return nil
	}

	entries := strings.Split(d, "~")

	for _, entry := range entries {
		if entry == "" {
			continue
		}
		userData := strings.Split(entry, "#")
		username := userData[0]
		dateTime := userData[1]
		message := userData[2]
		journalData[username] += dateTime + "#" + message + "#"
	}

	return nil
}

func LoadAuthData() error {

	for k := range authData {
		delete(authData, k)
	}
	totalUsers = 0

	encryptedData, err := ioutil.ReadFile(AUTH_DATA)
	if err != nil {
		return err
	}

	byteData, err := decryptFile(encryptedData)
	if err != nil {
		return err
	}

	d := string(byteData)

	if d == "" {
		return nil
	}

	entries := strings.Split(d, "~")

	for _, entry := range entries {
		if entry == "" {
			continue
		}
		userData := strings.Split(entry, "#")
		username := userData[0]
		password := userData[1]
		authData[username] = password
		totalUsers++
	}

	return nil
}
