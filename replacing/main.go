package main

import (
	"fmt"

	"github.com/go-resty/resty"
)

const description = `
This program is a case in which a pre-modules package had to be upgraded to go modules
but can not be upgraded completely, so a replace was used to handle it.

Check it out here: https://github.com/go-resty/resty/blob/v2.6.0
`

func main() {
	fmt.Println(description)
	_ = resty.New()
}
