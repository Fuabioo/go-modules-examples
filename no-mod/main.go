package main

import (
	"fmt"

	NullTime "github.com/kirillDanshin/nulltime"
)

const description = `
This program that requires a dependency that is not a go module.
If you check the go.mod you will see that 'go mod tidy' generated
a pseudo version based on the latest tag on git.

Check it out here: https://github.com/kirillDanshin/nulltime/tree/55e96302862431c0065423fd4a3341bed5b60866
`

func main() {
	fmt.Println(description)
	_ = &NullTime.NullTime{}
}
