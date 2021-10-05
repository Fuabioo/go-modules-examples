package examplepkg

import (
	"github.com/iancoleman/strcase"
)

// MapKeysToCamelCase changes all the keys on a map to camelCase
//
// This is not even tagged yet
func MapKeysToCamelCase(v map[string]string) map[string]string {
	res := make(map[string]string)
	for key, value := range v {
		key = strcase.ToLowerCamel(key)
		res[key] = value
	}
	return res
}
