package user

import (
	"io"
	"log"
	"os"

	"github.com/velotio-tech/go-k8s-training/pkg/secure"
)

const SessionFile = "session"

var activeUser string

func (u *User) setActive(temporaryLogin bool) {
	activeUser = u.Username
	if !temporaryLogin {
		createSession()
	}
}

func GetCurrentUser() *User {
	if activeUser != "" {
		return &User{
			Username: activeUser,
		}
	}
	return retrieveSession()
}

func retrieveSession() *User {
	file, err := os.OpenFile(SessionFile, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		log.Println("unable to open session file")
		return nil
	}
	sessionBytes, err := io.ReadAll(file)
	if err != nil {
		log.Println("unable to read session file")
		return nil
	}
	user := User{}
	err = secure.Decrypt(string(sessionBytes), &user)
	if err != nil {
		log.Println("unable to decrypt session file")
		return nil
	}
	return &user
}

func createSession() {
	u := GetCurrentUser()
	file, err := os.OpenFile(SessionFile, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		log.Panic("error while opening session file : ", err)
		return
	}
	encryptedData, err := secure.Encrypt(u)
	if err != nil {
		log.Panic("error while encrypting session creds : ", err)
		return
	}
	_, err = file.WriteString(encryptedData)
	if err != nil {
		log.Panic("error while storing session creds : ", err)
		return
	}
}
