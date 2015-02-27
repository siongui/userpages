// http://www.w3schools.com/rss/default.asp
// http://www.tutorialspoint.com/rss/what-is-atom.htm
// http://stackoverflow.com/questions/6619717/what-is-the-difference-between-rss-and-atom-feeds
// http://stackoverflow.com/questions/16309944/atom-to-rss-feeds-converter
package main

import (
	"io/ioutil"
	"encoding/xml"
	"html/template"
	"fmt"
)


// http://www.w3schools.com/rss/rss_syntax.asp
// http://www.w3schools.com/rss/rss_channel.asp
type Rss2 struct {
	XMLName		xml.Name	`xml:"rss"`
	Version		string		`xml:"version,attr"`
	// Required
	Title		string		`xml:"channel>title"`
	Link		string		`xml:"channel>link"`
	Description	string		`xml:"channel>description"`
	// Optional
	PubDate		string		`xml:"channel>pubDate"`
	ItemList	[]Item		`xml:"channel>item"`
}

// http://www.w3schools.com/rss/rss_item.asp
// http://stackoverflow.com/questions/7220670/difference-between-description-and-contentencoded-tags-in-rss2
// https://groups.google.com/d/topic/golang-nuts/uBMo1BpaQCM
type Item struct {
	// Required
	Title		string		`xml:"title"`
	Link		string		`xml:"link"`
	Description	template.HTML	`xml:"description"`
	// Optional
	Content		template.HTML	`xml:"encoded"`
	PubDate		string		`xml:"pubDate"`
	Comments	string		`xml:"comments"`
}


func main() {
	r := Rss2{}
	xmlContent, _ := ioutil.ReadFile("example-6.xml")
	err := xml.Unmarshal(xmlContent, &r)
	if err != nil { panic(err) }
	for _, item := range r.ItemList {
		fmt.Println(item)
	}
}
