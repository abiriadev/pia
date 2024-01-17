package api

import (
	"net/url"
	"strconv"
)

type Comments struct {
	Lines []string
}

type Page struct {
	current  int `json:"CNT"`
	all      int `json:"ALL"`
	max      int `json:"MAX"`
	max_page int `json:"MAX_PAGE"`
	start    int `json:"START"`
	end      int `json:"END"`
	back     int `json:"BACK"`
	next     int `json:"NEXT"`
}

type rawComments struct {
	CommentCounter int  `json:"comment_counter"`
	Page           Page `json:"page"`
	Result         any  `json:"result"`
}

func GetComments(id int) (Comments, error) {
	var comments Comments
	_, err := postJson[rawComments](endPointComments, url.Values{
		"mode":       []string{"get_comment_list"},
		"episode_no": []string{strconv.Itoa(id)},
	})

	if err != nil {
		return comments, err
	}

	// for _, line := range res.Result {
	// 	comments.Lines = append(comments.Lines, strings.ReplaceAll(line.Text, "&nbsp;", " "))
	// }
	comments.Lines = []string{}

	return comments, nil
}
