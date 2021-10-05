# example-pkg

```go
package main

import (
    "log"

	examplepkg "github.com/Fuabioo/example-pkg"
)

func main() {
    res := examplepkg.MapKeysToCamelCase(map[string]string{
		"key_one": "value_one",
	})
    log.Println(res)
}
```
