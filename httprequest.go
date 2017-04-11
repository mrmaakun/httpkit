package httpkit

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"strconv"
)

func httpRequest(method string, url string, headers map[string]string, payload []byte) (*http.Response, error) {

	var req = &http.Request{}
	var err error

	log.Println("Request URL: " + url)

	if payload != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(payload))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	log.Println("Response Code: " + strconv.Itoa(resp.StatusCode))

	// Throw an error if the response is over 400
	if resp.StatusCode >= 400 {
		err = errors.New("ERROR Status Code is " + strconv.Itoa(resp.StatusCode))
	}

	return resp, err
}
