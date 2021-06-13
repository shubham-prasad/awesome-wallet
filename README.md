# My Awesome Wallet

Because every other wallet is 💩


## Prerequisites
1. [Golang](https://golang.org/doc/install)
2. [Docker](https://www.docker.com/products/docker-desktop)
3. [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
4. [SQLC](https://docs.sqlc.dev/en/latest/overview/install.html)
5. [go2proto](https://github.com/anjmao/go2proto)
    `go install github.com/anjmao/go2proto@latest`
6. [protoc](https://grpc.io/docs/protoc-installation/)
   Install protobuf
    `brew install protobuf`
   Install plugins for go
    `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26`  
    `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1`  
   Export PATH from .bash_profile or .bashrc
    `export PATH="$PATH:$(go env GOPATH)/bin"`