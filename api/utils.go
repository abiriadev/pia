package api

import (
	"encoding/json"
	"net/http"
)

func getJson[T any](url string) (T, error) {
	var body T

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return body, err
	}

	req.Header.Add("User-Agent", "pia")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return body, err
	}

	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return body, err
	}

	return body, nil
}
