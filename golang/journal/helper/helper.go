package helper

import (
	"os"
)

func CheckFile(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		_, err := os.Create(filename)
		return err
	}
	return nil
}
