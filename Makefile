.PHONY: air build server

air:
	air -c .air.toml

build:
	go build -o tmp/server github.com/erichaase/fantasy/cmd/server

server: build
	tmp/server -localhost -port 3001