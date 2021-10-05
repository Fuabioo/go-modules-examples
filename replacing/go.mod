module github.com/Fuabioo/go-modules-examples/+incompatible/replacing

go 1.17

replace github.com/go-resty/resty => github.com/go-resty/resty/v2 v2.6.0

require (
	github.com/go-resty/resty v1.12.0 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
)
