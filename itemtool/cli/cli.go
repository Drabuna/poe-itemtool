package main

import (
	"log"
	"os"

	"github.com/drabuna/poebuildbuyer/itemtool"
	"github.com/urfave/cli/v2"
)

func main() {
	data := ""
	searchConf := itemtool.SearchConfig{}

	app := &cli.App{
		Name:     "Path of Exile Item Tool",
		HideHelp: false,

		Commands: []*cli.Command{
			{
				Name:      "get",
				Usage:     "Gets links to the items from PoB pastebin URL or PoB exported data.",
				UsageText: "poeit get [arguments...] [URL]  ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "data",
						Usage:       "Exported build data from Path of Building",
						Destination: &data,
					},
					&cli.StringFlag{
						Name:        "league",
						Value:       "Delirium",
						Usage:       "League to search in. Supported values: 'Delirium','Hardcore Delirium','Standard','Hardcore','PS4 - Delirium','PS4 - Hardcore Delirium','PS4 - Standard','PS4 - Hardcore','Xbox - Delirium','Xbox - Hardcore Delirium','Xbox - Standard','Xbox - Hardcore'",
						Destination: &searchConf.League,
					},
					&cli.StringFlag{
						Name:        "mode",
						Value:       "upgrade",
						Usage:       "Search mode to use. Supported values: 'upgrade', 'undercut'. If 'undercut' is used, can specify low and high parameters",
						Destination: &searchConf.Mode,
					},
					&cli.IntFlag{
						Name:        "low",
						Value:       85,
						Usage:       "Controls the lower bound of undercut search.",
						Destination: &searchConf.Undercut,
					},
					&cli.IntFlag{
						Name:        "high",
						Value:       125,
						Usage:       "Controls the upper bound of undercut search.",
						Destination: &searchConf.Uppercut,
					},
				},
				Action: func(c *cli.Context) error {
					if c.NArg() == 0 && c.NumFlags() == 0 {
						log.Fatalln("Plese, provide a valid URL or data")
					}

					//load from base64 input instead
					if data != "" {
						items, err := itemtool.GetBuildItemsLinksFromData(data, searchConf)
						if err != nil {
							log.Fatalln(err)
						}
						printItems(items)
					} else {
						link := c.Args().First()
						items, err := itemtool.GetBuildItemsLinks(link, searchConf)
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
