package toutf8

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func UrlToUtf8Encoding(url string) (r io.Reader, name string, certain bool, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("response status code: %d", resp.StatusCode)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	e, name, certain, err := DetermineEncodingFromReader(bytes.NewReader(b))
	if err != nil {
		return
	}

	r = transform.NewReader(bytes.NewReader(b), e.NewDecoder())
	return
}

func DetermineEncodingFromReader(r io.Reader) (e encoding.Encoding, name string, certain bool, err error) {
	b, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return
	}

	e, name, certain = charset.DetermineEncoding(b, "")
	return
}
