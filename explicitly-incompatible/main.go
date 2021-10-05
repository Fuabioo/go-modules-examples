package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

const description = `
This program that requires a dependency that is is tagged v3+.
If you check the go.mod you will see that it is explicitly marked as +incompatible

Check it out here: https://github.com/dgrijalva/jwt-go/tree/v3.2.0
`

func main() {
	fmt.Println(description)
	_ = &jwt.StandardClaims{}
}
