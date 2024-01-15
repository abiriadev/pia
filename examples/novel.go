package main

import (
	"github.com/abiriadev/pia/api"
	"github.com/kr/pretty"
)

func main() {
	res, err := api.GetNovel(145563)
	if err != nil {
		panic(err)
	}

	pretty.Println(res)
}
