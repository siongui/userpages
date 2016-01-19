package main

import "net/http"
import "encoding/json"

type Def struct {
	Data	[][]string
	Word	string
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
	println(w.Data[0][0])
	println(w.Data[0][1])
	println(w.Data[0][2])
}
