/*
Poll RSS/ATOM seeds

For handling Ctrl^C event, @see
http://stackoverflow.com/questions/11268943/golang-is-it-possible-to-capture-a-ctrlc-signal-and-run-a-cleanup-function-in
*/
package mylib

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)


func cleanupDB(ch chan os.Signal) {
	for sig := range ch {
		fmt.Println(sig)
		os.Exit(1)
	}
}

func Poll(sites []OpmlOutline) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go cleanupDB(c)

	for _, site := range sites {
		fmt.Println(site)
	}
}
