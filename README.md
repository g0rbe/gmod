# gmod

[![Go Reference](https://pkg.go.dev/badge/github.com/g0rbe/gmod.svg)](https://pkg.go.dev/github.com/g0rbe/gmod)
[![Go Report Card](https://goreportcard.com/badge/github.com/g0rbe/gmod)](https://goreportcard.com/report/github.com/g0rbe/gmod)

Package `gmod` is a collection of Go packages for many things.

Tried to give unique names to packages to avoid interference with the `std` package but still can infer the usage.

Get:
```bash
go get github.com/g0rbe/gmod@latest
```

Or get the latest commit (if Go module proxy is not updated):
```bash
go get "github.com/g0rbe/gmod@$(curl -s 'https://api.github.com/repos/g0rbe/gmod/commits' | jq -r '.[0].sha')"
```