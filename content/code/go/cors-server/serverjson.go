package servejson

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", servejson)
}

func servejson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, `{"a":"b","c":"d"}`)
}
