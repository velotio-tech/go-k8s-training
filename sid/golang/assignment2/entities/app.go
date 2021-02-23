package entities

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/farkaskid/go-k8s-training/assignment2/storage"
)

type App struct {
	Users      map[string]string
	Masterfile string
	Passphrase string
}

func (app *App) Dump() error {
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)

	err := enc.Encode(app.Users)

	if err != nil {
		log.Fatalln("Failed to encode users", err)

		return err
	}

	return storage.Dump(data.Bytes(), app.Masterfile, app.Passphrase)
}

func (app *App) Load() error {
	data, err := storage.Load(app.Masterfile, app.Passphrase)

	if err != nil {
		// log.Fatalln("Failed to decode users", err)

		app.Users = make(map[string]string)
		return err
	}

	users, decodedData := make(map[string]string), bytes.Buffer{}
	_, err = decodedData.Write(data)

	if err != nil {
		log.Fatalln("Failed to decode users", err)

		app.Users = make(map[string]string)
		return err
	}

	dec := gob.NewDecoder(&decodedData)
	err = dec.Decode(&users)

	app.Users = users

	return err
}

func (app *App) AddUser(username, password string) User {
	app.Users[username] = password

	return User{Username: username, Password: password, Journal: &Journal{}}
}

func (app *App) AuthenticateUser(username, password string) (User, bool) {
	pass, ok := app.Users[username]

	if ok && pass == password {
		user := User{Username: username, Password: password}
		user.Load()

		return user, true
	}

	return User{}, false
}
