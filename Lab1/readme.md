# Laboratory Work 1. Classical Ciphers.

### Course: Cryptography & Security
### Author: Nicolae Gherman

----

## Theory
In the classical cryptography the original data or the plain text is transformed into the encoded format or cipher text so that we can transmit this data through insecure communication channels. 
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

	for _, ch := range message {
		if ch == ' ' {
			encrypted += " "
			continue
		}

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

	for _, ch := range message {
		if ch == ' ' {
			decrypted += " "
			continue
		}

		deCh := math.Mod(float64(int(unicode.ToUpper(ch))-65-key+26), 26)
		decrypted += string(int(deCh) + 65)
	}

	return decrypted
} 
``` 

* If needed, screenshots.


## Conclusions / Screenshots / Results
