package main

import (
	"fmt"

	examplepkg "github.com/Fuabioo/example-pkg"
)

const description = `
This program uses a dependency that points to a pre-release tag rather than a released tag

GOPRIVATE=github.com/Fuabioo go mod vendor
`

func main() {
	fmt.Println(description)
	_ = examplepkg.MapKeysToCamelCase(map[string]string{
		"key_one": "value_one",
	})
}
