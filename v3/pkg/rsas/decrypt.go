package rsas

import "github.com/pyihe/secret"

// DecryptAEADAES256GCM RSA-EncryptOAEP解密
func DecryptAEADAES256GCM(cipher secret.Cipher, key, cipherText, associateData, nonce string) (plainText []byte, err error) {
	var decryptRequest = &secret.SymRequest{
		CipherData:  cipherText,
		Key:         []byte(key),
		Type:        secret.SymTypeAES,
		ModeType:    secret.BlockModeGCM,
		PaddingType: secret.PaddingTypeNoPadding,
		AddData:     []byte(associateData),
		Nonce:       []byte(nonce),
	}
	return cipher.SymDecrypt(decryptRequest)
}
