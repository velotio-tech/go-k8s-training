package utils

import (
	"errors"
	"os"
)

const APPLICATION_PATH string = "/tmp/journaler"
const USER_DATA_FILE string = "/tmp/journaler/.userdata"

func CheckApplicationExists() bool {
	_, err := os.Stat(APPLICATION_PATH)
	if err != nil {
		return false
	}
	_, err = os.Stat(USER_DATA_FILE)
	return err == nil
}

func SetupApplication() {
	err := os.Mkdir(APPLICATION_PATH, os.ModePerm)
	if err != nil {
		panic(err)
	}
	_, err = os.Create(USER_DATA_FILE)
	if err != nil {
		panic(err)
	}
}

func GetFileObj(filename string, writeMode bool) *os.File {
	if _, err := os.Stat(filename); err == nil {
		var fileObj *os.File
		var fErr error
		if writeMode {
			fileObj, fErr = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		} else {
			fileObj, fErr = os.Open(filename)
		}
		if fErr != nil {
			panic(fErr)
		}
		return fileObj
	} else if errors.Is(err, os.ErrNotExist) {
		fileObj, fErr := os.Create(filename)
		if fErr != nil {
			panic(fErr)
		}
		return fileObj
	} else {
		panic(err)
	}
}
