package main

import (
	"fmt"

	ciphers "github.com/senpainikolay/CS-sem5/Lab1/Ciphers"
)

func main() {
	c := ciphers.CaesarClasic{}
	fmt.Println(c.Encrypt("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG", 23))
	fmt.Println(c.Decrypt("QEB NRFZH YOLTK CLU GRJMP LSBO QEB IXWV ALD", 23))
}
