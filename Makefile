.PHONY: air build server

air:
	air -c .air.toml

build:
	go build -o bin/server github.com/erichaase/fantasy/cmd/server

server: build
	bin/server -localhost -port 3001