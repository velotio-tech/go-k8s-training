package database

import(
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func getSecret() ([]byte) {
	return make([]byte, 16)
}

func encrypt(plainData string, secret []byte) (string, error) {
	cipherBlock, err := aes.NewCipher(secret)
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aead.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}

	encryptedData := base64.URLEncoding.EncodeToString(aead.Seal(nonce, nonce, []byte(plainData), nil))
	return encryptedData, nil
}

func decrypt(encodedData string, secret []byte) (string, error) {
	encryptData, err := base64.URLEncoding.DecodeString(encodedData)
	if err != nil {
		return "", nil
	}

	cipherBlock, err := aes.NewCipher(secret)
	if err != nil {
		return "", nil
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", nil
	}

	nonceSize := aead.NonceSize()
	if len(encryptData) < nonceSize {
		return "", err
	}
 
	nonce, cipherText := encryptData[:nonceSize], encryptData[nonceSize:]
	plainData, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}
 
	return string(plainData), nil
}

