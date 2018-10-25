package guess

import (
	"bufio"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
)

func UrlEncoding(url string) (name string, certain bool, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("response status code: %d", resp.StatusCode)
		return
	}

	_, name, certain, err = DetermineEncodingFromReader(resp.Body)
	return
}

func DetermineEncodingFromReader(r io.Reader) (e encoding.Encoding, name string, certain bool, err error) {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return
	}

	e, name, certain = charset.DetermineEncoding(bytes, "")
	return
}
