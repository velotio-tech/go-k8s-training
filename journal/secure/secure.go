package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

var (
	SECRET = []byte("abc&1*~#^2^#s0^=)^^7%b34")
	bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
)

func Encode(message string) string {
	plainText := []byte(message)
	block, err := aes.NewCipher(SECRET)
	if err != nil {
		panic(err)
	}
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return ""
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	return base64.RawStdEncoding.EncodeToString(cipherText)
}

func Decode(secure string) string {
	cipherText, err := base64.RawStdEncoding.DecodeString(secure)
	if err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(SECRET)
	if err != nil {
		panic(err)
	}
	if len(cipherText) < aes.BlockSize {
		err = errors.New("ciphertext block size is too short")
		return ""
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText)
}