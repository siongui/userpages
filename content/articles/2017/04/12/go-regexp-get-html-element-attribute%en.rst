[Golang] Get HTML Element Attribute via Regular Expression
##########################################################

:date: 2017-04-12T15:04+08:00
:tags: Go, Golang, Regular Expression, String Manipulation
:category: Go
:summary: Get the attribute of a HTML element via *named group matches* in Go
          *regexp* package.
:og_image: https://files.readme.io/G4eDCX6TlKQr9JmwyyVB_RegExCaptureDiag.png
:adsu: yes

Given a short string of a HTML element, we want to extract atrribute of the
element from the given string. We can use `net/html`_ package or goquery_ to get
the job done. But I do not like to ``go get`` packages only for handling a short
string, so I use the `named group matches`_ in Go_ standard regexp_ package.

Problem
+++++++

Given the following string of YouTube embed code:

**<iframe width="560" height="315" src="https://www.youtube.com/embed/YpWFR-ioQlE" frameborder="0" allowfullscreen></iframe>**

Extract the attributes of *width* and *height* from the iframe element.


Solution
++++++++

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/2xA6Bf9JMp>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"errors"
  	"fmt"
  	"regexp"
  )

  const youtubeiframecode = `<iframe width="560" height="315" src="https://www.youtube.com/embed/YpWFR-ioQlE" frameborder="0" allowfullscreen></iframe>`

  func GetAttributes(c string) (w, h string, err error) {
  	pattern := `<iframe width="(?P<w>[0-9]+)" height="(?P<h>[0-9]+)" .*></iframe>`
  	re := regexp.MustCompile(pattern)
  	matches := re.FindStringSubmatch(c)
  	names := re.SubexpNames()

  	for i, match := range matches {
  		if names[i] == "w" {
  			w = match
  		}
  		if names[i] == "h" {
  			h = match
  		}
  	}

  	if w == "" || h == "" {
  		err = errors.New("cannot find attribute")
  		return
  	}

  	return
  }

  func main() {
  	w, h, err := GetAttributes(youtubeiframecode)
  	if err != nil {
  		fmt.Println(err)
  	} else {
  		fmt.Println("width: ", w)
  		fmt.Println("height: ", h)
  	}
  }

.. adsu:: 2

----

References:

.. [1] `[Golang] Regular Expression Named Group - Extract Metadata from File Path <{filename}../../../2016/02/20/go-regexp-named-group-match-path-metadata%en.rst>`_
.. [2] `Online regex tester and debugger: PHP, PCRE, Python, Golang and JavaScript <https://regex101.com/r/relwQD/3>`_
.. [3] | `Extract subexp names in Go regexps : golang <https://www.reddit.com/r/golang/comments/64wcdu/extract_subexp_names_in_go_regexps/>`_
       | `GitHub - reconquest/regexputil-go: Extract subexp names in golang regexp <https://github.com/reconquest/regexputil-go>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _regexp: https://golang.org/pkg/regexp/
.. _named group matches: https://golang.org/pkg/regexp/#Regexp.SubexpNames
.. _net/html: https://godoc.org/golang.org/x/net/html
.. _goquery: https://github.com/PuerkitoBio/goquery
