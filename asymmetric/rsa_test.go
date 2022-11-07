package asymmetric

import (
	"log"
	"testing"
)

var tempSavingDelimitersOfEncryption [][]int

func TestEncrypt(t *testing.T) {

	log.Println("Encryption tests: ")

	table := []struct {
		msg string
		en  []int64
	}{
		{msg: "FaF ze bEsT", en: []int64{343, 0, 729, 343, 343, 0, 27, 8, 1, 8, 8, 1, 0, 1, 27, 8, 729, 512, 216, 729, 1, 1, 125, 512, 64}},
		{msg: "The quick brown fox jumps over the lazy dog", en: []int64{512, 64, 1, 0, 64, 1, 0, 1, 27, 8, 1, 1, 27, 1, 1, 343, 1, 0, 125, 729, 729, 1, 0, 343, 27, 8, 729, 512, 1, 1, 64, 1, 1, 1, 1, 1, 729, 1, 1, 0, 27, 8, 1, 0, 8, 1, 1, 1, 1, 8, 0, 27, 8, 1, 0, 216, 1, 1, 343, 1, 0, 729, 1, 1, 8, 1, 1, 125, 27, 8, 1, 1, 1, 1, 1, 512, 1, 0, 1, 1, 1, 64, 27, 8, 1, 1, 216, 1, 0, 64, 1, 0, 1, 27, 8, 1, 0, 512, 729, 343, 1, 8, 8, 1, 8, 1, 27, 8, 1, 0, 0, 1, 1, 1, 1, 0, 27}},
		{msg: "Test :-)", en: []int64{512, 64, 1, 0, 1, 1, 1, 125, 1, 1, 216, 27, 8, 125, 512, 64, 125, 64, 1}},
		{msg: "AUF 420", en: []int64{216, 125, 512, 125, 343, 0, 27, 8, 125, 8, 125, 0, 64, 512}},
	}

	for _, data := range table {

		asymmetricCph := RSA{}
		en := CipherEncrypt(&asymmetricCph, data.msg)
		tempSavingDelimitersOfEncryption = append(tempSavingDelimitersOfEncryption, asymmetricCph.LenOfAsciiCodes)

		if !int64SlicesEqual(en, data.en) {
			t.Errorf(" The encryption of message: %v failed \n", data.msg)
		}

	}

	log.Println("Encryption tests passed")
}

func TestDencrypt(t *testing.T) {

	log.Println("Decryption tests: ")

	table := []struct {
		de string
		en []int64
	}{
		{de: "FaF ze bEsT", en: []int64{343, 0, 729, 343, 343, 0, 27, 8, 1, 8, 8, 1, 0, 1, 27, 8, 729, 512, 216, 729, 1, 1, 125, 512, 64}},
		{de: "The quick brown fox jumps over the lazy dog", en: []int64{512, 64, 1, 0, 64, 1, 0, 1, 27, 8, 1, 1, 27, 1, 1, 343, 1, 0, 125, 729, 729, 1, 0, 343, 27, 8, 729, 512, 1, 1, 64, 1, 1, 1, 1, 1, 729, 1, 1, 0, 27, 8, 1, 0, 8, 1, 1, 1, 1, 8, 0, 27, 8, 1, 0, 216, 1, 1, 343, 1, 0, 729, 1, 1, 8, 1, 1, 125, 27, 8, 1, 1, 1, 1, 1, 512, 1, 0, 1, 1, 1, 64, 27, 8, 1, 1, 216, 1, 0, 64, 1, 0, 1, 27, 8, 1, 0, 512, 729, 343, 1, 8, 8, 1, 8, 1, 27, 8, 1, 0, 0, 1, 1, 1, 1, 0, 27}},
		{de: "Test :-)", en: []int64{512, 64, 1, 0, 1, 1, 1, 125, 1, 1, 216, 27, 8, 125, 512, 64, 125, 64, 1}},
		{de: "AUF 420", en: []int64{216, 125, 512, 125, 343, 0, 27, 8, 125, 8, 125, 0, 64, 512}},
	}
	asymmetricCph := RSA{}
	for i, data := range table {
		asymmetricCph.LenOfAsciiCodes = tempSavingDelimitersOfEncryption[i]
		de := CipherDecrypt(&asymmetricCph, data.en)

		if de != data.de {
			t.Errorf(" The decyption of array failed on correct message: %v, Your message: %v\n", data.de, de)
		}

	}
	log.Println("Decryption tests passed ")
}

func int64SlicesEqual(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if int64(v) != b[i] {
			return false
		}
	}
	return true
}
