package main

import (
	"fmt"

	examplepkg "github.com/Fuabioo/example-pkg"
)

const description = `
This program stores a copy of the dependency on the vendor folder

GOPRIVATE=github.com/Fuabioo go mod vendor
`

func main() {
	fmt.Println(description)
	_ = examplepkg.MapKeysToCamelCase(map[string]string{
		"key_one": "value_one",
	})
}
