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

`make cli_build`

- On Mac:

`./out/cli/poeit get pastebin.com/qJhJdWkP`

- On Windows:

`poeit.exe get pastebin.com/qJhJdWkP`

# Flags

- Specify base64 data instead of url using `-data` flag:

`poeit get -data NrtPWtz2za2n6tfwfFM7yTj2CbAF...`

- Specify a league to search in using `-league` flag:

`poeit get -league "PS4 - Delirium" pastebin.com/qJhJdWkP`

- Specify a search mode using `-mode` flag:

`poeit get -league "Standard" -mode "upgrade" pastebin.com/qJhJdWkP`

or

`poeit get -league "Standard" -mode "undercut" -low 85 -high 125 pastebin.com/qJhJdWkP`

# Roadmap

- [ ] Some tests
- [ ] UI
- [ ] Items preview and price analysis
- [ ] Autoupdater
- [ ] Proper modifier parser and more acqurate seach
- [ ] Support multiple stores
