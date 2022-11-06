package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

func encrypt(keyString string, stringToEncrypt string) (encryptedString string) {
	// convert key to bytes
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func decrypt(keyString string, stringToDecrypt string) string {
	key, _ := hex.DecodeString(keyString)
	ciphertext, _ := base64.URLEncoding.DecodeString(stringToDecrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

type Person struct {
	Name string
	Age  int
}

func main() {
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.

	err := enc.Encode(Person{"Harry Potter", 1000})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	key := []byte("this's secret key.enough 32 bits")
	if _, err := rand.Read(key); err != nil {
		panic(err.Error())
	}
	keyStr := hex.EncodeToString(key) //convert to string for saving

	fmt.Println("Encrypting.....")
	// encrypt value to base64
	cryptoText := encrypt(keyStr, string(network.Bytes()))
	fmt.Println(cryptoText)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}

	fmt.Println("Decrypting.....")
	// encrypt base64 crypto to original value
	text := decrypt(keyStr, cryptoText)
	byteBuffer := bytes.NewBuffer([]byte(text))
	dec := gob.NewDecoder(byteBuffer) // Will read from byteBuffer
	var person Person
	err = dec.Decode(&person)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q: %d\n", person.Name, person.Age)
}
