// http://golang.org/pkg/io/ioutil/
// http://golang.org/pkg/encoding/xml/
package main

import (
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

type html struct {
	XMLName	xml.Name	`xml:"html"`
	// First letter must be capital. Cannot use inner
	Inner	string		`xml:",innerxml"`
}

func main() {
	h := html{}
	content, _ := ioutil.ReadFile("example-1.xml")
	err := xml.Unmarshal(content, &h)
	if err != nil { panic(err) }
	fmt.Println(h)
}
