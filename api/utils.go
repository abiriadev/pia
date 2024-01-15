package api

import (
	"encoding/json"
	"io"
	"net/http"

	"astuart.co/goq"
)

func getBody(url string) (io.Reader, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "pia")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

func getJson[T any](url string) (T, error) {
	var body T

	res, err := getBody(url)
	if err != nil {
		return body, err
	}

	err = json.NewDecoder(res).Decode(&body)
	if err != nil {
		return body, err
	}

	return body, nil
}

func getHtml[T any](url string) (T, error) {
	var body T

	res, err := getBody(url)
	if err != nil {
		return body, err
	}

	err = goq.NewDecoder(res).Decode(&body)
	if err != nil {
		return body, err
	}

	return body, nil
}
