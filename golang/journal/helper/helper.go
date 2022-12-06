package helper

import (
	"fmt"
	"journal/crypt"
	"os"
	"strings"
)

func CheckFile(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		_, err := os.Create(filename)
		return err
	}
	return nil
}

func Check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		// panic(err)
	}
}

func ByteToSliceOfStrings(ciper []byte) ([]string, error) {
	var sliceData []string
	if len(ciper) == 0 {
		return sliceData, nil
	}
	data, err := crypt.Decrypt(string(ciper))
	Check(err)

	sliceData = strings.Split(data, ",")
	return sliceData, nil
}

func CheckIfDuplicateUser(data []string, name string) bool {
	for _, user := range data {
		if user == name {
			return true
		}
	}
	return false
}

func SliceOfStringsToByte(data []string) []byte {
	str := strings.Join(data, ",")
	return []byte(str)
}

func EncryptAndAppendToFile(byteData []byte, file string) {
	str, err := crypt.Encrypt(byteData)
	Check(err)

	e := os.WriteFile(file, []byte(str), 0666)
	Check(e)
}
