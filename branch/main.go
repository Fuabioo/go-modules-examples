package main

import (
	"fmt"

	examplepkg "github.com/Fuabioo/example-pkg"
)

const description = `
This program uses a dependency that points to a branch rather than a released tag

GOPRIVATE=github.com/Fuabioo go get github.com/Fuabioo/example-pkg@ebcbaf5a23c544d09b8d5c4e74bc4c8e971aa1e5
`

func main() {
	fmt.Println(description)
	_ = examplepkg.MapKeysToCamelCase(map[string]string{
		"key_one": "value_one",
	})
}
