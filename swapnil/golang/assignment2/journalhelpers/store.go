package journalhelpers

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type entry struct {
	Text      string
	CreatedAt time.Time
}

type userdata struct {
	Password string
	Entries  []entry
}

func (u userdata) addEntry(e entry) userdata {
	u.Entries = append(u.Entries, e)
	cnt := len(u.Entries)
	if cnt > 50 {
		u.Entries = u.Entries[cnt-50 : cnt]
	}

	return u
}

type userData map[string]userdata

func (d userData) storeUserData() {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(d); err != nil {
		log.Fatal(err)
		return
	}
	// 256 bit key which will be used for encryption and decryption
	secretKey := &[32]byte{75, 201, 190, 111, 175, 219, 150, 22, 160, 84, 17, 59, 237, 118, 254, 24, 57, 255, 75, 135, 144, 125, 71, 175, 40, 24, 27, 230, 22, 7, 249, 226}
	encrypted, err := encrypt(buf.Bytes(), secretKey)
	if err != nil {
		fmt.Println("Error: couldn't encrypt the data", err)
	}
	homeDir := getHomeDir()
	filePath := homeDir + "/.journal_app_store"
	storeBytes(filePath, encrypted)
}

func (d userData) userExists(u string) bool {
	_, ok := d[u]
	return ok
}

func (d userData) listEntries(u string) {
	if len(d[u].Entries) == 0 {
		fmt.Println("It's empty here. Please add entries first.")
	}
	for _, val := range d[u].Entries {
		// fmt.Println(each.CreatedAt, each.Text)
		fmt.Println(val.CreatedAt.Format("02-01-2006 15:04:05"), val.Text)
	}
}

func (d userData) addUser(u string, p string) bool {
	entries, ok := d[u]
	fmt.Println(entries, ok)
	if ok == false {
		fmt.Println(len(d))
		if len(d) > 10 {
			fmt.Println("Error: only 10 users are supported")
			return false
		}
		d[u] = userdata{p, make([]entry, 0)}
		fmt.Println("Info: user added")
		return true
	} else {
		fmt.Printf("Error: user already exist %v \n", u)
		return false
	}

}

func (d userData) addEntry(u string, e entry) {
	_, ok := d[u]
	if ok == true {
		d[u] = d[u].addEntry(e)
		fmt.Println("Info: Entry added successfully")
	} else {
		fmt.Printf("Error: user %v doesn't exists\n", u)
	}

}
