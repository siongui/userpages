package main

import (
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

type div struct {
	XMLName		xml.Name	`xml:"div"`
	SpanList	[]span		`xml:"span"`
}

type span struct {
//	XMLName		xml.Name	`xml:"span"`
	Text		string		`xml:",chardata"`
}

func main() {
	d := div{}
	xmlContent, _ := ioutil.ReadFile("example-4.xml")
	err := xml.Unmarshal(xmlContent, &d)
	if err != nil { panic(err) }
	fmt.Println(d)
	// uncomment XMLName struct field in span struct
	// the output of fmt.Println(d) will be:
	// {{ div} [{{ span} SpanText1} {{ span} SpanText2} {{ span} SpanText3}]}
}
