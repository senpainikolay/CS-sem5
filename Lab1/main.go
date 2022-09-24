package main

import (
	"fmt"

	ciphers "github.com/senpainikolay/CS-sem5/Lab1/Ciphers"
)

func main() {
	c := ciphers.CaesarClasic{}
	fmt.Println("Caesar Clasic: ")
	fmt.Println("Encrpytion: ")
	fmt.Println(c.Encrypt("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG", 23))
	fmt.Println("Decryption:")
	fmt.Println(c.Decrypt("QEB NRFZH YOLTK CLU GRJMP LSBO QEB IXWV ALD", 23))
	fmt.Println()

	fmt.Println("Caesar with permutaions: ")
	cP := ciphers.CaesarPermutation{SecretWord: "Pneumonoultramicroscopicsilicovolcanoconiosis"}
	cP.InitializeAlphabet()
	fmt.Println("Encrpytion: ")
	fmt.Println(cP.Encrypt("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG", 24))
	fmt.Println("Decryption:")
	fmt.Println(cP.Decrypt("OFP JNRAH SLUKZ BUQ GNEYI UCPL OFP MTXW VUD", 24))
	fmt.Println()

	fmt.Println("Vigenere cipher:")
	v := ciphers.Vigenere{}
	msg := "Vigenere Cipher Encryption Decryption"
	key := "FAFKEK"
	fmt.Printf("The message: %v\n", msg)
	v.Initialize(key, msg)
	fmt.Println("Encrpytion: ")
	fmt.Println(v.Encrypt())
	fmt.Println()
	fmt.Println("Vigenere cipher:")
	msg = "AILOROWE HSTRJR JXGBDPYSSX IEHBCZYITX"
	fmt.Printf("The message: %v\n", msg)
	v.Initialize(key, msg)
	fmt.Println("Decryption:")
	fmt.Println(v.Decrypt())
	fmt.Println()

}
