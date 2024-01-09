package api

import "strconv"

type Script struct {
	Lines []string
}

func GetScript(id int) (Script, error) {
	s, err := getJson[Script](endPoint + strconv.Itoa(id))
	if err != nil {
		return s, err
	}

	return s, nil
}
