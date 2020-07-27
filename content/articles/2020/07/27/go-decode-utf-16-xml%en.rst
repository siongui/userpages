[Golang] Unmarshal UTF-16 XML
#############################

:date: 2020-07-27T10:19+08:00
:tags: Go, Golang, String Manipulation, Locale, i18n, XML
:category: Go
:summary: How to parse a UTF-16 XML file in Go.
:og_image: https://miro.medium.com/max/1250/1*Vu-esgd_xQFi4B5tIxtMqQ.jpeg
:adsu: yes


How to parse a UTF-16 XML file in Go.

The Go standard `encoding/xml`_ package provides basic methods for parsing XML
files. However, the `encoding/xml`_ package parses only UTF-8 encoded XML. After
spending a lot of time for search and trial and error [1]_ [2]_ [3]_, finally I
get my UTF-16 XML parsed correctly. I do not know why it works. Just show the
code for reference.

.. code-block:: go

  import (
  	"encoding/xml"
  	"io"

  	"golang.org/x/net/html/charset"
  )

  func BypassReader(label string, input io.Reader) (io.Reader, error) {
  	return input, nil
  }

  func DecodeUtf16XML(r io.Reader, v interface{}) (err error) {
  	// https://www.tipitaka.org/romn/cscd/vin01m.mul.toc.xml
  	// The Tipiṭaka XML is encoded in UTF-16
  	// Google search: golang xml utf-16
  	// https://stackoverflow.com/questions/6002619/unmarshal-an-iso-8859-1-xml-input-in-go
  	// https://groups.google.com/forum/#!topic/golang-nuts/tXcECEKC2rs
  	nr, err := charset.NewReader(r, "utf-16")
  	if err != nil {
  		return
  	}
  	decoder := xml.NewDecoder(nr)
  	decoder.CharsetReader = BypassReader
  	err = decoder.Decode(v)
  	return
  }

**Usage**:

.. code-block:: go

  import (
  	"encoding/xml"
  	"os"
  	"testing"
  )

  type Tree struct {
  	XMLName xml.Name `xml:"tree"`
  	Trees   []Tree   `xml:"tree"`
  	Text    string   `xml:"text,attr"`
  	Src     string   `xml:"src,attr"`
  	Action  string   `xml:"action,attr"`
  }

  func TestDecodeUtf16XML(t *testing.T) {
  	dst := "/tmp/romn/cscd/vin01m.mul.toc.xml"
  	err := CheckDownload("https://www.tipitaka.org/romn/cscd/vin01m.mul.toc.xml", dst, false)
  	if err != nil {
  		t.Error(err)
  		return
  	}

  	f16, err := os.Open(dst)
  	if err != nil {
  		t.Error(err)
  		return
  	}

  	tree := Tree{}
  	err = DecodeUtf16XML(f16, &tree)
  	if err != nil {
  		t.Error(err)
  		return
  	}
  	t.Log(tree)
  }

----

Tested on: ``Ubuntu Linux 20.04``, ``Go 1.12.17, 1.14.4``.

----

References:

.. [1] | `golang utf-16 xml - Google search <https://www.google.com/search?q=golang+utf-16+xml>`_
       | `golang utf-16 xml - DuckDuckGo search <https://duckduckgo.com/?q=golang+utf-16+xml>`_
       | `golang utf-16 xml - Ecosia search <https://www.ecosia.org/search?q=golang+utf-16+xml>`_
       | `golang utf-16 xml - Qwant search <https://www.qwant.com/?q=golang+utf-16+xml>`_
       | `golang utf-16 xml - Bing search <https://www.bing.com/search?q=golang+utf-16+xml>`_
       | `golang utf-16 xml - Yahoo search <https://search.yahoo.com/search?p=golang+utf-16+xml>`_
       | `golang utf-16 xml - Baidu search <https://www.baidu.com/s?wd=golang+utf-16+xml>`_
       | `golang utf-16 xml - Yandex search <https://www.yandex.com/search/?text=golang+utf-16+xml>`_

.. [2] `utf 8 - Unmarshal an ISO-8859-1 XML input in Go - Stack Overflow <https://stackoverflow.com/questions/6002619/unmarshal-an-iso-8859-1-xml-input-in-go>`_

.. [3] | `How to handle UTF-16 LE XML - Google Groups <https://groups.google.com/forum/#!topic/golang-nuts/tXcECEKC2rs>`_
       | `go - xml: encoding "utf-16" declared but Decoder.CharsetReader is nil unmarshal successful - Stack Overflow <https://stackoverflow.com/questions/50812137/xml-encoding-utf-16-declared-but-decoder-charsetreader-is-nil-unmarshal-succe>`_

.. [4] `decode utf-16 xml · siongui/gopalilib@370ba1f · GitHub <https://github.com/siongui/gopalilib/commit/370ba1fa6c6d18516e7ff08c354f9d0222d69e6b>`_

.. [5] `[Golang] Auto-Detect and Convert Encoding of HTML to UTF-8 <{filename}/articles/2018/10/27/auto-detect-and-convert-html-encoding-to-utf8-in-go%en.rst>`_

.. _Go: https://golang.org/
.. _encoding/xml: https://golang.org/pkg/encoding/xml/
