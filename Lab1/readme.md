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

* Clasic Caesar. Shifting the key by ASCII table on encryption and decryption.  

*Encryption: 
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
*Decryption: 
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

* Caesar with Secret Word Permutation. The unique letters from secret word are permutated to the begining of the alphabet and then the rest of letters are added. The encryption and decryption are based on indexes of the newly formed string alphabet  and follows the logic from Caesar above.
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
* If needed, screenshots.


## Conclusions / Screenshots / Results
