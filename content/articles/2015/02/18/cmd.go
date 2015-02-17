// http://golang.org/pkg/flag/
package main

import (
	"flag"
	"fmt"
)

var isProductonServer = flag.Bool("production", false, "if run in production mode")

func main() {
	flag.Parse()
	fmt.Println("Production Mode:", *isProductonServer)
}
