package asymmetric

type AsymmetricCipher interface {
	GenerateKeys()
	Encrypt(string) []int64
	Decrypt([]int64) string
}

func CipherEncrypt(c AsymmetricCipher, msg string) []int64 {
	c.GenerateKeys()
	return c.Encrypt(msg)
}
func CipherDecrypt(c AsymmetricCipher, en []int64) string {
	c.GenerateKeys()
	return c.Decrypt(en)
}
