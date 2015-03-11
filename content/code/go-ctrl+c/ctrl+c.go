// http://stackoverflow.com/questions/11268943/golang-is-it-possible-to-capture-a-ctrlc-signal-and-run-a-cleanup-function-in
// https://gobyexample.com/signals
package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func handleCtrlC(c chan os.Signal) {
	sig := <-c
	// handle ctrl+c event here
	// for example, close database
	fmt.Println("\nsignal: ", sig)
	os.Exit(0)
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go handleCtrlC(c)

	// block here
	for { select {} }
}
