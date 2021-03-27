

1. Build the plugin `go build -buildmode=plugin -o ./git/git.so ./git/git.go`
1. Build the cli `go build main.go`
1. Run git pull `./main git pull` which will show help from git