// http://www.tutorialspoint.com/rss/what-is-atom.htm
// http://stackoverflow.com/questions/16309944/atom-to-rss-feeds-converter
package main

import (
	"io/ioutil"
	"encoding/xml"
	"fmt"
)


// http://en.wikipedia.org/wiki/Atom_%28standard%29
// http://golang.org/src/pkg/encoding/xml/
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

func main() {
	a := Atom1{}
	xmlContent, _ := ioutil.ReadFile("example-7.xml")
	err := xml.Unmarshal(xmlContent, &a)
	if err != nil { panic(err) }
	for _, entry := range a.EntryList {
		fmt.Println(entry)
	}
}
