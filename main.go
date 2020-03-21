package main

import (
	"log"

	"github.com/drabuna/poebuildbuyer/itemtool"
)

func main() {

	log.Println("Starting...")

	//https://pastebin.com/qJhJdWkP - article build
	//https://pastebin.com/Xx3EAYWC - witch build

	items, err := itemtool.GetBuildItemsLinks("https://pastebin.com/Xx3EAYWC")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		if item.Error != nil {
			log.Println(item.Name, " - ", "Error: ", item.Error)
		} else {
			log.Println(item.Name, " - ", item.URL)
		}
	}
}
