package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/jshiwam/journal/secure"
)

type Manager struct {
	registeredUsers []*User
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

// TODO @shiwam: Figure out how to cache the data and not load the file everytime you run signup command
func (m *Manager) LoadRegisteredUsers() error {
	err := checkFile(REGUSERS)
	if err != nil {
		return err
	}
	cipher, err := os.ReadFile(REGUSERS)
	if err != nil {
		return err
	}
	var data []byte

	if len(cipher) > 0 {
		decrypt, err := secure.Decrypt(string(cipher))
		if err != nil {
			return err
		}
		data = []byte(decrypt)
	} else {
		data = make([]byte, 0)
	}

	json.Unmarshal(data, &m.registeredUsers)

	return nil
}

func (m *Manager) Register(user *User) error {
	err := m.LoadRegisteredUsers()
	if err != nil {
		return err
	}

	for _, usr := range m.registeredUsers {
		if usr.Email == user.Email {
			return errors.New("user with given email already exists please use a different email ID")
		}
	}
	if len(m.registeredUsers) > 10 {
		return errors.New("can't register more than 10 users")
	}
	m.registeredUsers = append(m.registeredUsers, user)
	user.LoginStatus = true
	m.Commit()
	return nil
}

func (m *Manager) GetUser(email string) *User {
	for _, usr := range m.registeredUsers {
		if usr.Email == email {
			return usr
		}
	}
	return nil
}

func (m *Manager) Login(email, password string) error {
	err := m.LoadRegisteredUsers()
	if err != nil {
		return err
	}
	exist := m.GetUser(email)
	if exist == nil {
		return fmt.Errorf("%s not registered, please enter the registered email", email)
	}

	if password != exist.Password {
		return errors.New("enter correct password ")
	}
	exist.LoginStatus = true
	m.Commit()
	return nil
}

func (m *Manager) Logout(email string) error {
	err := m.LoadRegisteredUsers()
	if err != nil {
		return err
	}
	var exist *User
	for _, usr := range m.registeredUsers {
		if usr.Email == email {
			exist = usr
			break
		}
	}
	if exist == nil {
		return fmt.Errorf("%s not registered, please sign up first", email)
	}
	exist.LoginStatus = false
	m.Commit()
	return nil
}

func (m *Manager) Commit() error {
	bytes, err := json.Marshal(m.registeredUsers)
	if err != nil {
		return err
	}
	encrypt, err := secure.Encrypt(bytes)
	if err != nil {
		return err
	}
	err = os.WriteFile(REGUSERS, []byte(encrypt), 0666)
	if err != nil {
		return err
	}
	return nil
}
