package mylib

import (
	"io/ioutil"
	"encoding/xml"
)


type Opml struct {
	XMLName	xml.Name	`xml:"opml"`
	Version	string		`xml:"version,attr"`
	Head	OpmlHead
	Body	OpmlBody
}

type OpmlHead struct {
	XMLName xml.Name	`xml:"head"`
	Title	string		`xml:"title"`
}

type OpmlBody struct {
	XMLName		xml.Name	`xml:"body"`
	SiteList	[]OpmlOutline	`xml:"outline"`
}

type OpmlOutline struct {
	Text	string		`xml:"text,attr"`
	Title	string		`xml:"title,attr"`
	Type	string		`xml:"type,attr"`
	XmlUrl	string		`xml:"xmlUrl,attr"`
	HtmlUrl	string		`xml:"htmlUrl,attr"`
	Favicon	string		`xml:"rssfr-favicon,attr"`
}


func GetOutlineList(filepath string) []OpmlOutline {
	v := Opml{}
	content, _ := ioutil.ReadFile(filepath)
	err := xml.Unmarshal(content, &v)
	if err != nil { panic(err) }
	return v.Body.SiteList
}
