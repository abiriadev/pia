package api

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/otterize/lox"
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

func GetToc(id int, page int) (Toc, error) {
	rawToc, err := postHtml[rawToc](endPointToc, url.Values{
		"novel_no": []string{strconv.Itoa(id)},
		"page":     []string{strconv.Itoa(page)},
	})
	if err != nil {
		return *new(Toc), err
	}

	Chapters, err := lox.MapErr(rawToc.Chapters, func(chapter rawChapter, _ int) (Chapter, error) {
		nameDoc, err := goquery.NewDocumentFromReader(strings.NewReader(chapter.Name))
		if err != nil {
			return *new(Chapter), err
		}

		Name := strings.TrimSpace(nameDoc.Children().Children().Eq(1).Contents().Last().Text())

		idDoc, err := goquery.NewDocumentFromReader(strings.NewReader(chapter.Id))
		if err != nil {
			return *new(Chapter), err
		}

		Id, err := strconv.Atoi(
			strings.TrimSuffix(
				strings.TrimPrefix(idDoc.Find("div").Get(0).Attr[1].Key, "novel_on_"),
				"\"",
			),
		)
		if err != nil {
			return *new(Chapter), err
		}

		return Chapter{
			Name: Name,
			Id:   Id,
			Type: chapter.Type,
		}, nil
	})
	if err != nil {
		return *new(Toc), err
	}

	return Toc{
		Chapters: Chapters,
	}, nil
}
