package api

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AuthorComment struct {
	Comment string `goquery:"#footer_plus"`
}

func GetAuthorComment(id int) (AuthorComment, error) {
	authorComment, err := getHtml[AuthorComment](endPointChapter + strconv.Itoa(id))
	if err != nil {
		return *new(AuthorComment), err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(authorComment.Comment))
	if err != nil {
		return *new(AuthorComment), err
	}

	authorComment.Comment = doc.Find("#writer_comments_box").
		Contents().Slice(2, goquery.ToEnd).Text()

	return authorComment, nil
}
