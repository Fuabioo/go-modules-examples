# Replacing import versions

You can replace the version of a module everywhere it is imported.

```go
import "github.com/my-organization/my-package"
```

```go.mod
require github.com/my-organization/my-package/v2 v2.44.88
```

This way any import of `github.com/my-organization/my-package` will actually use v2
