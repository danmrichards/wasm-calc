GOARCH=amd64
GO111MODULE=on

build: linux darwin windows freebsd

linux: frontend server-linux

darwin: frontend server-darwin

windows: frontend server-windows

freebsd: frontend server-freebsd

frontend:
	CGO_ENABLED=0 GOARCH=wasm GOOS=js go build -o main.wasm .

server-linux:
	CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=linux go build -o ./bin/server-linux-${GOARCH} ./server

server-darwin:
	CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=darwin go build -o ./bin/server-darwin-${GOARCH} ./server

server-windows:
	CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=windows go build -o ./bin/server-windows-${GOARCH}.exe ./server

server-freebsd:
	CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=freebsd go build -o ./bin/server-freebsd-${GOARCH} ./server

.PHONY: build