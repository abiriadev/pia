package main

import (
	"github.com/abiriadev/pia/api"
	"github.com/kr/pretty"
)

func main() {
	res, err := api.GetToc(145563, 1)
	if err != nil {
		panic(err)
	}

	pretty.Println(res)
}
