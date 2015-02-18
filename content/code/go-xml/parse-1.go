// http://golang.org/pkg/io/ioutil/
// http://golang.org/pkg/encoding/xml/
package main

import (
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

type div struct {
	XMLName	xml.Name	`xml:"div"`
	// First letter must be capital. Cannot use inner
	Content	string		`xml:",chardata"`
}

func main() {
	d := div{}
	xmlContent, _ := ioutil.ReadFile("example-1.xml")
	err := xml.Unmarshal(xmlContent, &d)
	if err != nil { panic(err) }
	fmt.Println("XMLName:", d.XMLName)
	fmt.Println("Content:", d.Content)
}
