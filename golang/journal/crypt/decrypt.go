package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

func Decrypt(encryptedtext string) (string, error) {
	// fmt.Println(encryptedtext)
	vp := viper.New()

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")

	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	keyString := vp.GetString("key")
	block, err := aes.NewCipher([]byte(keyString))
	if err != nil {
		return "", err
	}

	ciphertext, _ := base64.RawStdEncoding.DecodeString(encryptedtext)

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}
