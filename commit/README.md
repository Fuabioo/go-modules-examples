# Vendoring

Specify the commit id using go get, this will add it to the go.mod:

```sh
GOPRIVATE=github.com/Fuabioo go get github.com/Fuabioo/example-pkg@ebcbaf5a23c544d09b8d5c4e74bc4c8e971aa1e5
```

> Remember to use GOPRIVATE when fetching dependencies with `go mod tidy` or `go get`
