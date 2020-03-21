package core

import (
	"errors"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
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
	if err != nil {
		return "", err
	}
	if runtime.GOARCH == "wasm" {
		req.Header.Add("js.fetch:mode", "cors")
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", errors.New("Bad status code: " + strconv.Itoa(response.StatusCode))
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	body := string(bodyBytes)
	return body, nil
}
