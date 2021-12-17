package rsas

import (
	"crypto"

	"github.com/pyihe/secret"
)

// EncryptOAEP RSA EncryptOAEP 加密敏感信息
func EncryptOAEP(cipher secret.Cipher, data interface{}) (cipherText string, err error) {
	return cipher.RSAEncryptToString(data, crypto.SHA1, secret.RSAEncryptTypeOAEP, nil)
}
