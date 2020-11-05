package util

import (
	"bytes"
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

var Timeout = 600 * time.Millisecond
var roundtrip = &http.Transport{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
	// DisableKeepAlives: true,
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext,
	MaxIdleConns:          300,
	MaxIdleConnsPerHost:   50,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

var Client = &http.Client{Timeout: Timeout, Transport: roundtrip}

func HttpPOST(url string, header *http.Header, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header = *header
	return Client.Do(req)
}

func HttpGET(url string, header *http.Header) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = *header
	return Client.Do(req)
}