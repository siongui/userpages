package main

import (
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

type div struct {
	XMLName	xml.Name	`xml:"div"`
	Class	string		`xml:"class,attr"`
	Content	string		`xml:",chardata"`
}

func main() {
	d := div{}
	xmlContent, _ := ioutil.ReadFile("example-2.xml")
	err := xml.Unmarshal(xmlContent, &d)
	if err != nil { panic(err) }
	fmt.Println("XMLName:", d.XMLName)
	fmt.Println("Class:", d.Class)
	fmt.Println("Content:", d.Content)
}
