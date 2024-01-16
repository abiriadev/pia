package api

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

	for i, chapter := range toc.Chapters {
		nameDoc, err := goquery.NewDocumentFromReader(strings.NewReader(chapter.Name))
		if err != nil {
			return toc, err
		}

		Name := strings.TrimSpace(nameDoc.Children().Children().Eq(1).Contents().Last().Text())

		idDoc, err := goquery.NewDocumentFromReader(strings.NewReader(chapter.Id))
		if err != nil {
			return toc, err
		}

		Id := idDoc.Find("div").Get(0).Attr[1].Key
		Id = strings.TrimSuffix(strings.TrimPrefix(Id, "novel_on_"), "\"")

		toc.Chapters[i] = Chapter{
			Name: Name,
			Id:   Id,
			Type: chapter.Type,
		}
	}

	return toc, nil
}
