# Laboratory Work 1. Classical Ciphers.

### Course: Cryptography & Security
### Author: Nicolae Gherman

----

## Theory
In the classical cryptography the original data or the plain text is transformed into the encoded format or cipher text so that we can transmit this data through insecure communication channels. <br />
So we begin from first popular methods of encryptions, the  symmetric ciphers studied below. In the symmetric cryptography a single key is used for encrypting and decryption the data. This encryption key is private key. This is the limitation of this encryption technique that this private key must be distributed only among the authorized sender and receiver. 

## Objectives:

* Get familiar with classical ciphers which mostly are symmetric. 
* Implement the Encryption and Decryption methods of those ciphers.

## Implementation description

* Clasic Caesar. <br /> Shifting the key by ASCII table on encryption and decryption.  

1.  Encryption: 
```
func (c *CaesarClasic) Encrypt(message string, key int) string {

	var encrypted string 
	//Code for spacing

		enCh := math.Mod(float64(int(unicode.ToUpper(ch))-65+key), 26)
		encrypted += string(int(enCh) + 65)
	}

	return encrypted
}
``` 
2. Decryption: 
``` 
func (c *CaesarClasic) Decrypt(message string, key int) string {

	var decrypted string 
	//Code for spacing


		deCh := math.Mod(float64(int(unicode.ToUpper(ch))-65-key+26), 26)
		decrypted += string(int(deCh) + 65)
	}

	return decrypted
} 
```  

* Caesar with Secret Word Permutation. <br />  The unique letters from secret word are permutated to the begining of the alphabet and then the rest of letters are added. The encryption and decryption are based on indexes of the newly formed string alphabet  and follows the logic from Caesar above. 
1. Permutation logic
``` 
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
``` 

* Vigenere Cipher <br />
After forming the keystream which would equal to the length of the plain text, it is used ASCII TABLE to perform the calculation
1. Encryption:  
 The plaintext(P) and keystream(K) are added modulo 26.  
 Ei = (Pi + Ki) mod 26
```  
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

```

2. Decryoption: <br /> 
Di = (Ei - Ki + 26) mod 26 

```
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
``` 

* Playfair cipher. <br /> 
After adjusting the alphabet to a custom one like shown in Caesar with permutation above, it is placed in a 5x5 grid ( i/j beiing the same letter). <br/> 

1. Then the message are splitted in  diagraphs corresponding to the use cases: 
```  
func (c *Playfair) InitSteps() {
	// two letter pair elimination
	for {
		if checkDublicates(c.Msg) {
			break
		}
		for i := 2; i < len(c.Msg); i += 2 {
			kek := c.Msg[i-2 : i]
			if kek[0] == kek[1] {
				c.Msg = Insert(c.Msg, "x", i-1)
				break
			}
		}
	}
	// adding extra bogus letter if there is a non paired diagraph
	if math.Mod(float64(len(c.Msg)), 2) == 1 {
		c.Msg += "z"

	}
	// forming the actual Diagraphs
	for i := 0; i < len(c.Msg); i += 2 {
		c.Diagraphs = append(c.Diagraphs, c.Msg[i:i+2])
	}

}
``` 
2. The 3 casses for: <br /> 
2.1 Encryption:  
``` 
func (c *Playfair) Encrypt() string {
	c.InitSteps()

	var encrypted string

	for _, elem := range c.Diagraphs {
		i1, j1 := c.GetGridPos(string(elem[0]))
		i2, j2 := c.GetGridPos(string(elem[1]))

		// case 1 : in the same column
		if j1 == j2 {
			en1 := int(math.Mod(float64(i1+1), 5))
			en2 := int(math.Mod(float64(i2+1), 5))
			encrypted += c.grid[en1][j1]
			encrypted += c.grid[en2][j2]
			continue
		}
		// case 2 : in the same row
		if i1 == i2 {
			en1 := int(math.Mod(float64(j1+1), 5))
			en2 := int(math.Mod(float64(j2+1), 5))
			encrypted += c.grid[i1][en1]
			encrypted += c.grid[i2][en2]
			continue
		}

		// case 3: rectangle
		encrypted += c.grid[i1][j2]
		encrypted += c.grid[i2][j1]

	}

	return encrypted

}
```


2.2: Decryption:  
``` 
func (c *Playfair) Decrypt() string {

	c.InitSteps()

	var decrypted string

	for _, elem := range c.Diagraphs {
		i1, j1 := c.GetGridPos(string(elem[0]))
		i2, j2 := c.GetGridPos(string(elem[1]))

		// case 1 : in the same column
		if j1 == j2 {
			en1 := i1 - 1
			en2 := i2 - 1
			if en1 == -1 {
				en1 = 4
			}

			if en2 == -1 {
				en2 = 4
			}
			decrypted += c.grid[en1][j1]
			decrypted += c.grid[en2][j2]
			continue
		}
		// case 2 : in the same row
		if i1 == i2 {
			en1 := j1 - 1
			en2 := j2 - 1
			if en1 == -1 {
				en1 = 4
			}
			if en2 == -1 {
				en2 = 4
			}
			decrypted += c.grid[i1][en1]
			decrypted += c.grid[i2][en2]
			continue
		}

		// case 3: rectangle
		decrypted += c.grid[i1][j2]
		decrypted += c.grid[i2][j1]

	}
	fnDe := decrypted

	for i := 2; i < len(decrypted); i++ {
		if decrypted[i-2] == decrypted[i] {
			fnDe = Rm(fnDe, i-1)

		}
	}

	return fnDe

} 
``` 




## The Screenshot Regarding the outputs:  
![alt text](https://github.com/senpainikolay/CS-sem5/tree/main/Lab1/output.png?raw=true)





