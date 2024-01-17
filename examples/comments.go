package main

import (
	"github.com/abiriadev/pia/api"
	"github.com/kr/pretty"
)

func main() {
	res, err := api.GetComments(1720152)
	if err != nil {
		panic(err)
	}

	pretty.Println(res)
}
