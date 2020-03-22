package core

import "encoding/xml"

const ITEM_LB = "\r\n"
const ITEM_SECTIONS_SEPERATOR = "--------" + ITEM_LB

type PathOfBuilding struct {
	XMLName xml.Name `xml:"PathOfBuilding"`
	Items   struct {
		List []PobItem `xml:"Item"`
	} `xml:"Items"`
}

type PobItem struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr"`
}

type Item struct {
	Rarity     string
	Name       string
	Type       string
	Properties map[string]string
	Modifiers  []string
}

func (item Item) Export() string {
	result := ""

	//first part - "header"
	result += "Rarity: " + item.Rarity + ITEM_LB
	result += item.Name + ITEM_LB
	if len(item.Type) > 0 {
		result += item.Type + ITEM_LB
	}
	result += ITEM_SECTIONS_SEPERATOR

	//second part - "properties"
	for k, v := range item.Properties {
		result += k + ": " + v + ITEM_LB + ITEM_SECTIONS_SEPERATOR
	}

	//third part - "modifers"
	for _, m := range item.Modifiers {
		result += m + ITEM_LB
	}

	return result
}
