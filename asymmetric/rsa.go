package asymmetric

import (
	"log"
	"math"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

type RSA struct {
	PublicKey  PublicKey
	PrivateKey int64
	// Saves the length of AsciiCode converted letter in order to decrypt the string properly.
	LenOfAsciiCodes []int
}

type PublicKey struct {
	n int
	e int
}

func (c *RSA) GenerateKeys() {
	rand.Seed(time.Now().UTC().UnixNano())
	p := GenerateFirstPrimeNumbersInRange(50, 80)
	q := GenerateFirstPrimeNumbersInRange(80, 100)
	e := 2
	phi := (p - 1) * (q - 1)
	for e < phi {
		if gcd(e, phi) == 1 {
			break
		}
		e = e + 1
	}
	c.PublicKey = PublicKey{(p * q), e}
	// math.Mod(float64(e*c.PrivateKey), float64(phi))) == 1 Condition
	k := 2
	c.PrivateKey = int64((1 + (k * phi)) / e)

}

func (cph *RSA) Encrypt(msg string) []int64 {
	var en []int64
	var toEncrypt []int
	k := 0
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
	for _, c := range toEncrypt {
		exp := math.Pow(float64(c), float64(cph.PublicKey.e))
		mod := math.Mod(exp, float64(cph.PublicKey.n))
		en = append(en, int64(mod))
	}
	return en

}

func (cph *RSA) Decrypt(msg []int64) string {
	var de string
	var deFinal string

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

	k := 0
	for _, i := range cph.LenOfAsciiCodes {
		AsciiCode, _ := strconv.Atoi(de[k:i])
		deFinal += string(rune(AsciiCode))
		k = i
	}

	return deFinal
}

func GenerateFirstPrimeNumbersInRange(num1, num2 int) int {
	if num1 < 2 || num2 < 2 {
		log.Fatalln("Numbers must be greater than 2.")
		return -1
	}
	for num1 <= num2 {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num1))); i++ {
			if num1%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			return num1
		}
		num1++
	}

	log.Fatalln("Coulnt find a prime number is this range")
	return -1
}

func gcd(a int, h int) int {
	temp := 0
	for {
		temp = int(math.Mod(float64(a), float64(h)))
		if temp == 0 {
			return int(h)
		}
		a = h
		h = temp
	}

}
