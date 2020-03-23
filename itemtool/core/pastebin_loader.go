package core

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
)

func LoadDataFromPastebinUrl(url string) (string, error) {
	rawUrl, err := validateAndTransformUrl(url)
	if err != nil {
		return "", err
	}
	data, err := loadBuildData(rawUrl)
	if err != nil {
		return "", err
	}
	return data, nil
}

func validateAndTransformUrl(link string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	if u.Scheme != "https" {
		u.Scheme = "https"
		return validateAndTransformUrl(u.String())
	}

	if u.Host != "pastebin.com" {
		return "", errors.New("Provided unsupported hostname " + u.Host + ", please use pastebin.com")
	}

	if len(u.Path) == 0 || u.Path == "/" {
		return "", errors.New("Provided URL is invalid")
	}

	u.Path = "/raw" + u.Path

	return u.String(), nil
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
		return "", errors.New("HTTP code: " + strconv.Itoa(response.StatusCode))
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	body := string(bodyBytes)
	return body, nil
}
