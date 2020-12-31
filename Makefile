.PHONY: air build_all build_importer build_server importer server

air:
	air -c .air.toml

build_all: build_importer build_server

build_importer:
	go build -v -o tmp/importer github.com/erichaase/fantasy/cmd/importer

build_server:
	go build -v -o tmp/server github.com/erichaase/fantasy/cmd/server

importer: build_importer
	tmp/importer

server: build_server
	tmp/server -localhost