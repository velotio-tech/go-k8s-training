package main

import "os"

func save(command string) {

	file, err := os.OpenFile("history.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	if _, err = file.WriteString(command); err != nil {
		if err != nil {
			panic(err.Error())
		}
	}
}
