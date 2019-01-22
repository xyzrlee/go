@echo off
go get -v github.com/xyzrlee/go/logger
go build -o target/server.exe server.go
