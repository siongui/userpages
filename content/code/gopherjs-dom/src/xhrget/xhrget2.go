package main

import "net/http"
import "encoding/json"

type Exp [3]string

type Def struct {
	Data []Exp
	Word string
}

func main() {
	resp, err := http.Get("/sukhada.json")
	if err != nil {
		// handle error
		println(err)
	}
	defer resp.Body.Close()
	w := Def{}
	dec := json.NewDecoder(resp.Body)
	err2 := dec.Decode(&w)
	if err2 != nil {
		// handle error
		println(err2)
	}
	println(w.Word)
	println(w.Data[0])
}
