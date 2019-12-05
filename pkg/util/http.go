package util

import (
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type PostRequest struct {
	Body io.Reader

	Url         string
	ContentType string
}

//向微信服务器发送POST请求
func PostToWx(req *PostRequest, result interface{}) error {
	if req == nil {
		return errors.New("have no PostRequest")
	}
	if req.Body == nil {
		return errors.New("body cannot be nil")
	}

	response, err := http.Post(req.Url, req.ContentType, req.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("http StatusCode: %v", response.StatusCode))
	}

	byteContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(byteContent, result)
	if err != nil {
		return err
	}
	return nil
}

//带证书的Post请求
func PostToWxWithCert(req *PostRequest, p12Cert *tls.Certificate, result interface{}) error {
	if p12Cert == nil {
		return errors.New("need p12Cert")
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{*p12Cert},
		},
		DisableCompression: true,
	}
	httpClient := http.Client{
		Transport: transport,
	}
	//发送请求
	response, err := httpClient.Post(req.Url, req.ContentType, req.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("http StatusCode: %v", response.StatusCode))
	}

	byteContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(byteContent, result)
	if err != nil {
		return err
	}
	return nil
}
