package main

import (
	"fmt"

	"github.com/abiriadev/pia/api"
)

func main() {
	res, err := api.GetScript(1720152)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Textify())
}
