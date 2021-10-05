# Private packages

Asuming you have a private package you want to import into your project then you
need to use the GOPRIVATE environment variable to mark a namespace as private,
this will let go mod know that it should use the git authorization that you have
configured on your machine.

```sh
GOPRIVATE=github.com/Fuabioo go mod tidy
```
