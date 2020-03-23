package itemtool

import (
	"log"

	"github.com/drabuna/poebuildbuyer/itemtool/core"
)

type ItemInfo struct {
	Name  string
	Error error
	URL   string
	Item  core.Item
}

type SearchConfig struct {
	League   string
	Mode     string
	Undercut int
	Uppercut int
}

func GetBuildItemsLinks(pastebinURL string, config SearchConfig) ([]ItemInfo, error) {
	data, err := core.LoadDataFromPastebinUrl(pastebinURL)
	if err != nil {
		log.Println("Failed to load Pastebin data")
		return nil, err
	}
	return GetBuildItemsLinksFromData(data, config)
}

func GetBuildItemsLinksFromData(data string, config SearchConfig) ([]ItemInfo, error) {
	pob, err := core.ExtractPathOfBuildingData(data)
	if err != nil {
		log.Println("Failed to import PoB data")
		return nil, err
	}
	items, err := core.ParseItems(pob.Items.List)
	if err != nil {
		log.Println("Failed to parse items")
		return nil, err
	}

	results := []ItemInfo{}
	for _, item := range items {
		url, err := core.FetchItemImportUrl(item, config.League, config.Mode, config.Undercut, config.Uppercut)

		info := ItemInfo{Name: item.Name, Item: item}

		if err != nil {
			info.Error = err
		} else {
			info.URL = url
		}
		results = append(results, info)
	}
	return results, nil
}
