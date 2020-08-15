package dota2api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

//http get
func (d Dota2) Get(u string) ([]byte, error) {
	var body []byte

	request, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return body, err
	}
	resp, err := d.client.Do(request)
	if err != nil {
		return body, err
	}
	defer func() { _ = resp.Body.Close() }()

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

func ArrayIntToStr(arr []uint64) string {
	var strArr string
	strArr = "["

	for i, v := range arr {
		strArr += strconv.FormatUint(v, 10)
		if i+1 < len(arr) {
			strArr += ","
		}
	}

	return strArr + "]"
}
