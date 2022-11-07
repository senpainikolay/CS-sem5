package main

import (
	"fmt"
	"log"

	ciphers "github.com/senpainikolay/CS-sem5/Ciphers"
	asymmetric "github.com/senpainikolay/CS-sem5/asymmetric"
)

func main() {
	RunAsymmetricCipher()
	RunClassicalCiphers()
}

func RunAsymmetricCipher() {
	asymmetricCph := asymmetric.RSA{}
	en := asymmetric.CipherEncrypt(&asymmetricCph, "Test :-)")
	de := asymmetric.CipherDecrypt(&asymmetricCph, en)
	log.Println(en)
	log.Println(de)
}

func RunClassicalCiphers() {
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

	fmt.Println("Playfair cipher:")
	msg = "hammerhellohamm"
	key = "FAFKEKLOL"
	fmt.Printf("The message: %v\n", msg)
	fmt.Printf("The Key: %v\n", key)
	cPf := ciphers.Playfair{Msg: msg, Key: key}
	cPf.Init()
	fmt.Println("Encrpytion: ")
	fmt.Println(cPf.Encrypt())
	fmt.Println()
	fmt.Println("Playfair cipher:")
	msg = "ijfsknkqijlffgijfss"
	fmt.Printf("The message: %v\n", msg)
	fmt.Printf("The Key: %v\n", key)
	cPf2 := ciphers.Playfair{Msg: msg, Key: key}
	cPf2.Init()
	fmt.Println("Decryption:")
	fmt.Println(cPf2.Decrypt())
}
