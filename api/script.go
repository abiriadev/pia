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

func GetScript(id int) (Script, error) {
	var script Script
	res, err := getJson[rawScript](endPoint + strconv.Itoa(id))
	if err != nil {
		return script, err
	}

	for _, line := range res.S {
		script.Lines = append(script.Lines, line.Text)
	}

	return script, nil
}
