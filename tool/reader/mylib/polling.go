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
	"database/sql"
)


const polling_interval = 10 * time.Second

type dbInfo struct {
	site	OpmlOutline
	db	*sql.DB
}

func cleanupDB(ch chan os.Signal, openedDBs chan dbInfo) {
	for sig := range ch {
		for {
			select {
			case di := <- openedDBs:
				di.db.Close()
				log.Println(di.site.Title, " db closed!")
			default:
				log.Fatal(sig)
			}
		}

	}
}

func Poll(sites []OpmlOutline) {
	openedDBs := make(chan dbInfo, len(sites))
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go cleanupDB(c, openedDBs)

	for _, site := range sites {
		go getSiteSeed(site, openedDBs)
	}
}

func getSiteSeed(site OpmlOutline, openedDBs chan dbInfo) {
	db := InitDB(XmlUrl2DBPath(site.XmlUrl))
	openedDBs <- dbInfo{site, db}
	for {
		select {
		default:
			v := GetSeed(site.XmlUrl)
			storeItems(db, v.Channel.ItemList)
			log.Println("store: ", site.Title)
			time.Sleep(polling_interval)
		}
	}
}
