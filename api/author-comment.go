package api

import (
	"strconv"
)

type AuthorComment struct {
	Comment string `goquery:"#footer_plus"`
}

func GetAuthorComment(id int) (AuthorComment, error) {
	authorComment, err := getHtml[AuthorComment](endPointChapter + strconv.Itoa(id))
	if err != nil {
		return *new(AuthorComment), err
	}

	return authorComment, nil
}
