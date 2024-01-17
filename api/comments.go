package api

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/otterize/lox"
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

	comments, err := lox.MapErr(
		rawComments.Result,
		func(rawComment Comment, _ int) (Comment, error) {
			nameDoc, err := goquery.NewDocumentFromReader(strings.NewReader(rawComment.UserName))
			if err != nil {
				return *new(Comment), err
			}

			rawComment.UserName = nameDoc.Find("b").Text()

			return rawComment, nil
		},
	)
	if err != nil {
		return *new(Comments), err
	}

	return Comments{
		Comments: comments,
		Page:     rawComments.Page,
	}, nil
}
