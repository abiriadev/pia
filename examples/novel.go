package main

import (
	"fmt"

	"github.com/abiriadev/pia/api"
)

func main() {
	res, err := api.GetNovel(145563)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
