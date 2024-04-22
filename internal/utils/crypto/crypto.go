package crypto

var KeySize int = 32

type Crypto interface {
	Encrypt(key, plainText string) (string, error)
	Decrypt(key, cipherText string) (string, error)
}
