package user

import (
	"fmt"
	"os"
	"strings"

	"github.com/velotio-tech/go-k8s-training/pkg/journal"
)

const MinPasswordLength = 6
const UsersFileName = "users"
const UsersMaxLimit = 10

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func InitializeUser(username, password string) *User {
	return &User{Username: username, Password: password}
}

func (e *User) Validate() error {
	if strings.TrimSpace(e.Username) == "" || len(e.Password) < MinPasswordLength {
		return fmt.Errorf("user validation failed")
	}
	return nil
}

func (u *User) Create() error {
	userCount, err := getTotalUserCount()
	if err != nil {
		return err
	}
	if userCount > UsersMaxLimit {
		return fmt.Errorf("user limit reached")
	}
	if err = u.Validate(); err != nil {
		return err
	}
	if u.Exists() {
		return fmt.Errorf("user with name already exists")
	}
	if err = u.insert(); err != nil {
		return fmt.Errorf("unable to create user")
	}
	u.setActive(false)
	return nil
}

func (u *User) Exists() bool {
	iterator, err := getIterator()
	if err != nil {
		return false
	}
	for iterator.HasNext() {
		userIface := iterator.Get()
		user, ok := userIface.(User)
		if !ok {
			return false
		}
		if user.Username == u.Username {
			return true
		}
	}
	return false
}

func (u *User) Login(temproryLogin bool) error {
	iterator, err := getIterator()
	if err != nil {
		return err
	}
	for iterator.HasNext() {
		userIface := iterator.Get()
		user, ok := userIface.(*User)
		if !ok {
			return fmt.Errorf("invalid user cast")
		}
		if u.Username == user.Username && u.Password == user.Password {
			u.setActive(temproryLogin)
			return nil
		}
	}
	return fmt.Errorf("invalid creds")
}

func (u *User) GetJournal() (*journal.Journal, error) {
	file, err := os.OpenFile(u.Username+".journal", os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}
	return journal.Get(file)
}
