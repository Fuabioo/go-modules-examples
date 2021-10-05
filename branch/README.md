# Vendoring

Specify the branch using go get, this will add it to the go.mod:

```sh
GOPRIVATE=github.com/Fuabioo go get github.com/Fuabioo/example-pkg@a-branch
```

> Remember to use GOPRIVATE when fetching dependencies with `go mod tidy` or `go get`
