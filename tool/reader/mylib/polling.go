/*
Poll RSS/ATOM seeds

^C event handling
http://stackoverflow.com/questions/11268943/golang-is-it-possible-to-capture-a-ctrlc-signal-and-run-a-cleanup-function-in

A Tour of Go - Concurrency
http://tour.golang.org/#64
*/
package mylib

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)


const polling_interval = 10 * time.Second

func cleanupDB(ch chan os.Signal, openedDBs chan string) {
	for sig := range ch {
		for {
			select {
			case db := <- openedDBs:
				log.Println(db, " closed!")
			default:
				log.Fatal(sig)
			}
		}

	}
}

func Poll(sites []OpmlOutline) {
	openedDBs := make(chan string, len(sites))
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go cleanupDB(c, openedDBs)

	for _, site := range sites {
		go getSiteSeed(site, openedDBs)
	}
}

func getSiteSeed(site OpmlOutline, openedDBs chan string) {
	openedDBs <- site.Title
	for {
		select {
		default:
			v := GetSeed(site.XmlUrl)
			for _, item := range v.Channel.ItemList {
				log.Println(item)
			}
			time.Sleep(polling_interval)
		}
	}
}
