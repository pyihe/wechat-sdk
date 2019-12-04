package util

import (
	"crypto/tls"
	"encoding/pem"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
)

func P12ToPem(p12Path string, password string) (*tls.Certificate, error) {
	p12, err := ioutil.ReadFile(p12Path)
	if err != nil {
		return nil, err
	}
	blocks, err := pkcs12.ToPEM(p12, password)
	if err != nil {
		return nil, err
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}
	pemCert, err := tls.X509KeyPair(pemData, pemData)
	return &pemCert, err
}
