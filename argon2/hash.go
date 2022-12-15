package hashing

import (
	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func GetTheHashOnText(password string) []byte {
	// Establish the parameters to use for Argon2.
	p := &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	// Pass the plaintext password and parameters to our generateFromPassword
	// helper function.
	salt := []byte{113, 75, 156, 151, 17, 222, 194, 185, 42, 59, 119, 32, 22, 216, 81, 1}
	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)
	return hash
}
