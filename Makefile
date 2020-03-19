.PHONY: env wasm

wasm:
	go build -o main.wasm .

env:
	go env -w GOARCH=wasm
	go env -w GOOS=js