package main

import (
	"fmt"

	"github.com/abiriadev/pia"
)

func main() {
	client := pia.NewPiaClient("")

	content, err := client.Novel(1720152).Content()
	if err != nil {
		panic(err)
	}

	fmt.Println(content)
}
