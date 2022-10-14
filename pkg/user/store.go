package user

import (
	"bufio"
	"io/fs"
	"os"
	"strings"

	"github.com/velotio-tech/go-k8s-training/pkg/secure"
)

func (u *User) insert() error {
	secureData, err := secure.Encrypt(u)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(UsersFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, fs.FileMode(0600))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(secureData + "\n")
	return err
}

func getIterator() (*userIterator, error) {
	file, err := os.OpenFile(UsersFileName, os.O_RDONLY|os.O_CREATE, fs.FileMode(0600))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	users := []User{}
	for scanner.Scan() {
		encryptedUser := scanner.Text()
		encryptedUser = strings.TrimSuffix(encryptedUser, "\n")
		decryptedUser := User{}
		err := secure.Decrypt(encryptedUser, &decryptedUser)
		if err != nil {
			return nil, err
		}
		users = append(users, decryptedUser)
	}
	return createUserIterator(users), nil
}

func getTotalUserCount() (int, error) {
	userCount := 0
	file, err := os.OpenFile(UsersFileName, os.O_RDONLY|os.O_CREATE, fs.FileMode(0600))
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		userCount++
	}
	return userCount, nil
}
