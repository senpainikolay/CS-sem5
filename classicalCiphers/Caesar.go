package classicalciphers

import (
	"math"
	"unicode"
)

type CaesarClasic struct{}

func (c *CaesarClasic) Encrypt(message string, key int) string {

	var encrypted string

	for _, ch := range message {
		if ch == ' ' {
			encrypted += " "
			continue
		}

		enCh := math.Mod(float64(int(unicode.ToUpper(ch))-65+key), 26)
		encrypted += string(int(enCh) + 65)
	}

	return encrypted
}

func (c *CaesarClasic) Decrypt(message string, key int) string {

	var decrypted string

	for _, ch := range message {
		if ch == ' ' {
			decrypted += " "
			continue
		}

		deCh := math.Mod(float64(int(unicode.ToUpper(ch))-65-key+26), 26)
		decrypted += string(int(deCh) + 65)
	}

	return decrypted
}
