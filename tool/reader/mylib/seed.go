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
	// Required
	Title		string		`xml:"title"`
	Link		string		`xml:"link"`
	Description	string		`xml:"description"`
	// Optional
	PubDate		string		`xml:"pubDate"`
	ItemList	[]Item		`xml:"item"`
}

// http://www.w3schools.com/rss/rss_item.asp
type Item struct {
	// Required
	Title		string		`xml:"title"`
	Link		string		`xml:"link"`
	Description	template.HTML	`xml:"description"`
	// Optional
	PubDate		string		`xml:"pubDate"`
	Comments	string		`xml:"comments"`
	// for reader only
	IsRead		bool
	ReadDatetime	string
}


func parseSeedContent(content []byte) Rss2 {
	v := Rss2{}
	xml.Unmarshal(content, &v)
	return v
}

func GetSeed(url string) (Rss2, bool) {
	resp, err := http.Get(url)
	if err != nil { return Rss2{}, false }
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { return Rss2{}, false }
	return parseSeedContent(body), true
}
