package main

import (
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

type opml struct {
	XMLName		xml.Name	`xml:"opml"`
	Version		string		`xml:"version,attr"`
	OpmlTitle	string		`xml:"head>title"`
	Outlines	[]outline	`xml:"body>outline"`
}

type outline struct {
	Text		string		`xml:"text,attr"`
	Title		string		`xml:"title,attr"`
	Type		string		`xml:"type,attr"`
	XmlUrl		string		`xml:"xmlUrl,attr"`
	HtmlUrl		string		`xml:"htmlUrl,attr"`
	Favicon		string		`xml:"rssfr-favicon,attr"`
}

func main() {
	o := opml{}
	xmlContent, _ := ioutil.ReadFile("example-5.xml")
	err := xml.Unmarshal(xmlContent, &o)
	if err != nil { panic(err) }
	for _, outline := range o.Outlines {
		fmt.Println(outline)
	}
}
