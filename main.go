package main

import (
	"log"

	"github.com/drabuna/poebuildbuyer/itemtool"
)

func main() {

	log.Println("Starting...")

	items, err := itemtool.GetBuildItemsLinks("pastebin.com/qJhJdWkP")
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
