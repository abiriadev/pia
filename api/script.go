package api

import "strconv"

type Script struct {
	Lines []string
}

type rawScript struct {
	S []struct {
		Text string `json:"text"`
	} `json:"s"`
}

func GetScript(id int) (rawScript, error) {
	s, err := getJson[rawScript](endPoint + strconv.Itoa(id))
	if err != nil {
		return s, err
	}

	return s, nil
}
