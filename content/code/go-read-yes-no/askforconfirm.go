package main

import (
	"fmt"
	"strings"
)

func Ask4confirm() bool {
	var s string

	fmt.Printf("(y/N): ")
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true
	}
	return false
}

func main() {
	isConfirmed := Ask4confirm()
	if isConfirmed {
		fmt.Println("confirmed")
	} else {
		fmt.Println("not confirmed")
	}
}
