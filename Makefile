PHONY: run

setup: pk
	GO111MODULE=off go get github.com/codegangsta/gin
	go build

pk:
	export SESSION_KEY="A�&H�Lz��2y����T<x:}���-[0�}��"
run:
	gin run router.go