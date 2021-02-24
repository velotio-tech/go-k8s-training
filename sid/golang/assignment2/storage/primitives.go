package storage

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/gob"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func InitUserDatabase() {
	_, err := os.Stat("users.gob")
	if err == nil {
		return
	}

	users := make(map[string]string)

	var data bytes.Buffer
	enc := gob.NewEncoder(&data)

	err = enc.Encode(users)
	if err != nil {
		log.Fatal("Encoder error:", err)
	}

	file, err := os.Create("./users.gob")
	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(data.Bytes())

	if err != nil {
		log.Fatal(err)
	}
}

func Dump(data []byte, filename, passpharse string) error {
	encryptedData, err := encrypt(data, passpharse)

	if err != nil {
		log.Fatalln("Failed to encrypt data", err)

		return err
	}

	err = ioutil.WriteFile(filename, encryptedData, 0777)

	if err != nil {
		log.Fatalln("Failed to write to file", filename)
	}

	return err
}

func encrypt(data []byte, passphrase string) ([]byte, error) {
	c, err := aes.NewCipher(getpassphrase(passphrase))

	if err != nil {
		log.Fatalln("Failed to create cipher", err)

		return nil, err
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		log.Fatalln("Failed to create cipher", err)

		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)

	if err != nil {
		log.Fatalln("Failed to populate nonce", err)

		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

func decrypt(data []byte, passphrase string) ([]byte, error) {
	c, err := aes.NewCipher(getpassphrase(passphrase))

	if err != nil {
		log.Fatalln("Failed to create cipher", err)

		return nil, err
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		log.Fatalln("Failed to create cipher", err)

		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, actualData := data[:nonceSize], data[nonceSize:]
	decryptedData, err := gcm.Open(nil, nonce, actualData, nil)

	if err != nil {
		log.Fatalln("Failed to decrypt data", err)

		return nil, err
	}

	return decryptedData, nil
}

func getpassphrase(str string) []byte {
	bytes := []byte(str)

	if len(bytes) > 32 {
		return bytes[:32]
	}

	newpasspharse, oldLength := make([]byte, 32), len(bytes)

	for i := 0; i < 32; i++ {
		newpasspharse[i] = bytes[i%oldLength]
	}

	return newpasspharse
}

func Load(filename, passphrase string) ([]byte, error) {

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		// log.Fatalln("Failed to load data", err)

		return nil, err
	}

	decryptedData, err := decrypt(data, passphrase)

	if err != nil {
		log.Fatalln("Failed to decrypt data", err)

		return nil, err
	}

	return decryptedData, err
}

func AddUser(username, password string) {

}

func DeleteUser(user string) {

}

func AuthenticateUser(username, password string) {

}
