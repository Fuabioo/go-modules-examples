package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

const description = `
This program that requires a dependency that is is tagged v2+.
If you check the go.mod you will see that because we imported using the /v2 suffix
then the latest v2 version tag was selected.

Check it out here: https://github.com/go-resty/resty/blob/v2.6.0/go.mod
`

func main() {
	fmt.Println(description)
	_ = resty.New()
}
