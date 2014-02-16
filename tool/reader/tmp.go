// http://www.w3schools.com/rss/default.asp
// http://www.tutorialspoint.com/rss/what-is-atom.htm
// http://stackoverflow.com/questions/6619717/what-is-the-difference-between-rss-and-atom-feeds
// http://stackoverflow.com/questions/16309944/atom-to-rss-feeds-converter
package tmp

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"html/template"
	"log"
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


// http://en.wikipedia.org/wiki/Atom_(standard)
type Atom1 struct {
	XMLName		xml.Name	`xml:"feed"`
	Xmlns		string		`xml:"xmlns,attr"`
	Title		string		`xml:"title"`
	Subtitle	string		`xml:"subtitle"`
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
	Id		string		`xml:"id"`
	Link		Link		`xml:"link"`
	Author		Author		`xml:"author"`
}


func parseSeedContent(content []byte) Rss2 {
	v := Rss2{}
	err := xml.Unmarshal(content, &v)
	if err != nil {
		log.Println(err)
		return v
	}

	if v.Version == "2.0" {
		// RSS 2.0
		return v
	}
	return v
}

func GetSeed(url string) (Rss2, bool) {
	resp, err := http.Get(url)
	if err != nil { return Rss2{}, false }
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil { return Rss2{}, false }
	return parseSeedContent(body), true
}


func GetAtom(url string) (Atom1, bool) {
	a := Atom1{}
	resp, err := http.Get(url)
	if err != nil { return a, false }
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil { return a, false }

	err3 := xml.Unmarshal(body, &a)
	if err3 != nil {
		log.Println(err3)
		return a, false
	}

	if a.Xmlns == "http://www.w3.org/2005/Atom" {
		// Atom 1.0
		return a, true
	}
	return a, false
}
