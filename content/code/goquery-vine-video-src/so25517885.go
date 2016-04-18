package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
)

type SharedContent struct {
	ContentUrl string `json:"contentUrl"`
}

type VineVideoMetadata struct {
	SC SharedContent `json:"sharedContent"`
}

func DecodeVineJsonBlob(blob string) VineVideoMetadata {
	meta := VineVideoMetadata{}
	err := json.Unmarshal([]byte(blob), &meta)
	if err != nil {
		panic(err)
	}
	return meta
}

func GetVineVideoJsonBlob(url string) string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	return doc.Find("script[type=\"application/ld+json\"]").Text()
}

func GetVineVideoSrc(url string) string {
	jsonBlob := GetVineVideoJsonBlob(url)
	meta := DecodeVineJsonBlob(jsonBlob)
	return meta.SC.ContentUrl
}

func main() {
	println(GetVineVideoSrc("https://vine.co/v/MlWtKgwh7WY"))
}
