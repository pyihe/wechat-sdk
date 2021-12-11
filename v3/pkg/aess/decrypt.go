package aess

import "github.com/pyihe/secret"

// DecryptAES128CBCPKCS7 AES-128-CBC PKCS#7 解密
func DecryptAES128CBCPKCS7(cipher secret.Cipher, cipherText string, key, iv []byte) (plainText []byte, err error) {
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
