package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	cdp "github.com/knq/chromedp"
	cdptypes "github.com/knq/chromedp/cdp"

	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

func getCredentials() (u, p string) {
	fmt.Print("Username: ")
	_, err := fmt.Scan(&u)
	if err != nil {
		panic(err)
	}

	fmt.Print("Password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	p = string(password)

	return
}

func facebookLogin(username, password string) cdp.Tasks {
	selname := `//input[@id="email"]`
	selpass := `//input[@id="pass"]`
	var buf []byte

	return cdp.Tasks{
		cdp.Navigate(`https://www.facebook.com`),
		cdp.WaitVisible(selpass),
		cdp.SendKeys(selname, username),
		cdp.SendKeys(selpass, password),
		cdp.Submit(selpass),
		cdp.WaitVisible(`//a[@title="Profile"]`),
		cdp.CaptureScreenshot(&buf),
		cdp.ActionFunc(func(context.Context, cdptypes.Handler) error {
			return ioutil.WriteFile("myfb.png", buf, 0644)
		}),
	}
}

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := cdp.New(ctxt, cdp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	err = c.Run(ctxt, facebookLogin(getCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
