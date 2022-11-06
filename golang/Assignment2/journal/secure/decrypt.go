package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"

	"github.com/spf13/viper"
)

// decrypt from base64 to decrypted string
func Decrypt(encrypttext string) (string, error) {
	keyString := viper.GetString("key")

	block, err := aes.NewCipher([]byte(keyString))
	if err != nil {
		return "", err
	}
	ciphertext, _ := base64.RawStdEncoding.DecodeString(encrypttext)
	// fmt.Println(len(ciphertext), aes.BlockSize)
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}
