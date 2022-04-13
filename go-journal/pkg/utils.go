package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type User struct {
	username string
	password string
}

func AddNewUser(username, password string) {
	path := "user.txt"
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	cnt := 0
	for scanner.Scan() {
		cnt++
	}
	if cnt >= 10 {
		fmt.Println("New Users Limit Exceed.")
		return
	}
	key := viperEnvVariable("KEY")
	val, _ := strconv.ParseUint(key, 10, 32)
	username = CaesarCipherEncrypt(string(username), uint(val))
	password = CaesarCipherEncrypt(string(password), uint(val))
	str := username + " " + password

	str = str + "\n"
	filename := "user.txt"
	if err := WriteToFile(str, filename); err != nil {
		fmt.Println("could not write to file:", err)
	}

}

// Reference: https://stackoverflow.com/questions/7151261/append-to-a-file-in-go
func WriteToFile(data, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()
	if _, err := f.Write([]byte(data)); err != nil {
		log.Fatal(err)
		return err
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func ReadLinesScanner() ([]User, error) {
	path := "user.txt"
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var Users []User

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Fields(scanner.Text())
		uname := text[0]
		key := viperEnvVariable("KEY")
		val, _ := strconv.ParseUint(key, 10, 32)
		uname = CaesarCipherDecrypt(string(uname), uint(val))

		pword := text[1]
		pword = CaesarCipherDecrypt(string(pword), uint(val))
		Users = append(Users, User{uname, pword})
	}
	return Users, scanner.Err()
}
