package crypto

type Crypto interface {
	Encrypt(key, plainText string) (string, error)
	Decrypt(key, cipherText string) (string, error)
}
