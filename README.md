# Go Modules Examples
<img alt="Go version" src="https://img.shields.io/badge/go-1.17-blue">
<img alt="Init branch" src="https://img.shields.io/badge/branch-mod%20init%20only-red">

This repository contains multiple examples of go modules scenarios:

- [Indirect dependencies](indirect/go.mod)
- [Explicit incompatibility](explicitly-incompatible/go.mod)
- [The dependency is not a module](no-mod/go.mod)
- [Replacing a dependency version](replacing/go.mod)
- [Semver with a 2+ major](suffix/go.mod)
- [Private repositories](private/go.mod)
- [Vendor dependencies](vendor/go.mod)
- [Pre-releasing a tag](pre-release/go.mod)
- [Fetching a specific branch](branch/go.mod)
- [Fetching a specific commit](commit/go.mod)

## Useful information

- [How to release a v2+ module](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher)
- [How to point to latest commit?](https://stackoverflow.com/questions/53682247/how-to-point-go-module-dependency-in-go-mod-to-a-latest-commit-in-a-repo/53682399)
- [What is _incompatible and will it cause harm?](https://stackoverflow.com/questions/57355929/what-does-incompatible-in-go-mod-mean-will-it-cause-harm)

## Relevant changes by version

- [Go 1.13](https://golang.org/doc/go1.13#modules)
- [Go 1.16](https://golang.org/doc/go1.16#go-command)
- [Go 1.17](https://golang.org/doc/go1.17#go-command)
