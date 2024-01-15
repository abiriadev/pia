package api

import (
	"net/url"
	"strconv"
)

type Toc struct {
	Chapters []string `goquery:"#episode_list_viewer > table > tr > .ion-bookmark"`
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
