package core

import (
	"errors"
	"log"
	"regexp"
	"strings"
)

func ParseItems(pobItems []PobItem) ([]Item, error) {
	items := []Item{}
	for index, pobItem := range pobItems {
		item, err := ParseItem(pobItem)
		if err != nil {
			log.Print("Couldn't parse item at index ", index, " due to error: ", err)
		} else {
			items = append(items, item)
		}
	}
	if len(items) == 0 {
		return nil, errors.New("No items were parsed")
	}
	return items, nil
}

/*
Expected format:
First line alway contains Rarity: type
Second line contains the name of the item
Third line contains type of the item
Every next line that has a type: value is something that can be parsed
Every line that doesn't have a "type:" is a modifier
*/

func ParseItem(pobItem PobItem) (Item, error) {
	item := Item{}

	hasType := false

	lines := strings.Split(pobItem.Text, "\n")

	//remove empty lines and trim spaces, and other special modifies and meta info like {stuff}
	cbracketsRegexp := regexp.MustCompile(`\{(.*?)\}`)
	cleanLines := []string{}
	for _, line := range lines {
		tmpLine := strings.TrimSpace(line)
		if len(tmpLine) > 0 {
			tmpLine = cbracketsRegexp.ReplaceAllString(tmpLine, "")
			cleanLines = append(cleanLines, tmpLine)
		}
	}

	//validate a min required lines to be an "item"
	if len(cleanLines) < 2 {
		return Item{}, errors.New("Too few lines, unsupported item format")
	}

	//parse rarity
	if !isItemProperty(cleanLines[0]) {
		return Item{}, errors.New("Unable to extract rarity, unsupported item format")
	}

	rarityKey, rarityValue, extracted := parseItemPropety(cleanLines[0])
	if !extracted || rarityKey != "Rarity" || len(rarityValue) == 0 {
		return Item{}, errors.New("Unable to extract rarity, unsupported item format")
	}
	item.Rarity = strings.Title(strings.ToLower(rarityValue))

	//set name
	item.Name = cleanLines[1]

	//type might not be provided, need to validate
	if len(cleanLines) > 2 {
		if !isItemProperty(cleanLines[2]) {
			hasType = true
			item.Type = cleanLines[2]
		}
	}

	//parse properties and modifiers
	properties := map[string]string{}
	modifiers := []string{}

	startFromIndex := 2
	if hasType {
		startFromIndex = 3
	}

	for index, line := range cleanLines[startFromIndex:] {
		if isItemProperty(line) {
			key, value, success := parseItemPropety(line)
			if !success || len(key) == 0 || len(value) == 0 {
				log.Println("Couldn't extract property from line:", line, "at index: ", index)
			} else {
				properties[key] = value
			}
		} else {
			modifiers = append(modifiers, line)
		}
	}

	item.Properties = properties
	item.Modifiers = modifiers

	return item, nil
}

func isItemProperty(line string) bool {
	return strings.Contains(line, ":")
}

//returns key, value, false if doesnt work
func parseItemPropety(property string) (string, string, bool) {
	keyValue := strings.Split(property, ":")
	if len(keyValue) != 2 {
		return "", "", false
	}
	return strings.TrimSpace(keyValue[0]), strings.TrimSpace(keyValue[1]), true
}
