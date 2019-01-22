all: init build

init:
	go get -v github.com/xyzrlee/go/logger

build:
	go build -o target/server ./server.go

clean:
	rm -rf ./target
