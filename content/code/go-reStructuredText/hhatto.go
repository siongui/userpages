package main

import (
	"bufio"
	"bytes"
	"github.com/hhatto/gorst"
	"io/ioutil"
	"os"
)

func main() {
	p := rst.NewParser(nil)

	w := bufio.NewWriter(os.Stdout)
	d, _ := ioutil.ReadFile("test.rst")
	p.ReStructuredText(bytes.NewReader(d), rst.ToHTML(w))
	w.Flush()
}
