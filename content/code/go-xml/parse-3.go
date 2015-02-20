package main

import (
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

type div struct {
	XMLName		xml.Name	`xml:"div"`
	ChildSpan	span
}

type span struct {
	XMLName		xml.Name	`xml:"span"`
	Text		string		`xml:",chardata"`
}

func main() {
	d := div{}
	xmlContent, _ := ioutil.ReadFile("example-3.xml")
	err := xml.Unmarshal(xmlContent, &d)
	if err != nil { panic(err) }
	fmt.Println("ChildSpan:", d.ChildSpan)
}
