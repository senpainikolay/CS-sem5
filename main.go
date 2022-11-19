package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	asymmetric "github.com/senpainikolay/CS-sem5/asymmetric"
	classicalCiphers "github.com/senpainikolay/CS-sem5/classicalCiphers"
	serverDBsimulation "github.com/senpainikolay/CS-sem5/db_simulation/server"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//RunAsymmetricCipher()
	// RunClassicalclassicalCiphers()
	serverDBsimulation.RunDBSimulationServer()

}

func RunAsymmetricCipher() {
	asymmetricCph := asymmetric.RSA{}
	en := asymmetric.CipherEncrypt(&asymmetricCph, "Test :-)")
	de := asymmetric.CipherDecrypt(&asymmetricCph, en)
	log.Println(en)
	log.Println(de)
}

func RunClassicalclassicalCiphers() {
	c := classicalCiphers.CaesarClasic{}
	fmt.Println("Caesar Clasic: ")
	fmt.Println("Encrpytion: ")
	fmt.Println(c.Encrypt("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG", 23))
	fmt.Println("Decryption:")
	fmt.Println(c.Decrypt("QEB NRFZH YOLTK CLU GRJMP LSBO QEB IXWV ALD", 23))
	fmt.Println()

	fmt.Println("Caesar with permutaions: ")
	cP := classicalCiphers.CaesarPermutation{SecretWord: "Pneumonoultramicroscopicsilicovolcanoconiosis"}
	cP.InitializeAlphabet()
	fmt.Println("Encrpytion: ")
	fmt.Println(cP.Encrypt("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG", 24))
	fmt.Println("Decryption:")
	fmt.Println(cP.Decrypt("OFP JNRAH SLUKZ BUQ GNEYI UCPL OFP MTXW VUD", 24))
	fmt.Println()

	fmt.Println("Vigenere cipher:")
	v := classicalCiphers.Vigenere{}
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
	cPf := classicalCiphers.Playfair{Msg: msg, Key: key}
	cPf.Init()
	fmt.Println("Encrpytion: ")
	fmt.Println(cPf.Encrypt())
	fmt.Println()
	fmt.Println("Playfair cipher:")
	msg = "ijfsknkqijlffgijfss"
	fmt.Printf("The message: %v\n", msg)
	fmt.Printf("The Key: %v\n", key)
	cPf2 := classicalCiphers.Playfair{Msg: msg, Key: key}
	cPf2.Init()
	fmt.Println("Decryption:")
	fmt.Println(cPf2.Decrypt())
}
