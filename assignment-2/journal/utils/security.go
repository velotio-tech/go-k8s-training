package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

const KEY = "encryptiondecryptionkeywith32byt"

func encryptFile(plainText []byte) ([]byte, error) {

	if string(plainText) == "" {
		return nil, nil
	}

	key := []byte(KEY)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nonce, nonce, plainText, nil)

	return cipherText, nil
}

func decryptFile(cipherText []byte) ([]byte, error) {

	if string(cipherText) == "" {
		return nil, nil
	}

	key := []byte(KEY)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := cipherText[:gcm.NonceSize()]

	cipherText = cipherText[gcm.NonceSize():]

	plainText, err := gcm.Open(nil, nonce, cipherText, nil)

	if err != nil {
		return nil, err
	}

	return plainText, nil
}
