package util

import "os"

func InitFile(filename string) {
	file, err := os.OpenFile("journal/"+filename, os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile("journal/" + filename)
	if err != nil {
		return data, err
	}
	return data, nil
}

func WriteFile(filename string, data []byte) error {

	err := os.WriteFile("journal/"+filename, data, 0777)
	if err != nil {
		return err
	}
	return nil
}
