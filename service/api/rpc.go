package main

import (
	"net/http"
	"time"
	"io"
	"io/ioutil"
	"fmt"
)

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
	Transport: &http.Transport{},
}

func httpReq(method string, path string, body io.ReadCloser) ([]byte, error) {
	url := "http://localhost:9113" + path
	req, err := http.NewRequest(method, url, body) // if body is also an io.Closer then client.Do will close it
	if err != nil {
		return nil, err
	}
	res, err := httpDo(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func checkUser(authHeader string) error {
	url := "http://localhost:9112"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", authHeader)
	_, err = httpDo(req)
	if err != nil {
		return err
	}
	return nil
}

func httpDo(req *http.Request) ([]byte, error) {
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%d %s", res.StatusCode, body)
	}

	return body, nil
}
