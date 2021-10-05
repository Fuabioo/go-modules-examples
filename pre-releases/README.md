# Vendoring

First, fetch the dependencies as usual:

```sh
GOPRIVATE=github.com/Fuabioo go mod tidy
```

> Remember to use GOPRIVATE when fetching dependencies with `go mod tidy` or `go get`

You may now edit the go.mod file to point at the pre-release tag.
