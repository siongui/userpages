package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "my 404 page!")
}

func FileServerWithCustom404(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			NotFound(w, r)
			return
		}
		fsh.ServeHTTP(w, r)
	})
}

func main() {
	http.ListenAndServe(":8080", FileServerWithCustom404(http.Dir("/usr/share/doc")))
}
