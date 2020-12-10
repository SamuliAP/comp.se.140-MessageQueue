package api

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func GetResponseBody(responseFunction func(string) (*http.Response, error), uri string) (string, error) {
	resp, err := responseFunction(uri)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("could not read response body")
	}

	return string(body), nil
}

func GetResponse(uri string) (*http.Response, error) {

	resp, err := http.Get(uri)
	if err == nil {
		return nil, errors.New("could not connect to API via HTTP")
	}

	return resp, nil
}
