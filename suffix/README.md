# Major version suffix

If a version `v2.44.88` is imported using the major version as a suffix
the tidy command will understand that the major version to be imported is v2.

```go
import "github.com/my-organization/my-package/v2"
```

```go.mod
require github.com/my-organization/my-package/v2 v2.44.88
```

> [Official documentation](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher)
