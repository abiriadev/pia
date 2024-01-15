package api

import (
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Novel struct {
	Title    string   `goquery:".epnew-novel-title"`
	Author   string   `goquery:".in-writer .writer-name"`
	Synopsis string   `goquery:".epnew-novel-info .synopsis"`
	Tags     []string `goquery:".epnew-novel-info .writer-tag > .tag"`
}

func GetNovel(id int) (Novel, error) {
	novel, err := getHtml[Novel](endPointNovel + strconv.Itoa(id))
	if err != nil {
		return *new(Novel), err
	}

	novel.Tags = lo.Map(novel.Tags, func(tag string, _ int) string {
		return strings.TrimPrefix(tag, "#")
	})

	return novel, nil
}
