package api

import (
	"strconv"
)

type Novel struct {
	Title string
}

func GetNovel(id int) (Novel, error) {
	novel, err := getHtml[Novel](endPointNovel + strconv.Itoa(id))
	if err != nil {
		return *new(Novel), err
	}

	return novel, nil
}
