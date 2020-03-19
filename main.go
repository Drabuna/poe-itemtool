package main

import (
	"log"
)

/*
name:
link_min: 3
link_max:
sockets_r: 3
sockets_g:
sockets_b: 1
sockets_w:
*/

var LEAGUES = []string{
	"Delirium",
	"Hardcore Delirium",
	"Standard",
	"Hardcore",
	"PS4 - Delirium",
	"PS4 - Hardcore Delirium",
	"PS4 - Standard",
	"PS4 - Hardcore",
	"Xbox - Delirium",
	"Xbox - Hardcore Delirium",
	"Xbox - Standard",
	"Xbox - Hardcore",
}

var MODES = []string{
	"udercut",
	"upgrade",
}

func main() {

	//"https://pastebin.com/5F4Lh16T"
	//"https://pastebin.com/JsRTzEBc"
	data, err := LoadDataFromPastebinUrl("https://pastebin.com/5F4Lh16T")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(data)

	pob, err := ExtractPathOfBuildingData(data)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Got ", len(pob.Items.List), " items")

	items, err := ParseItems(pob.Items.List)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items[0:1] {
		log.Println("\n---------------")
		log.Println("\n" + item.Export())

		url, err := FetchItemImportUrl(item, LEAGUES[0], MODES[1], 85, 125)
		if err != nil {
			log.Println("!!! ERROR !!! ", err)
		} else {
			log.Println("URL is:", url)
		}
		log.Println("\n---------------")
	}

}
