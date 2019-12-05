package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

//HMAC-SHA256签名方式
func SignHMACSHA256(s string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//MD5签名方式
func SignMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//填充方式为PKCS7的AES128CBC解密， 用于小程序敏感信息解密
func AES128CBCDecrypt(encryptData, key, iv []byte) (origData []byte, err error) {
	defer func() {
		if pErr := recover(); pErr != nil {
			err = errors.New("key is not match data")
		}
	}()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData = make([]byte, len(encryptData))
	blockMode.CryptBlocks(origData, encryptData)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

//PKCS7解填充
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//填充方式为PKCS7的AES256解密， 用于解密微信退款通知
func AES256ECBDecrypt(data, key []byte) (realData []byte, err error) {
	defer func() {
		if pErr := recover(); pErr != nil {
			err = errors.New("key is not match data")
		}
	}()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	if len(data)%blockSize != 0 {
		return nil, errors.New("block size cannot match data size")
	}

	realData = make([]byte, 0)
	text := make([]byte, 16)
	for len(data) > 0 {
		block.Decrypt(text, data)
		data = data[blockSize:]
		realData = append(realData, text...)
	}
	realData = PKCS7UnPadding(realData)
	return realData, nil
}
