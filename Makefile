bin_name = poeit
cli_bin_name = poeit

is_win = 
is_mac = 

ifeq ($(OS),Windows_NT)
    os_name := Windows
else
    os_name := $(shell uname -s)
endif

ifeq ($(os_name), Windows)
	is_win = true
    target = $(bin_name).exe
	cli_target = $(cli_bin_name).exe
endif
ifeq ($(os_name), Darwin)
	is_mac = true
    target = $(bin_name)
	cli_target = $(cli_bin_name)
endif


.PHONY: all

all: clean build run

wasm: wasm_clean wasm_build wasm_serve

clean:
ifdef is_win 
	cmd /c if exist "out\bin" cmd /c rmdir /Q /S out\bin
else
	rm -rf ./out/bin/
endif


build: 
	go build -o out/bin/$(target) 

run:
	./out/bin/$(target)

start: 
	go run .


#CLI stuff
cli_build:
	go build -o out/cli/${cli_target} ./itemtool/cli/cli.go 

cli_build_win:
ifdef is_mac 
	GOOS=windows go build -o out/cli/${cli_bin_name}.exe ./itemtool/cli/cli.go 
else 
	go build -o out/cli/${cli_bin_name}.exe ./itemtool/cli/cli.go 
endif



#WASM stuff
wasm_build:
	GOOS=js GOARCH=wasm go build -o out/wasm/$(bin_name).wasm

wasm_clean: 
	rm -f out/wasm/$(bin_name).wasm

wasm_serve:
	go run webserver/main.go

wasm_tiny:
	tinygo build -o out/wasm/$(bin_name).wasm -target wasm
