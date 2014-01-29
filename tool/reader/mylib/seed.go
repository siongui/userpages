package mylib

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"html/template"
)


type Rss2 struct {
	XMLName	xml.Name	`xml:"rss"`
	Version	string		`xml:"version,attr"`
	Channel	RssChan
}

type RssChan struct {
	XMLName	xml.Name	`xml:"channel"`
	Title	string		`xml:"title"`
	Link	string		`xml:"link"`
	Description	string	`xml:"description"`
	PubDate		string	`xml:"pubDate"`
	ItemList	[]Item	`xml:"item"`
}

type Item struct {
	Title	string		`xml:"title"`
	Link	string		`xml:"link"`
	Comments	string	`xml:"comments"`
	Description	template.HTML	`xml:"description"`
	PubDate		string	`xml:"pubDate"`
}

func GetSeed(url string) Rss2 {
	resp, err := http.Get(url)
	if err != nil { panic(err) }
	defer resp.Body.Close()

	v := Rss2{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { panic(err) }
	xml.Unmarshal(body, &v)
	return v
}
