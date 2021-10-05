# Vendoring

First, fetch the dependencies as usual:

```sh
GOPRIVATE=github.com/Fuabioo go mod tidy
```

> Remember to use GOPRIVATE when fetching dependencies with `go mod tidy` or `go get`

Now that the dependencies are in your system, you may run the vendor (no GOPRIVATE required)

```sh
go mod vendor
```

> Be careful vendoring private repositories on public ones!
