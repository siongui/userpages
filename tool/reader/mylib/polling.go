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


var polling_interval = 8 * time.Minute

type dbInfo struct {
	site	OpmlOutline
	db	*sql.DB
}

var dbMap map[string]dbInfo

func setTestPollingInterval() {
	polling_interval = 5 * time.Second
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
	dbMap = make(map[string]dbInfo)

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
	di := dbInfo{site, db}
	openedDBs <- di
	dbMap[site.XmlUrl] = di
	for {
		select {
		default:
			v, ok := GetSeed(site.XmlUrl)
			if ok {
				storeItems(db, v.ItemList)
				//log.Println("store: ", site.Title)
			} else {
				log.Println("fail to get: ", site.Title)
			}
			time.Sleep(polling_interval)
		}
	}
}

func GetItemsFromDB(xmlUrl string) []Item {
	di, ok := dbMap[xmlUrl]
	if !ok { return nil }
	return ReadItems(di.db)
}
