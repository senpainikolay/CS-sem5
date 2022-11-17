package client_simulation

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"

	hashArgon2 "github.com/senpainikolay/CS-sem5/db_simulation/argon2"
	inMemoryDB "github.com/senpainikolay/CS-sem5/db_simulation/database"
)

type ClientInterfaceSimulation struct {
	db            *inMemoryDB.Database
	RSAPrivateKey *rsa.PrivateKey
	RSAlabel      []byte
}

func GetClientInterfaceSimulation() *ClientInterfaceSimulation {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	return &ClientInterfaceSimulation{
		db:            inMemoryDB.GetDB(),
		RSAPrivateKey: privateKey,
		RSAlabel:      []byte("OAEP Encrypted"),
	}
}

func (c *ClientInterfaceSimulation) RegisterCredentials(username string, pw string) bool {

	hashedPassword := hashArgon2.GetTheHashOnPassword(pw)
	en, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &c.RSAPrivateKey.PublicKey, hashedPassword, c.RSAlabel)
	if err != nil {
		return true
	}

	c.db.RegisterUser(username, en)
	return false

}

func (c *ClientInterfaceSimulation) LogInCredentials(username string, pw string) bool {

	hashedPassword := hashArgon2.GetTheHashOnPassword(pw)
	db_pass := c.db.GetUserPassword(username)
	de, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, c.RSAPrivateKey, db_pass, c.RSAlabel)
	if err != nil {
		return true
	}
	res := bytes.Compare(hashedPassword, de)
	if res == 0 {
		return false
	}

	return true

}

func (c *ClientInterfaceSimulation) DeleteCredentials(username string, pw string) bool {

	hashedPassword := hashArgon2.GetTheHashOnPassword(pw)
	db_pass := c.db.GetUserPassword(username)
	de, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, c.RSAPrivateKey, db_pass, c.RSAlabel)
	if err != nil {
		return true
	}
	res := bytes.Compare(hashedPassword, de)
	if res != 0 {
		log.Println("NOT CORRECT DATA!")
		return true

	}

	c.db.Delete(username)
	return false

}
