package main

import (
	"log"
	"os"

	"github.com/drabuna/poebuildbuyer/itemtool"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "Path of Exile Item Tool",
		Commands: []*cli.Command{
			{
				Name:      "get",
				Usage:     "Gets links to the items from PoB pastebin URL or PoB exported data.",
				UsageText: "poeit get [URL] [arguments...]  ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "data",
						Usage: "Exported build data from Path of Building",
					},
				},
				Action: func(c *cli.Context) error {
					if c.NArg() == 0 && c.NumFlags() == 0 {
						log.Fatalln("Plese, provide a valid URL or data")
					}

					//load from base64 input instead
					if c.String("data") != "" {
						items, err := itemtool.GetBuildItemsLinksFromData(c.String("data"))
						if err != nil {
							log.Fatalln(err)
						}
						printItems(items)
					} else {
						link := c.Args().First()
						items, err := itemtool.GetBuildItemsLinks(link)
						if err != nil {
							log.Fatalln(err)
						}
						printItems(items)
					}
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

func printItems(items []itemtool.ItemInfo) {
	for _, item := range items {
		if item.Error != nil {
			log.Println(item.Name, " - ", "Error: ", item.Error)
		} else {
			log.Println(item.Name, " - ", item.URL)
		}
	}
}
