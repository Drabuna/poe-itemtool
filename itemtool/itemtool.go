package itemtool

import (
	"github.com/drabuna/poebuildbuyer/itemtool/core"
)

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

type ItemInfo struct {
	Name  string
	Error error
	URL   string
	Item  core.Item
}

func GetBuildItemsLinks(pastebinURL string) ([]ItemInfo, error) {
	data, err := core.LoadDataFromPastebinUrl(pastebinURL)
	if err != nil {
		return nil, err
	}
	return GetBuildItemsLinksFromData(data)
}

func GetBuildItemsLinksFromData(data string) ([]ItemInfo, error) {
	pob, err := core.ExtractPathOfBuildingData(data)
	if err != nil {
		return nil, err
	}
	items, err := core.ParseItems(pob.Items.List)
	if err != nil {
		return nil, err
	}

	results := []ItemInfo{}
	for _, item := range items {
		url, err := core.FetchItemImportUrl(item, LEAGUES[0], MODES[1], 85, 125)

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
