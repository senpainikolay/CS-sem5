# Laboratory Work 3. Asymmetric Ciphers

### Course: Cryptography & Security
### Author: Nicolae Gherman

----

## Theory 
Asymmetric cryptography, also known as public-key cryptography, is a process that uses a pair of related keys: one public key and one private key,  to encrypt and decrypt a message and protect it from unauthorized access or use. <br> 
The idea of RSA, which I have implemented, is based on the fact that it is difficult to factorize a large integer. The public key consists of two numbers where one number is a multiplication of two large prime numbers. And private key is also derived from the same two prime numbers. So if somebody can factorize the large number, the private key is compromised. Therefore encryption strength totally lies on the key size and if we double or triple the key size, the strength of encryption increases exponentially. 


## Objectives:

* Get familiar with asymmetric ciphers. 
* Implement the Genereting Keys, Encryption and Decryption methods of a choosen algorithm,  RSA in this case.

## Implementation description
The implementaions follows 3 steps, The Public and Private key generations, Encryption and Decryption. 
 
1. Keys Generation. 
Firstly, 2  prime numbers are considered. 
```
p := GenerateFirstPrimeNumbersInRange(50, 80)
q := GenerateFirstPrimeNumbersInRange(80, 100) 
``` 
The Euler’s Totient Function is defined through  phi:  
```  
phi := (p - 1) * (q - 1)
```
Then, the 'e' is choosed relatively small,  it should not be a factor of p*q and should follow the condition: 1 < e < phi && it should not have common divisors expect 1 with 'phi'.
```
e := 2
	for e < phi {
		if gcd(e, phi) == 1 {
			break
		}
		e = e + 1
	} 


```  
The next step is to calculate Private key such (e*d)mod phi == 1 
The ( p*q, e ) would form the Public Key 
``` 
c.PublicKey = PublicKey{(p * q), e}
k := 2
c.PrivateKey = int64((1 + (k * phi)) / e)
``` 

2. Encryption:  

It uses the actually Public Key and follow the fomrula: encrypted = ( msg^e) mod n, n:p*q. 
The output is actually an array of intergers in my code which is taken further in Decryption
```
for _, c := range toEncrypt { // The parts of ascii Code numbers in the code.
		exp := math.Pow(float64(c), float64(cph.PublicKey.e))
		mod := math.Mod(exp, float64(cph.PublicKey.n))
		en = append(en, int64(mod))
	} 
``` 
 

 3. Decryption:   
 It uses the actually Private Key and follow the fomrula: decrypted = ( encrypted^private_key) mod n, n:p*q.  
 A special module is used to operate with Big numbers.

```
	for _, c := range msg {
		//exp := math.Pow(float64(c), float64(cph.PrivateKey))
		var exp, e = big.NewInt(c), big.NewInt(cph.PrivateKey)
		// msg^e
		exp.Exp(exp, e, nil)
		// convering msg^e for modular operations
		numMod1 := new(big.Int)
		numMod1.SetBytes(exp.Bytes())
		// converting n to modular operation
		numMod2 := new(big.Int)
		convertToBigNum := big.NewInt(int64(cph.PublicKey.n))
		numMod2.SetBytes(convertToBigNum.Bytes())
		// (msg^e) mod n
		mod := new(big.Int)
		mod = mod.Mod(numMod1, numMod2)
		de += strconv.FormatInt(mod.Int64(), 10)

	} 
```  
---  

To avoid Overflow, I have applied encryption on single numbers from Ascii Codes from letters/inputs.  
The Cipher structure has an attribute (LenOfAsciiCode) which represents an integer array to keep the original lenghts of ascii codes. It is applied in encryption method and used in Decrpytion method. 
```  
// Encryption 
for _, c := range msg {
		// Ascii Letter Code converted  to string value (to avoid overflow!)
		asciiConverted := strconv.FormatInt(int64(c), 10)
		// Saving the lens of those ascii code letters
		k += len(asciiConverted)
		cph.LenOfAsciiCodes = append(cph.LenOfAsciiCodes, k)
		// parsing the string and encpryting converted numbers(0-9)

		for _, l := range asciiConverted {
			Li, _ := strconv.Atoi(string(l))
			toEncrypt = append(toEncrypt, Li)
		}
	} 


``` 

``` 
// Decryption
k := 0
	for _, i := range cph.LenOfAsciiCodes {
		AsciiCode, _ := strconv.Atoi(de[k:i])
		deFinal += string(rune(AsciiCode))
		k = i
	}

```


## The Screenshot Regarding the runing tests form res_test.go:  
![Screenshot](output.png)




