package stream

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestEncrypt(t *testing.T) {

	log.Println("Encryption tests: ")

	table := []struct {
		pt  string
		key string
		en  string
	}{
		{pt: "The quick brown fox jumps over the lazy dog.", key: "AUF", en: "fb8b56d6a56eff0539f1ea63eb98c615ecfd75389fcf92a31cf15ba5c0c0edfc6dc0344e0695bf834d9a0fef"},
		{pt: "This plain", key: "AUF", en: "fb8b5a85f46bfa073bbf"},
		{pt: "FaF KeK", key: "AUF", en: "e98275d69f7edd"},
		{pt: "Last one lol", key: "FAF", en: "81e64481683d3e8431ad1dba"},
	}

	NR_OF_BYTES := 256
	c := RC4_Init(NR_OF_BYTES)
	for _, data := range table {

		en := c.Encrypt(data.key, data.pt)
		if hex.EncodeToString([]byte(en)) != data.en {
			t.Errorf(" The text %v : The ecryption %v , \n THe correct encryption: %v \n", data.pt, hex.EncodeToString([]byte(en)), data.en)
		}

	}

	log.Println("Encryption tests passed")
}

func TestDencrypt(t *testing.T) {

	log.Println("Decryption tests: ")

	table := []struct {
		pt  string
		key string
		de  string
	}{
		{pt: "fb8b56d6a56eff0539f1ea63eb98c615ecfd75389fcf92a31cf15ba5c0c0edfc6dc0344e0695bf834d9a0fef", key: "AUF", de: "The quick brown fox jumps over the lazy dog."},
		{pt: "fb8b5a85f46bfa073bbf", key: "AUF", de: "This plain"},
		{pt: "e98275d69f7edd", key: "AUF", de: "FaF KeK"},
		{pt: "81e64481683d3e8431ad1dba", key: "FAF", de: "Last one lol"},
	}

	NR_OF_BYTES := 256
	c := RC4_Init(NR_OF_BYTES)
	for _, data := range table {
		hexToByte, err := hex.DecodeString(data.pt)
		if err != nil {
			panic(err)
		}
		de := c.Decrypt(data.key, string(hexToByte))
		if de != data.de {
			t.Errorf(" The text %v : The decryption: %v , \n THe correct decryption: %v \n", data.pt, de, data.de)
		}

	}
	log.Println("Decryption tests passed ")
}
