package api

import (
	"strconv"
	"strings"
)

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
	res, err := getJson[rawScript](endPointScript + strconv.Itoa(id))
	if err != nil {
		return script, err
	}

	for _, line := range res.S {
		script.Lines = append(script.Lines, strings.ReplaceAll(line.Text, "&nbsp;", " "))
	}

	return script, nil
}

func (script *Script) Textify() string {
	return strings.Join(script.Lines, "")
}
