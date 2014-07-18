package dota2api

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

//http get
func Get(u string) ([]byte, error) {
	var body []byte

	timeout := time.Duration(10) * time.Second

	transport := &http.Transport{
		ResponseHeaderTimeout: timeout,
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		},
		DisableKeepAlives: true,
	}

	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Get(u)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}

	return body, nil
}

func parseUrl(u string, param map[string]interface{}) (string, error) {
	ur, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	q := ur.Query()

	for k, v := range param {
		q.Set(k, fmt.Sprintf("%v", v))
	}

	ur.RawQuery = q.Encode()
	return ur.String(), nil
}
