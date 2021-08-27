package encdec

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
)

func Encrypt(plaintext []byte) []byte {
	key := getKey()
	// text := []byte(plaintext)
	text := plaintext
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error while creating cipher key")
		return nil
	}
	b := base64.StdEncoding.EncodeToString(text)
	cipherText := make([]byte, aes.BlockSize+len(b))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println("Error while reading full data")
		return nil
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(b))

	return cipherText
}

func Decrypt(text []byte) []byte {
	key := getKey()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	if len(text) < aes.BlockSize {
		return nil
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil
	}

	return data
}

func getKey() []byte {
	buf, err := ioutil.ReadFile("./key")
	if err != nil {
		fmt.Println("Error reading key ::", err)
	}

	return buf
}
