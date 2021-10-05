package examplepkg

import (
	"github.com/iancoleman/strcase"
)

// MapKeysToCamelCase changes all the keys on a map to camelCase
func MapKeysToCamelCase(v map[string]string) map[string]string {
	res := make(map[string]string)
	for key, value := range v {
		key = strcase.ToLowerCamel(key)
		res[key] = value
	}
	return res
}
