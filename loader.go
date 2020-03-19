package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func LoadDataFromPastebinUrl(url string) (string, error) {
	rawUrl := pastebinNormalToRaw(url)
	data, err := loadBuildData(rawUrl)
	if err != nil {
		return "", err
	}
	return data, nil
}

func pastebinNormalToRaw(url string) string {
	rawUrl := strings.Replace(url, "pastebin.com/", "pastebin.com/raw/", 1)
	return rawUrl
}

func loadBuildData(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("js.fetch:mode", "no-cors")

	if err != nil {
		return "", err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", err
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	body := string(bodyBytes)
	return body, nil
}
