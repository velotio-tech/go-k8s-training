package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Manager struct {
	registeredUsers map[string]*User
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

func (m *Manager) loadRegisteredUsers() error {
	fmt.Println("Loading all the registered Users")
	err := checkFile(REGUSERS)
	if err != nil {
		return err
	}
	data, err := os.ReadFile(REGUSERS)
	var users []User
	if err != nil {
		return err
	}
	json.Unmarshal(data, &users)
	m.registeredUsers = make(map[string]*User)
	for _, user := range users {
		m.registeredUsers[user.Email] = &user
	}
	fmt.Println("Loading registered Users successful")
	return nil
}

func (m *Manager) register(user *User) error {
	if m.registeredUsers == nil {
		err := m.loadRegisteredUsers()
		if err != nil {
			return err
		}
	}
	if _, present := m.registeredUsers[user.Email]; present {
		fmt.Println("user with given email already exists please use a different email ID")
		return errors.New("user with given email already exists please use a different email ID")
	}
	if len(m.registeredUsers) > 10 {
		fmt.Println("can't register more than 10 users")
		return errors.New("can't register more than 10 users")
	}
	m.registeredUsers[user.Email] = user
	user.LoginStatus = true
	return nil
}

func (m *Manager) login(email, password string) error {
	if m.registeredUsers == nil {
		m.loadRegisteredUsers()
	}
	if _, present := m.registeredUsers[email]; !present {
		fmt.Printf("%s not registered, please sign up first\n", email)
		return fmt.Errorf("%s not registered, please sign up first", email)
	}
	if password != m.registeredUsers[email].Password {
		fmt.Println("enter correct password ")
		return errors.New("enter correct password ")
	}
	m.registeredUsers[email].LoginStatus = true
	return nil
}

func (m *Manager) commit() error {
	var users []User

	for _, user := range m.registeredUsers {
		users = append(users, *user)
	}
	bytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	err = os.WriteFile(REGUSERS, bytes, 0666)
	if err != nil {
		return err
	}
	fmt.Println("Data saved to file successfully")
	return nil
}
