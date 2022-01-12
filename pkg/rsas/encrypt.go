package rsas

import (
	"crypto"

	"github.com/pyihe/go-pkg/errors"

	"github.com/pyihe/secret"
)

// EncryptOAEP RSA EncryptOAEP 加密敏感信息
func EncryptOAEP(cipher secret.Cipher, data interface{}) (cipherText string, err error) {
	if cipher == nil {
		err = errors.New("未初始化的cipher!")
		return
	}
	return cipher.RSAEncryptToString(data, crypto.SHA1, secret.RSAEncryptTypeOAEP, nil)
}

// DecryptOAEP RSA EncryptOAEP 解密
func DecryptOAEP(cipher secret.Cipher, data interface{}) (plainText []byte, err error) {
	if cipher == nil {
		err = errors.New("未初始化的cipher!")
		return
	}
	return cipher.RSADecrypt(data, crypto.SHA1, secret.RSAEncryptTypeOAEP, nil)
}
