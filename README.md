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

## Structure

```
├── clock
├── cryptography
│   ├── checksum
│   └── random
├── freax
├── inout
│   ├── colorz
│   ├── dntfy
│   ├── logz
│   └── pinentry
├── net
│   ├── arp
│   ├── blocklist
│   ├── capability
│   ├── ctlog
│   ├── dns
│   │   └── hetzner
│   ├── iface
│   ├── ip
│   ├── portscan
│   ├── raw
│   ├── route
│   ├── tcp
│   ├── tls
│   │   ├── certificate
│   │   ├── ciphersuite
│   │   │   └── tools
│   │   ├── ssl30
│   │   ├── tls10
│   │   ├── tls11
│   │   ├── tls12
│   │   └── tls13
│   └── validator
├── octets
└── slicer
```