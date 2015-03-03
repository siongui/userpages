package main

import (
	"io/ioutil"
	"encoding/xml"
	"html/template"
	"log"
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


const atomErrStr = "expected element type <rss> but have <feed>"

func parseAtom(content []byte) (Rss2, bool){
	a := Atom1{}
	err := xml.Unmarshal(content, &a)
	if err != nil {
		log.Println(err)
		return Rss2{}, false
	}
	return atom1ToRss2(a), true
}

func parseFeedContent(content []byte) (Rss2, bool) {
	v := Rss2{}
	err := xml.Unmarshal(content, &v)
	if err != nil {
		if err.Error() == atomErrStr {
			// try Atom 1.0
			return parseAtom(content)
		}
		log.Println(err)
		return v, false
	}

	if v.Version == "2.0" {
		// RSS 2.0
		for i, _ := range v.ItemList {
			if v.ItemList[i].Content != "" {
				v.ItemList[i].Description = v.ItemList[i].Content
			}
		}
		return v, true
	}

	log.Println("not RSS 2.0")
	return v, false
}


func main() {
	// parse sample rss feed
	xmlContent1, _ := ioutil.ReadFile("example-6.xml")
	r1, ok1 := parseFeedContent(xmlContent1)
	if ok1 {
		log.Println(r1.Title)
	} else {
		log.Println("fail to read example-6")
	}

	// parse sample atom feed
	xmlContent2, _ := ioutil.ReadFile("example-7.xml")
	r2, ok2 := parseFeedContent(xmlContent2)
	if ok2 {
		log.Println(r2.Title)
	} else {
		log.Println("fail to read example-7")
	}

	// parse opml
	xmlContent3, _ := ioutil.ReadFile("example-5.xml")
	r3, ok3 := parseFeedContent(xmlContent3)
	if ok3 {
		log.Println(r3.Title)
	} else {
		log.Println("fail to read example-5")
	}
}
