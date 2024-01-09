package api

import (
	"encoding/json"
	"net/http"
)

func getJson[T any](url string) (T, error) {
	var body T

	res, err := http.Get(url)
	if err != nil {
		return body, err
	}

	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return body, err
	}

	return body, nil
}
