.PHONY: all

all: clean build run


clean:
	rm -rf /out/bin/

build: 
	go build -o out/bin/itemtool 

run:
	./out/bin/itemtool

start: 
	go run .


wasm: wasm_clean wasm_build wasm_serve

wasm_build:
	GOOS=js GOARCH=wasm go build -o out/wasm/itemtool.wasm

wasm_clean: 
	rm -f out/wasm/itemtool.wasm

wasm_serve:
	go run webserver/main.go

wasm_tiny:
	tinygo build -o out/wasm/itemtool.wasm -target wasm
