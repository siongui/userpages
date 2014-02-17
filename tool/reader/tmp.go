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
// http://golang.org/src/pkg/encoding/xml/
type Atom1 struct {
	XMLName		xml.Name	`xml:"http://www.w3.org/2005/Atom feed"`
	Title		string		`xml:"title"`
	Subtitle	string		`xml:"subtitle"`
	Id		string		`xml:"id"`
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


const atomErrStr = "expected element type <rss> but have <feed>"

func parseAtom(content []byte) (Atom1, bool){
	a := Atom1{}
	err := xml.Unmarshal(content, &a)
	if err != nil {
		log.Println(err)
		return a, false
	}
	return a, true
}

func parseSeedContent(content []byte) (Rss2, bool) {
	v := Rss2{}
	err := xml.Unmarshal(content, &v)
	if err != nil {
		if err.Error() == atomErrStr {
			// try Atom 1.0
			log.Println(parseAtom(content))
		}
		log.Println(err)
		return v, false
	}

	if v.Version == "2.0" {
		// RSS 2.0
		return v, true
	}

	log.Println("not RSS 2.0")
	return v, false
}

func GetSeed(url string) (Rss2, bool) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return Rss2{}, false
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Println(err2)
		return Rss2{}, false
	}

	return parseSeedContent(body)
}

