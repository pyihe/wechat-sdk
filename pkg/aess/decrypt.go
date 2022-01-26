package aess

import (
	"github.com/pyihe/secret"
	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
)

// DecryptAEADAES256GCM RSA-EncryptOAEP解密
func DecryptAEADAES256GCM(cipher secret.Cipher, key string, cipherData interface{}, associateData, nonce string) (plainText []byte, err error) {
	if cipher == nil {
		err = errors.ErrNoCipher
		return
	}
	var decryptRequest = &secret.SymRequest{
		CipherData:  cipherData,
		Key:         []byte(key),
		Type:        secret.SymTypeAES,
		ModeType:    secret.BlockModeGCM,
		PaddingType: secret.PaddingTypeNoPadding,
		AddData:     []byte(associateData),
		Nonce:       []byte(nonce),
	}
	return cipher.SymDecrypt(decryptRequest)
}

// DecryptAES128CBCPKCS7 AES-128-CBC PKCS#7 解密
func DecryptAES128CBCPKCS7(cipher secret.Cipher, cipherText interface{}, key, iv []byte) (plainText []byte, err error) {
	if cipher == nil {
		err = errors.ErrNoCipher
		return
	}
	var request = &secret.SymRequest{
		CipherData:  cipherText,
		Key:         key,
		Type:        secret.SymTypeAES,
		ModeType:    secret.BlockModeCBC,
		PaddingType: secret.PaddingTypePKCS7,
		Iv:          iv,
	}
	plainText, err = cipher.SymDecrypt(request)
	return
}
