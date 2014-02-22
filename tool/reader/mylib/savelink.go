// http://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-denoted-by-a-path-exists-in-golang
// http://stackoverflow.com/questions/14249467/os-mkdir-and-os-mkdirall-permission-value
// http://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
package mylib

import (
	"net/url"
	"encoding/json"
	"os"
	"io/ioutil"
	"fmt"
	"log"
)

type LinkInfo struct {
	// required
	Title		string
	Link		string
	Tag		string
	Language	string
	// optional
	HNComments	string
}

func SaveLinkAsJson(l LinkInfo) {
	const linkdir = "./links/"
	if _, err := os.Stat(linkdir); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(linkdir, 0755)
		} else {
			log.Println(err)
		}
	}

	path := fmt.Sprint(linkdir, url.QueryEscape(l.Link))
	os.Remove(path)

	b, err := json.Marshal(l)
	if err != nil { log.Println(err) }

	ioutil.WriteFile(path, b, 0644)
}
