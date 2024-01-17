package api

import (
	"fmt"
	"net/url"
	"strconv"
)

type Comments struct {
	Comments []Comment
	Page     Page
}

type Comment struct {
	Id            int    `json:"comment_idx"`
	UserName      string `json:"user_name"`
	RegDate       string `json:"comment_regdate"`
	Message       string `json:"comment_msg"`
	ChildrenCount string `json:"comment_child_cnt"`
	UserImage     string `json:"user_img"`
	Image         string `json:"comment_img"`
	Upvotes       int    `json:"comment_vote"`
	Downvotes     int    `json:"comment_bad"`
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
	CommentCounter int       `json:"comment_counter"`
	Page           Page      `json:"page"`
	Result         []Comment `json:"result"`
}

func GetComments(id int) (Comments, error) {
	rawComments, err := postJson[rawComments](endPointComments, url.Values{
		"mode":       []string{"get_comment_list"},
		"episode_no": []string{strconv.Itoa(id)},
	})

	if err != nil {
		return *new(Comments), err
	}

	fmt.Println(rawComments.Page)

	return Comments{
		Comments: rawComments.Result,
		Page:     rawComments.Page,
	}, nil
}
