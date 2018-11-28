package http

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpMethod(method string, url string, content []byte, headparameters ...string) (string, error) {
	if len(headparameters)%2 == 1 {
		return "", errors.New("parameters error")
	}
	data := strings.NewReader(string(content))
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(headparameters)-1; i = i + 2 {
		req.Header.Add(headparameters[i], headparameters[i+1])
	}
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
