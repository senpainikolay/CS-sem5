package ciphers

import (
	"fmt"
	"math"
	"unicode"
)

type CaesarPermutation struct {
	SecretWord string
	Alphabet   []rune
}

func (c *CaesarPermutation) InitializeAlphabet() {
	tempAlphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, ch := range c.SecretWord {
		ch = unicode.ToUpper(ch)
		if arrContainsElem(c.Alphabet, ch) {
			continue
		}
		c.Alphabet = append(c.Alphabet, ch)
	}

	for _, ch := range tempAlphabet {
		ch = unicode.ToUpper(ch)
		if arrContainsElem(c.Alphabet, ch) {
			continue
		}
		c.Alphabet = append(c.Alphabet, ch)
	}
	fmt.Printf(" The secret word: %v \n", c.SecretWord)
	fmt.Printf("The permutated alphabet: %v \n", string(c.Alphabet))
}

func (c *CaesarPermutation) Encrypt(message string, key int) string {

	var encrypted string

	for _, ch := range message {
		if ch == ' ' {
			encrypted += " "
			continue
		}
		ch = unicode.ToUpper(ch)
		idx := float64(returnIndexOfElemFromArr(c.Alphabet, ch) + key)

		enCh := c.Alphabet[int(math.Mod(idx, 26))]
		encrypted += string(enCh)
	}

	return encrypted
}

func (c *CaesarPermutation) Decrypt(message string, key int) string {

	var decrypted string

	for _, ch := range message {
		if ch == ' ' {
			decrypted += " "
			continue
		}

		ch = unicode.ToUpper(ch)
		idx := float64(returnIndexOfElemFromArr(c.Alphabet, ch) - key + 26)

		enCh := c.Alphabet[int(math.Mod(idx, 26))]
		decrypted += string(enCh)
	}

	return decrypted
}

func arrContainsElem(arr []rune, ch rune) bool {
	for _, elem := range arr {
		if elem == ch {
			return true
		}

	}
	return false
}

func returnIndexOfElemFromArr(arr []rune, ch rune) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == ch {
			return i
		}
	}
	return -1
}
