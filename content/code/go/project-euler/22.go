package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

func GetFileContentFromUrl() string {
	resp, err := http.Get("https://projecteuler.net/project/resources/p022_names.txt")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func GetAllSortedNames() []string {
	raw := GetFileContentFromUrl()
	raw = raw[1 : len(raw)-1]
	names := strings.Split(raw, `","`)
	sort.Strings(names)
	return names
}

func main() {
	names := GetAllSortedNames()
	totalscores := 0
	for i, name := range names {
		pos := i + 1

		sum := 0
		for i := 0; i < len(name); i++ {
			sum += int(name[i] - byte('A') + 1)
		}

		score := sum * pos
		totalscores += score
	}
	fmt.Println(totalscores)
}
