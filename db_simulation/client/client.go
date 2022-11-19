package client_simulation

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"
	mathRand "math/rand"

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

func (c *ClientInterfaceSimulation) RegisterCredentials(username string, pw string) string {

	hashedPassword := hashArgon2.GetTheHashOnText(pw)
	c.db.RegisterUser(username, hashedPassword)
	token := randStr(10)
	hashedToken := hashArgon2.GetTheHashOnText(token)
	en, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &c.RSAPrivateKey.PublicKey, hashedToken, c.RSAlabel)
	if err != nil {
		log.Fatal(" RSA Encrytion failed ")
	}
	c.db.UserToken[username] = en
	return token

}

func (c *ClientInterfaceSimulation) LogInCredentials(username string, pw string, tk string) string {

	hashedPassword := hashArgon2.GetTheHashOnText(pw)
	db_pass := c.db.GetUserPassword(username)
	res := bytes.Compare(hashedPassword, db_pass)
	if res != 0 {
		return "wrong password or user does not exist"
	}
	hashedToken := hashArgon2.GetTheHashOnText(tk)
	de, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, c.RSAPrivateKey, c.db.UserToken[username], c.RSAlabel)
	if err != nil {
		return "Error on Authentication"
	}
	res = bytes.Compare(hashedToken, de)
	if res != 0 {
		return "Unsucessful Authentication"
	}

	return "Success Log In"

}

func randStr(n int) string {
	charset := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]byte, n)
	for i := range b {
		b[i] = charset[mathRand.Intn(len(charset))]
	}
	return string(b)
}
