package api

import (
	"net/url"
	"strconv"
)

type Toc struct {
	Chapters string `goquery:"td.font12 > b"`
}

func GetToc(id int) (Toc, error) {
	toc, err := postHtml[Toc](endPointToc, url.Values{
		"novel_no": []string{strconv.Itoa(id)},
	})
	if err != nil {
		return *new(Toc), err
	}

	return toc, nil
}
