package main

import (
	"fmt"

	"github.com/spf13/viper"
)

const description = `
This program that requires a dependency with indirect dependencies.
If you compare the go.mod file with viper's go.mod you may notice that
all the indirects of this program are direct requirements for viper.

Check it out here: https://github.com/spf13/viper/blob/v1.9.0/go.mod
`

func main() {
	fmt.Println(description)
	viper.AutomaticEnv()
}
