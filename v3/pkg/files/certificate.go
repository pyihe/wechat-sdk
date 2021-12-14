package files

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"strings"

	"github.com/pyihe/go-pkg/errors"
)

// LoadRSAPrivateKey 加载RSA PRIVATE KEY
func LoadRSAPrivateKey(file string) (privateKey *rsa.PrivateKey, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	block, _ := pem.Decode(data)
	if block.Type != "PRIVATE KEY" {
		err = errors.New("证书类型必须是PRIVATE KEY!")
		return
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	var ok bool
	privateKey, ok = key.(*rsa.PrivateKey)
	if !ok {
		err = errors.New("请提供RSA私钥文件!")
	}
	return
}

// LoadRSAPublicKey 加载RSA PUBLIC KEY
func LoadRSAPublicKey(file string) (publicKey *rsa.PublicKey, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	block, _ := pem.Decode(data)
	if block.Type != "PUBLIC KEY" {
		err = errors.New("证书类型必须是PUBLIC KEY!")
		return
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	var ok bool
	publicKey, ok = key.(*rsa.PublicKey)
	if !ok {
		err = errors.New("请提供RSA公钥文件!")
	}
	return
}

// LoadCertificate 加载证书
func LoadCertificate(file string) (certificate *x509.Certificate, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	block, _ := pem.Decode(data)
	certificate, err = x509.ParseCertificate(block.Bytes)
	return
}

// LoadRSAPublicKeyWithSerialNo 加载本地RSA公钥，同时返回证书序列号
func LoadRSAPublicKeyWithSerialNo(file string) (serialNo string, publicKey *rsa.PublicKey, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	block, _ := pem.Decode(data)
	if block.Type != "PUBLIC KEY" {
		err = errors.New("证书类型必须是PUBLIC KEY!")
		return
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return
	}
	serialNo = strings.ToUpper(cert.SerialNumber.Text(16))
	publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		err = errors.New("invalid certificate!")
		return
	}
	return
}

// UnmarshalRSAPublicKey 反序列化RSA PublicKey, 同时返回证书序列号
func UnmarshalRSAPublicKey(key []byte) (serialNo string, publicKey *rsa.PublicKey, err error) {
	block, _ := pem.Decode(key)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return
	}
	serialNo = strings.ToUpper(cert.SerialNumber.Text(16))
	publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		err = errors.New("invalid certificate!")
		return
	}
	return
}

// UnmarshalCertificate 反序列化data到证书
func UnmarshalCertificate(data []byte) (serialNo string, cert *x509.Certificate, err error) {
	block, _ := pem.Decode(data)
	cert, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		return
	}
	serialNo = strings.ToUpper(cert.SerialNumber.Text(16))
	return
}
