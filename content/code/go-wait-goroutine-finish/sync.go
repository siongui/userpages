// search keyword: golang wait for goroutine to finish
// https://groups.google.com/d/topic/golang-nuts/mNhXnWLFOo4
package main

import (
	"fmt"
)


func routine(site string, c chan int) {
	// do something
	fmt.Println(site)
	c <- 1
}


func main() {
	c := make(chan int)

	sites := []string{
	"https://www.google.com/",
	"https://duckduckgo.com/",
	"https://www.bing.com/"}

	for _, site := range sites {
		go routine(site, c)
	}

	// wait all goroutines to finish
	for i:=0; i < len(sites); i++ { <-c }
}
