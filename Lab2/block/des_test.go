package block

import (
	"log"
	"testing"

	p "github.com/senpainikolay/CS-sem5/Lab2/block/permutations"
)

func TestEncrypt(t *testing.T) {

	log.Println("Encryption tests: ")

	table := []struct {
		pt  string
		key string
		en  string
	}{
		{pt: "utmusmas", key: "caesarok", en: "4ba9fc3cb6e3e949"},
		{pt: "chisinau", key: "hmoklolw", en: "a2e2d061425dcdf2"},
		{pt: "ATBACKSU", key: "SECONDKE", en: "8abebfadc966b3c2"},
		{pt: "LMAOCIPH", key: "whygodwh", en: "3b566979fe9440f2"},
	}
	des := GetDES()
	for _, data := range table {

		des.KeyInit(data.key)
		hexSt, _ := des.Encrypt(p.StringToHex(data.pt))

		if hexSt != data.en {
			t.Errorf(" The text %v : The ecryption %v , \n THe correct encryption: %v \n", data.pt, hexSt, data.en)
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
		{pt: "4ba9fc3cb6e3e949", key: "caesarok", de: "utmusmas"},
		{pt: "a2e2d061425dcdf2", key: "hmoklolw", de: "chisinau"},
		{pt: "8abebfadc966b3c2", key: "SECONDKE", de: "ATBACKSU"},
		{pt: "3b566979fe9440f2", key: "whygodwh", de: "LMAOCIPH"},
	}

	for _, data := range table {
		des := GetDES()

		des.KeyInit(data.key)
		_, str := des.Decrypt(data.pt)

		if str != data.de {
			t.Errorf(" The text %v : The decryption: %v , \n THe correct decryption: %v \n", data.pt, str, data.de)
		}

	}
	log.Println("Decryption tests passed ")
}
