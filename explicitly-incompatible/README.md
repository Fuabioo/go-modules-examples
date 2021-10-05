# Explicitly marking the module as incompatible

You can import the module package as usual, but you need to specify on go.mod
that the module is incompatible

```go
import "github.com/my-organization/my-package"
```

```go.mod
require github.com/my-organization/my-package v2.44.88+incompatible
```
