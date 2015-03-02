package main

import (
	"io/ioutil"
	"encoding/xml"
	"html/template"
	"fmt"
)


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


type Atom1 struct {
	XMLName		xml.Name	`xml:"http://www.w3.org/2005/Atom feed"`
	Title		string		`xml:"title"`
	Subtitle	string		`xml:"subtitle"`
	Id		string		`xml:"id"`
	Updated		string		`xml:"updated"`
	Rights		string		`xml:"rights"`
	Link		Link		`xml:"link"`
	Author		Author		`xml:"author"`
	EntryList	[]Entry		`xml:"entry"`
}

type Link struct {
	Href		string		`xml:"href,attr"`
}

type Author struct {
	Name		string		`xml:"name"`
	Email		string		`xml:"email"`
}

type Entry struct {
	Title		string		`xml:"title"`
	Summary		string		`xml:"summary"`
	Content		string		`xml:"content"`
	Id		string		`xml:"id"`
	Updated		string		`xml:"updated"`
	Link		Link		`xml:"link"`
	Author		Author		`xml:"author"`
}

func atom1ToRss2(a Atom1) Rss2 {
	r := Rss2{
		Title: a.Title,
		Link: a.Link.Href,
		Description: a.Subtitle,
		PubDate: a.Updated,
	}
	r.ItemList = make([]Item, len(a.EntryList))
	for i, entry := range a.EntryList {
		r.ItemList[i].Title = entry.Title
		r.ItemList[i].Link = entry.Link.Href
		if entry.Content == "" {
			r.ItemList[i].Description = template.HTML(entry.Summary)
		} else {
			r.ItemList[i].Description = template.HTML(entry.Content)
		}
	}
	return r
}


func main() {
	a := Atom1{}
	xmlContent, _ := ioutil.ReadFile("example-7.xml")
	err := xml.Unmarshal(xmlContent, &a)
	if err != nil { panic(err) }
	r := atom1ToRss2(a)
	fmt.Println(r)
}
