package api

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Toc struct {
	Chapters []Chapter
}

type Chapter struct {
	Name string
	Id   int
	Type string
}

type rawToc struct {
	Chapters []rawChapter `goquery:"tr"`
}

type rawChapter struct {
	Name string `goquery:"td.font12 > b,html"`
	Id   string `goquery:"tr > td:first-child,html"`
	Type string `goquery:"td.font12 > b > span"`
}

func GetToc(id int) (Toc, error) {
	rawToc, err := postHtml[rawToc](endPointToc, url.Values{
		"novel_no": []string{strconv.Itoa(id)},
	})
	if err != nil {
		return *new(Toc), err
	}

	toc := Toc{
		Chapters: make([]Chapter, len(rawToc.Chapters)),
	}

	for i, chapter := range rawToc.Chapters {
		nameDoc, err := goquery.NewDocumentFromReader(strings.NewReader(chapter.Name))
		if err != nil {
			return *new(Toc), err
		}

		Name := strings.TrimSpace(nameDoc.Children().Children().Eq(1).Contents().Last().Text())

		idDoc, err := goquery.NewDocumentFromReader(strings.NewReader(chapter.Id))
		if err != nil {
			return *new(Toc), err
		}

		Id, err := strconv.Atoi(
			strings.TrimSuffix(
				strings.TrimPrefix(idDoc.Find("div").Get(0).Attr[1].Key, "novel_on_"),
				"\"",
			),
		)
		if err != nil {
			return *new(Toc), err
		}

		toc.Chapters[i] = Chapter{
			Name: Name,
			Id:   Id,
			Type: chapter.Type,
		}
	}

	return toc, nil
}
