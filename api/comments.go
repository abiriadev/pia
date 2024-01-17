package api

import (
	"fmt"
	"net/url"
	"strconv"
)

type Comments struct {
	Lines []string
}

type Page struct {
	Current int `json:"CNT"`
	All     int `json:"ALL"`
	Max     int `json:"MAX"`
	MaxPage int `json:"MAX_PAGE"`
	Start   int `json:"START"`
	End     int `json:"END"`
	Back    int `json:"BACK"`
	Next    int `json:"NEXT"`
}

type rawComments struct {
	CommentCounter int  `json:"comment_counter"`
	Page           Page `json:"page"`
	Result         any  `json:"result"`
}

func GetComments(id int) (Comments, error) {
	var comments Comments
	rawComments, err := postJson[rawComments](endPointComments, url.Values{
		"mode":       []string{"get_comment_list"},
		"episode_no": []string{strconv.Itoa(id)},
	})

	if err != nil {
		return comments, err
	}

	fmt.Println(rawComments.Page)

	// for _, line := range res.Result {
	// 	comments.Lines = append(comments.Lines, strings.ReplaceAll(line.Text, "&nbsp;", " "))
	// }
	comments.Lines = []string{}

	return comments, nil
}
