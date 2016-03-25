Vim編輯器將檔案格式從windows(dos)轉成unix(linux)
################################################

:date: 2016-03-26T06:11+08:00
:tags: Linux, Vim
:category: Linux
:summary: Vim_ 編輯器將檔案格式從windows(dos)轉成unix(linux)


Vim_ 編輯器將檔案格式從windows(dos)轉成unix(linux)，從 [1]_ 的回答：

.. code-block:: vim

  :set ff=unix

----

References:

.. [1] `vim dos to unix conversion <https://www.google.com/search?q=vim+dos+to+unix+conversion>`_

       `Convert DOS line endings to Linux line endings in vim - Stack Overflow <http://stackoverflow.com/a/82743>`_

.. _Vim: http://www.vim.org/

.. ``(註1-1)`` => `` [1]_ ``
   :%s/(註1-\(\d\{1}\))/ [\1]_ /gc
   https://www.google.com/search?q=vim+regex+replace
   http://stackoverflow.com/questions/11850033/vim-regex-replace

.. ``〔註1-1〕：`` => `` .. [1] ``
   :%s/^〔註1-\(\d\{1}\)〕：/.. [\1] /gc

.. go read form post value
  package main
  import (
  	"fmt"
  	"net/http"
  )
  var indexHtml = `<!doctype html>
  <html>
  <head><title>Link to Rst Image</title></head>
  <body>
  <form action="/post" method="post">
    <input name="myValue">
    <button>Send</button>
  </form>
  </body>
  </html>`
  func handler(w http.ResponseWriter, r *http.Request) {
  	fmt.Fprintf(w, indexHtml)
  }
  func postHandler(w http.ResponseWriter, r *http.Request) {
  	fmt.Fprintf(w, r.PostFormValue("myValue"))
  }
  func main() {
  	http.HandleFunc("/", handler)
  	http.HandleFunc("/post", postHandler)
  	http.ListenAndServe(":8000", nil)
  }
