package main

import (
	"flag"
	"fmt"
)

func main() {
	arg := flag.String("input", "defaultValue", "Command line argument")
	flag.Parse()
	fmt.Println(*arg)
}
