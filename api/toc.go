package api

import (
	"net/url"
	"strconv"
)

type Chapter struct {
	Name string `goquery:"td.font12 > b,html"`
	Id   string `goquery:"tr > td:first-child,html"`
	Type string `goquery:"td.font12 > b > span"`
}

type Toc struct {
	Chapters []Chapter `goquery:"tr"`
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
