package main

import (
	"fmt"

	examplepkg "github.com/Fuabioo/example-pkg"
)

const description = `
This program uses a private repository, so in order to fetch the package you need
to use the GOPRIVATE environment variable:

GOPRIVATE=github.com/Fuabioo go mod tidy
`

func main() {
	fmt.Println(description)
	_ = examplepkg.MapKeysToCamelCase(map[string]string{
		"key_one": "value_one",
	})
}
