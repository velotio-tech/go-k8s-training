package entities

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"

	"github.com/farkaskid/go-k8s-training/assignment2/storage"
)

type User struct {
	Username, Password string
	Journal            *Journal
}

func (user *User) Load() error {
	data, err := storage.Load(user.Username+".gob", user.Password)

	if err != nil {
		log.Fatalln("Failed to decode journal", err)

		user.Journal = &Journal{}
		return err
	}

	journal, decryptedData := Journal{}, bytes.Buffer{}
	_, err = decryptedData.Write(data)

	if err != nil {
		log.Fatalln("Failed to decode users", err)

		user.Journal = &Journal{}
		return err
	}

	dec := gob.NewDecoder(&decryptedData)
	err = dec.Decode(&journal)

	user.Journal = &journal

	return err
}

func (user *User) Dump() error {
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)

	err := enc.Encode(user.Journal)

	if err != nil {
		log.Fatalln("Failed to encode user's journal", err)

		return err
	}

	return storage.Dump(data.Bytes(), user.Username+".gob", user.Password)
}

func (user *User) AddEntry(text string, timestamp int64) {
	user.Journal.Entries = append(user.Journal.Entries, JournalEntry{text, timestamp})
}

func (user *User) ReadJournal() {
	for _, entry := range user.Journal.Entries {
		fmt.Println(time.Unix(entry.Timestamp, 0), ": --> ", entry.Data)
	}
}
