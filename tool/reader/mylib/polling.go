/*
Poll RSS/ATOM seeds

^C event handling
http://stackoverflow.com/questions/11268943/golang-is-it-possible-to-capture-a-ctrlc-signal-and-run-a-cleanup-function-in

Is there a way to do repetitive tasks at intervals in golang?
http://stackoverflow.com/questions/16466320/is-there-a-way-to-do-repetitive-tasks-at-intervals-in-golang

10 things you (probably) don't know about Go
http://nf.wh3rd.net/10things/

multiple goroutines listening on one channel
http://stackoverflow.com/questions/15715605/multiple-goroutines-listening-on-one-channel
*/
package mylib

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func cleanupDB(ch chan os.Signal, openedDBs chan string) {
	for sig := range ch {
		fmt.Println(sig)
		for {
			select {
			case db := <- openedDBs:
				fmt.Print(db, " closed!\n")
			default:
				os.Exit(1)
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
	count := 0
	for {
		select {
		default:
			fmt.Println(site.Title, count)
			count += 1
			time.Sleep(1 * time.Second)
		}
	}
}
