package core

import (
	"errors"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"
)

func FetchItemImportUrl(item Item, league string, mode string, undercut int, uppercut int) (string, error) {
	poeTradeUrl := "https://poe.trade/import"
	method := "POST"

	payloadData := ""
	payloadData += "data=" + escapeData(item.Export()) + "&"
	payloadData += "league=" + escapeData(league) + "&"
	payloadData += "mode=" + escapeData(mode) + "&"
	payloadData += "undercut=" + strconv.Itoa(undercut) + "&"
	payloadData += "uppercut=" + strconv.Itoa(uppercut) + "&"
	payloadData += "name=" + escapeData(item.Name)

	payload := strings.NewReader(payloadData)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	req, err := http.NewRequest(method, poeTradeUrl, payload)

	if err != nil {
		return "", err
	}

	if runtime.GOARCH == "wasm" {
		req.Header.Add("js.fetch:mode", "cors")
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode == 302 {
		return res.Header.Get("Location"), nil
	}
	return "", errors.New("Failed to fetch search url; HTTP status code: " + strconv.Itoa(res.StatusCode))
}

func escapeData(d string) string {
	res := url.QueryEscape(d)
	return res
}
