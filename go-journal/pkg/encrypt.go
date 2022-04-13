package pkg

import "strings"

// Reference: https://golangprojectstructure.com/caesar-cipher-secret-messages/
func CaesarCipherEncrypt(value string, shift uint) string {
	var builder strings.Builder
	for _, r := range value {
		if r >= 'A' && r <= 'Z' {
			r = caesarCipherShiftRune(r, shift)
		} else if r >= 'a' && r <= 'z' {
			r = caesarCipherShiftRune2(r, shift)
		}
		builder.WriteRune(r)
	}
	return builder.String()
}

func caesarCipherShiftRune(r rune, shift uint) rune {
	s := rune(shift % 26)
	if s == 0 {
		return r
	}
	return (((r - 'A') + s) % 26) + 'A'
}
func caesarCipherShiftRune2(r rune, shift uint) rune {
	s := rune(shift % 26)

	if s == 0 {
		return r
	}

	return (((r - 'a') + s) % 26) + 'a'
}
func CaesarCipherDecrypt(value string, shift uint) string {
	return CaesarCipherEncrypt(value, 26-(shift%26))
}
