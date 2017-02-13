[Golang] Online Concatenate and Compress JavaScript Files
#########################################################

:date: 2016-03-07T01:07+08:00
:tags: Go, Golang, JavaScript, Commandline, HTTP POST, Google Closure,
       Minify HTML/CSS/JavaScript, File Input/Output, Go net/http
:category: Go
:summary: Concatenate and compress JavaScript_ files via Go_ programming
          language and online `Google Closure Compiler`_.
:adsu: yes

Concatenate and compress JavaScript_ files via Golang_ and online
`Google Closure Compiler`_.

.. code-block:: go

  package main

  import (
  	"io/ioutil"
  	"net/http"
  	"net/url"
  	"path"
  )

  func minjs(baseDir string, jsFiles []string, outputPath string) {
  	var jsCode []byte
  	for _, file := range jsFiles {
  		jsPath := path.Join(baseDir, file)
  		println("concatenating " + jsPath + " ...")
  		b, err := ioutil.ReadFile(jsPath)
  		if err != nil {
  			panic(err)
  		}
  		jsCode = append(jsCode, b...)
  	}

  	params := url.Values{}
  	//params.Set("code_url", "https://github.com/twnanda/twnanda/raw/master/theme/javascript/tongwen_core.js")
  	//params.Set("code_url", "https://github.com/twnanda/twnanda/raw/master/theme/javascript/tongwen_table_ps2t.js")
  	//params.Set("code_url", "https://github.com/twnanda/twnanda/raw/master/theme/javascript/tongwen_table_pt2s.js")
  	//params.Set("code_url", "https://github.com/twnanda/twnanda/raw/master/theme/javascript/tongwen_table_s2t.js")
  	//params.Set("code_url", "https://github.com/twnanda/twnanda/raw/master/theme/javascript/tongwen_table_t2s.js")
  	params.Set("js_code", string(jsCode))
  	params.Set("compilation_level", "SIMPLE_OPTIMIZATIONS")
  	params.Set("language", "ECMASCRIPT5")
  	params.Set("output_format", "text")
  	params.Set("output_info", "compiled_code")

  	println("\nCompressing combined js online ...")
  	resp, err := http.PostForm("https://closure-compiler.appspot.com/compile", params)
  	if err != nil {
  		panic(err)
  	}
  	defer resp.Body.Close()
  	b, err := ioutil.ReadAll(resp.Body)
  	if err != nil {
  		panic(err)
  	}
  	ioutil.WriteFile(path.Join(baseDir, outputPath), b, 0644)
  }

  func main() {
  	baseDir := "dictionary"
  	jsFiles := []string{
  		"common/app/scripts/services/data/dicBooks.js",
  		"common/app/scripts/services/data/succinctTrie.js",
  		"common/app/scripts/services/data/i18nStrings.js",
  		"app/scripts/app.js",
  		"app/scripts/controllers.js",
  		"app/scripts/directives/inputSuggest.js",
  		"app/scripts/directives/draggableAndEvents.js",
  		"common/app/scripts/services/paliWordJson.js",
  		"common/app/scripts/services/shortExp.js",
  		"common/app/scripts/services/ngBits.js",
  		"common/app/scripts/services/wordSearch.js",
  		"common/app/scripts/directives/dropdown.js",
  		"common/app/scripts/filters/expOrder.js",
  		"common/app/scripts/i18n.js",
  	}
  	minjs(baseDir, jsFiles, "app/all_compiled.js")
  }

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `Getting Started with the API  |  Closure Compiler  |  Google Developers <https://developers.google.com/closure/compiler/docs/gettingstarted_api>`_

.. [2] `[Python] Online Concatenate and Compress JavaScript Files <{filename}../../02/26/online-concatenate-and-compress-js-files%en.rst>`_

.. [3] `[Golang] Concatenate JavaScript Files <{filename}../06/go-concatenate-js-files%en.rst>`_

.. [4] `pali/minjs.go at 07cf02ac49af1674a9a5d4e78ca7aebf2dc456b0 · siongui/pali · GitHub <https://github.com/siongui/pali/blob/07cf02ac49af1674a9a5d4e78ca7aebf2dc456b0/dictionary/minjs.go>`_

.. [5] `url - The Go Programming Language <https://golang.org/pkg/net/url/>`_

.. [6] `http - The Go Programming Language <https://golang.org/pkg/net/http/>`_

.. [7] `go - Make a URL-encoded POST request using \`http.NewRequest(...)\` - Stack Overflow <http://stackoverflow.com/questions/19253469/make-a-url-encoded-post-request-using-http-newrequest>`_

.. [8] `go - How to send a POST request in Golang? - Stack Overflow <http://stackoverflow.com/questions/24493116/how-to-send-a-post-request-in-golang>`_

.. [9] `golang get current file path <https://www.google.com/search?q=golang+get+current+file+path>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _JavaScript: https://www.google.com/search?q=javascript
.. _Google Closure Compiler: https://developers.google.com/closure/compiler/
