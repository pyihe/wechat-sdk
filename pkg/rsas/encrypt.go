package rsas

import (
	"crypto"

	"github.com/pyihe/secret"
	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
)

// EncryptOAEP RSA EncryptOAEP 加密敏感信息
func EncryptOAEP(cipher secret.Cipher, data interface{}) (cipherText string, err error) {
	if cipher == nil {
		err = errors.ErrNoCipher
		return
	}
	return cipher.RSAEncryptToString(data, crypto.SHA1, secret.RSAEncryptTypeOAEP, nil)
}

// DecryptOAEP RSA EncryptOAEP 解密
func DecryptOAEP(cipher secret.Cipher, data interface{}) (plainText []byte, err error) {
	if cipher == nil {
		err = errors.ErrNoCipher
		return
	}
	return cipher.RSADecrypt(data, crypto.SHA1, secret.RSAEncryptTypeOAEP, nil)
}
