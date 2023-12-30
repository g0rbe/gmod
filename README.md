# go-gmod

Get:
```bash
go get github.com/g0rbe/gmod@latest
```

Or get the latest commit (if Go module proxy is not updated)"
```bash
go get "github.com/g0rbe/gmod@$(curl -s 'https://api.github.com/repos/g0rbe/gmod/commits' | jq -r '.[0].sha')"
```