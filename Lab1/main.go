package main

import (
	"fmt"

	ciphers "github.com/senpainikolay/CS-sem5/Lab1/Ciphers"
)

func main() {
	c := ciphers.CaesarClasic{}
	fmt.Println("Caesar Clasic: ")
	fmt.Println(c.Encrypt("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG", 23))
	fmt.Println(c.Decrypt("QEB NRFZH YOLTK CLU GRJMP LSBO QEB IXWV ALD", 23))

	fmt.Println("Caesar with permutaions: ")
	cP := ciphers.CaesarPermutation{SecretWord: "Pneumonoultramicroscopicsilicovolcanoconiosis"}
	cP.InitializeAlphabet()
	fmt.Println(cP.Encrypt("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG", 24))
	fmt.Println(cP.Decrypt("OFP JNRAH SLUKZ BUQ GNEYI UCPL OFP MTXW VUD", 24))

}
