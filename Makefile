.PHONY: all

all: clean test build run

dev: test-verbose start

wasm: wasm_clean test wasm_build wasm_serve

clean:
	rm -rf /out/bin/

build: 
	go build -o out/bin/poeit 

run:
	./out/bin/poeit

start: 
	go run .

test:
	go test ./tests

test-verbose:
	go test ./tests -v


#CLI stuff
cli_build_mac:
	go build -o out/cli/poeit ./itemtool/cli/cli.go 

cli_build_win:
	GOOS=windows go build -o out/cli/poeit.exe ./itemtool/cli/cli.go 


#WASM stuff
wasm_build:
	GOOS=js GOARCH=wasm go build -o out/wasm/itemtool.wasm

wasm_clean: 
	rm -f out/wasm/itemtool.wasm

wasm_serve:
	go run webserver/main.go

wasm_tiny:
	tinygo build -o out/wasm/itemtool.wasm -target wasm
