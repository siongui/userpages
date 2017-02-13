[Golang] Hacker News Link to reStructuredText
#############################################

:date: 2016-04-04T22:04+08:00
:tags: Go, Golang, Golang template, goquery, Web Scrape, reStructuredText,
       Copy to Clipboard, Go net/http
:category: Go
:summary: Extract title and URL from `Hacker News comment`_ via goquery_, and
          then output the info to `reStructuredText footnote`_.
:adsu: yes


Extract title and URL from `Hacker News comment`_ via goquery_, and then output
the info to `reStructuredText footnote`_.

Install Go_ and goquery_ package first. Then ``go run`` the following code:


.. code-block:: go

  package main

  import (
  	"bytes"
  	"github.com/PuerkitoBio/goquery"
  	"html/template"
  	"net/http"
  )

  type HNews struct {
  	Title       string
  	URL         string
  	CommentsURL string
  }

  type Index struct {
  	Textarea template.HTML
  }

  var htmlTmpl = `<!doctype html><html>
  <head><title>Link to Rst Image</title></head>
  <body>
  <form action="/" method="post">
    <input size="100" name="url" placeholder="HN comments URL">
    <button>Send</button><br><br>
    <textarea id="ta" rows="5" cols="80">{{.Textarea}}</textarea><br>
    <button type="button" id="copy">Copy textarea to clipboard</button>
  </form>
  <script>
    var textareaElm = document.getElementById("ta");
    var copyElm = document.getElementById("copy");
    copyElm.onclick = function(event) {
      textareaElm.select();
      var isSuccessful = document.execCommand('copy');
      if (isSuccessful) {
        textareaElm.value = "Copy OK";
      } else {
        textareaElm.value = "Copy Fail";
      }
    }
  </script>
  </body></html>`

  var rstTmpl = ".. [1] `{{.Title}} <{{.URL}}>`_\n" +
  	"       (`HN comments <{{.CommentsURL}}>`__)\n"

  func processHNURL(url string) string {
  	doc, err := goquery.NewDocument(url)
  	if err != nil {
  		panic(err)
  	}

  	news := HNews{CommentsURL: url}
  	newsElm := doc.Find("td.title > a")
  	news.Title = newsElm.Text()
  	news.URL, _ = newsElm.Attr("href")

  	tmpl, err := template.New("hn").Parse(rstTmpl)
  	if err != nil {
  		panic(err)
  	}

  	var rst bytes.Buffer
  	err = tmpl.Execute(&rst, news)
  	if err != nil {
  		panic(err)
  	}

  	return rst.String()
  }

  func handler(w http.ResponseWriter, r *http.Request) {
  	idx := Index{}
  	if r.Method == "POST" {
  		hnurl := r.PostFormValue("url")
  		idx.Textarea = template.HTML(processHNURL(hnurl))
  	}
  	tmpl, err := template.New("index").Parse(htmlTmpl)
  	if err != nil {
  		panic(err)
  	}
  	tmpl.Execute(w, idx)
  }

  func main() {
  	http.HandleFunc("/", handler)
  	http.ListenAndServe(":8000", nil)
  }

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `jquery find direct child - Google search <https://www.google.com/search?q=jquery+find+direct+child>`_

       `jquery find direct child - DuckDuckGo search <https://duckduckgo.com/?q=jquery+find+direct+child>`_

       `jquery find direct child - Bing search <https://www.bing.com/search?q=jquery+find+direct+child>`_

       `jquery find direct child - Yahoo search <https://search.yahoo.com/search?p=jquery+find+direct+child>`_

       `jquery find direct child - Baidu search <https://www.baidu.com/s?wd=jquery+find+direct+child>`_

       `jquery find direct child - Yandex search <https://www.yandex.com/search/?text=jquery+find+direct+child>`_

       `How to get only direct child elements by jQuery function - Stack Overflow <http://stackoverflow.com/questions/3687637/how-to-get-only-direct-child-elements-by-jquery-function>`_

.. [2] `golang backtick - Google search <https://www.google.com/search?q=golang+backtick>`_

       `golang backtick - DuckDuckGo search <https://duckduckgo.com/?q=golang+backtick>`_

       `golang backtick - Bing search <https://www.bing.com/search?q=golang+backtick>`_

       `golang backtick - Yahoo search <https://search.yahoo.com/search?p=golang+backtick>`_

       `go - Golang how to escape back ticks - Stack Overflow <http://stackoverflow.com/questions/21198980/golang-how-to-escape-back-ticks>`_

       `How do you write multiline strings in Go? - Stack Overflow <http://stackoverflow.com/questions/7933460/how-do-you-write-multiline-strings-in-go>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _goquery: https://github.com/PuerkitoBio/goquery
.. _Hacker News comment: https://news.ycombinator.com/item?id=11410894
.. _reStructuredText footnote: http://docutils.sourceforge.net/docs/user/rst/quickref.html#footnotes
