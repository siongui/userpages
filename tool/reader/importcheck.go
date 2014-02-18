package main

import (
	"./mylib"
	"fmt"
)


const o1 = "../../content/extra/Feeder.opml"


func checkSite(site mylib.OpmlOutline, c chan int) {
	_, ok := mylib.GetSeed(site.XmlUrl)
	if !ok {
		fmt.Println(site)
	}
	c <- 1
}

// search keyword: golang wait for goroutine to finish
// https://groups.google.com/d/topic/golang-nuts/mNhXnWLFOo4
func main() {
	c := make(chan int)
	sites := mylib.GetOutlineList(o1)
	for _, site := range sites {
		go checkSite(site, c)
	}

	for i:=0 ; i < len(sites); i++ { <- c }
}
