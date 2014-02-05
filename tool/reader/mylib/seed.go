// http://www.w3schools.com/rss/default.asp
// http://www.tutorialspoint.com/rss/what-is-atom.htm
package mylib

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"html/template"
)


// http://www.w3schools.com/rss/rss_syntax.asp
type Rss2 struct {
	XMLName	xml.Name	`xml:"rss"`
	Version	string		`xml:"version,attr"`
	Channel	RssChan
}

// http://www.w3schools.com/rss/rss_channel.asp
type RssChan struct {
	XMLName		xml.Name	`xml:"channel"`
	Title		string		`xml:"title"`
	Link		string		`xml:"link"`
	Description	string		`xml:"description"`
	PubDate		string		`xml:"pubDate"`
	ItemList	[]Item		`xml:"item"`
}

// http://www.w3schools.com/rss/rss_item.asp
type Item struct {
	Title		string		`xml:"title"`
	Link		string		`xml:"link"`
	Comments	string		`xml:"comments"`
	Description	template.HTML	`xml:"description"`
	PubDate		string		`xml:"pubDate"`
	IsRead		bool
	ReadDatetime	string
}


func parseSeedContent(content []byte) Rss2 {
	v := Rss2{}
	xml.Unmarshal(content, &v)
	return v
}

func GetSeed(url string) Rss2 {
	resp, err := http.Get(url)
	if err != nil { panic(err) }
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { panic(err) }
	return parseSeedContent(body)
}
