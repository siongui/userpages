package main

import (
	"fmt"
	"os"
)

func main() {
	fo, err := os.Create("hello.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	fmt.Fprintf(fo, "Hello\n")
	fmt.Fprintf(fo, "World\n")
}
