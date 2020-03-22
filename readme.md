# PoE Item Tool

Given a [Path of Building](https://github.com/Openarl/PathOfBuilding) export data, give me all the links to [poe.trade](https://poe.trade/) to buy them.
Currently only availiable as a CLI tool.

# How to use

Download latest release binary from [releases](https://github.com/Drabuna/poe-itemtool/releases) 
Run it from terminal on Mac:
`./poeit get pastebin.com/qJhJdWkP`

or from command line on Windows:
`poeit.exe get pastebin.com/qJhJdWkP`

# Building 
Or clone the repo and then:
`go get`
 - On Mac: 
	 - `make cli_build_mac`
	 - `./out/cli/poeit get pastebin.com/qJhJdWkP`
 - On Windows:
	 - `make cli_build_win`
	 - `poeit.exe get pastebin.com/qJhJdWkP`

# Flags
 - get -data base64_export
	 - `poeit get -data NrtPWtz2za2n6tfwfFM7yTj2CbAF...` 

# Roadmap

 - [ ] Some tests
 - [ ] UI
 - [ ] Items preview and price analysis
 - [ ] Autoupdater
 - [ ] Proper modifier parser and more acqurate seach
 - [ ] Support multiple stores