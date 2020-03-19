package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/xml"
	"io/ioutil"
	"strings"
)

func ExtractPathOfBuildingData(data string) (PathOfBuilding, error) {
	xml, err := decodeAndDecompress(data)
	if err != nil {
		return PathOfBuilding{}, err
	}

	pob, err := parseXml(xml)
	if err != nil {
		return PathOfBuilding{}, err
	}
	return pob, nil
}

func decodeAndDecompress(data string) (string, error) {
	cleanBase64 := strings.ReplaceAll(strings.ReplaceAll(data, "-", "+"), "_", "/")

	compressedData, err := base64.StdEncoding.DecodeString(cleanBase64)
	if err != nil {
		return "", err
	}

	bytesReader := bytes.NewReader(compressedData)
	compressionReader, err := zlib.NewReader(bytesReader)
	if err != nil {
		return "", err
	}
	defer compressionReader.Close()

	decompressedData, err := ioutil.ReadAll(compressionReader)
	if err != nil {
		return "", err
	}
	return string(decompressedData), nil
}

func parseXml(xmlData string) (PathOfBuilding, error) {
	var pob PathOfBuilding

	err := xml.Unmarshal([]byte(xmlData), &pob)
	if err != nil {
		return PathOfBuilding{}, err
	}

	return pob, nil
}
