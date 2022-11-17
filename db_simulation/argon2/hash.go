package hashing

import (
	"crypto/rand"
	"log"

	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func GetTheHashOnPassword(password string) []byte {
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
	hash, err := generateFromPassword(password, p)
	if err != nil {
		log.Fatal(err)
	}

	return hash
}

func generateFromPassword(password string, p *params) (hash []byte, err error) {
	// Generate a cryptographically secure random salt.
	salt, err := []byte{113, 75, 156, 151, 17, 222, 194, 185, 42, 59, 119, 32, 22, 216, 81, 1}, nil // generateRandomBytes(p.saltLength)
	if err != nil {
		return nil, err
	}

	// Pass the plaintext password, salt and parameters to the argon2.IDKey
	// function. This will generate a hash of the password using the Argon2id
	// variant.
	hash = argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	return hash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
