package rsas

import (
	"crypto"

	"github.com/pyihe/secret"
	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
)

// SignSHA256WithRSA SHA-256 with RSA 签名
func SignSHA256WithRSA(cipher secret.Cipher, data interface{}) (signature string, err error) {
	if cipher == nil {
		err = errors.ErrNoCipher
		return
	}
	signature, err = cipher.RSASignToString(data, secret.RSASignTypePKCS1v15, crypto.SHA256)
	return
}

// VerifySHA256WithRSA 验证SHA256-RSA签名
func VerifySHA256WithRSA(cipher secret.Cipher, signText, plainText string) (err error) {
	if cipher == nil {
		err = errors.ErrNoCipher
		return
	}
	ok, err := cipher.RSAVerify(signText, plainText, secret.RSASignTypePKCS1v15, crypto.SHA256)
	if err != nil {
		return
	}
	if !ok {
		err = errors.ErrNoCipher
	}
	return
}
