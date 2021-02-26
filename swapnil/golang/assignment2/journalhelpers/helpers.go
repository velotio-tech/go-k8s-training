package journalhelpers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func getHome() string {
	homeMessage := `Welcome to awesome journals
For logging in enter 'login <username> <password>'
For signing up enter 'signup <username> <password>'
For quitting enter 'quit'
For help enter 'help'`
	return homeMessage
}

func getHomeDir() string {
	// returns current user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return homeDir
}

func loadUserData() userData {
	homeDir := getHomeDir()
	filePath := homeDir + "/.journal_app_store"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		// returning empty history slice
		return make(userData)
	}
	// fmt.Println(string(data))
	// 256 bit key which will be used for encryption and decryption
	secretKey := &[32]byte{75, 201, 190, 111, 175, 219, 150, 22, 160, 84, 17, 59, 237, 118, 254, 24, 57, 255, 75, 135, 144, 125, 71, 175, 40, 24, 27, 230, 22, 7, 249, 226}
	descrypted, err := decrypt(data, secretKey)
	buf := bytes.NewBuffer(descrypted)
	dec := gob.NewDecoder(buf)

	m := make(map[string]userdata)

	if err := dec.Decode(&m); err != nil {
		log.Fatal(err)
	}
	return m

}

func storeBytes(path string, data []byte) {

	// create or overrite the file
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error: creating history file", err)
	}
	// finally close the file
	defer file.Close()
	// join every command with new line char for storing in file
	_, error := file.Write(data)
	if error != nil {
		fmt.Println("Error: writing data to file.")
	}
}

func encrypt(plaintext []byte, key *[32]byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func decrypt(ciphertext []byte, key *[32]byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}

// ListEntries for listing entries with username and password
func ListEntries(username, password string) {
	data := loadUserData()
	result := loginHandler(data, []string{"list", username, password})
	if result == true {
		data.listEntries(username)
	}
}

// AddEntry for adding entry with username and password
func AddEntry(username, password, entryText string) {
	data := loadUserData()
	result := loginHandler(data, []string{"list", username, password})
	if result == true {
		data.addEntry(username, entry{Text: entryText, CreatedAt: time.Now()})
		data.storeUserData()
	}
}
