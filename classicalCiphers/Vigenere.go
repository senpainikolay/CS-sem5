package classicalciphers

import (
	"fmt"
	"math"
	"strings"
)

type Vigenere struct {
	// package private
	keystream     string
	msg           string
	spaceIndexArr []int
}

func (c *Vigenere) Initialize(key string, msg string) {
	c.spaceIndexArr = make([]int, 0)
	for i := 0; i < len(msg); i++ {
		if msg[i] == 32 {
			c.spaceIndexArr = append(c.spaceIndexArr, i)
		}
	}
	c.msg = strings.ReplaceAll(strings.ToUpper(msg), " ", "")
	c.keystream = strings.ToUpper(key)
	remainder := math.Mod(float64(len(c.msg)), float64(len(c.keystream)))

	for {
		if len(c.keystream) >= len(c.msg) {
			break
		}
		c.keystream += strings.ToUpper(key)
	}

	if remainder != 0.0 {
		adjustedSlice := c.keystream[:len(c.keystream)-len(key)+int(remainder)]
		c.keystream = adjustedSlice
	}
	fmt.Printf("The plaintext: %v \n", string(c.msg))
	fmt.Printf("The key transfomed in keystream: %v \n", string(c.keystream))
}

func (c *Vigenere) Encrypt() string {

	var encrypted string

	for i := 0; i < len(c.msg); i++ {
		enCh := math.Mod(float64(int(c.msg[i])+int(c.keystream[i])-2*65), 26)
		encrypted += string(int(enCh) + 65)

	}
	for _, idx := range c.spaceIndexArr {
		newSlice := encrypted[:idx] + " " + encrypted[idx:]
		encrypted = newSlice
	}

	return encrypted
}

func (c *Vigenere) Decrypt() string {

	var decrypted string

	for i := 0; i < len(c.msg); i++ {
		deCh := math.Mod(float64(int(c.msg[i])-int(c.keystream[i])-2*65), 26)
		decrypted += string(int(deCh) + 26 + 65)
	}

	for _, idx := range c.spaceIndexArr {
		newSlice := decrypted[:idx] + " " + decrypted[idx:]
		decrypted = newSlice
	}

	return decrypted
}
