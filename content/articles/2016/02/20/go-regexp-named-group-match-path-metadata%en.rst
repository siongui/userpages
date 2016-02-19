[Golang] Regular Expression Named Group - Extract Metadata from File Path
#########################################################################

:date: 2016-02-20T02:23+08:00
:tags: Go, Golang, Regular Expression, Pelican, String Manipulation
:category: Go
:summary: Go_ regexp_ example of `named group matches`_ - extract metadata from
          file path.


Golang_ regexp_ example of `named group matches`_ - extract metadata from file
path.


Problem
+++++++

We have a *string* of file path:

**articles/2016/01/05/deploy-website-by-pelican-travis-ci-github-pages%en.rst**

We want to extract the metadata from the string in the following form:

*date* - **2016/01/05**

*slug* - **deploy-website-by-pelican-travis-ci-github-pages**

*lang* - **en**


Solution
++++++++

`Run Code on Go Playground <https://play.golang.org/p/cXgukkhTTu>`__

.. code-block:: go

  package main

  import (
          "fmt"
          "regexp"
  )

  var path1 = "articles/2016/01/05/deploy-website-by-pelican-travis-ci-github-pages%en.rst"
  var path2 = "articles/2016/01/21/everything-is-teaching-us-ajahn-chah%zh.rst"

  func main() {
          pattern := `articles/(?P<date>\d{4}/\d{2}/\d{2})/(?P<slug>[-a-zA-Z0-9]*)%(?P<lang>[_a-zA-Z]{2,5})\.rst`
          pathMetadata := regexp.MustCompile(pattern)

          matches := pathMetadata.FindStringSubmatch(path1)
          names := pathMetadata.SubexpNames()
          for i, match := range matches {
                  if i != 0 {
                          fmt.Println(names[i], match)
                  }
          }

          matches = pathMetadata.FindStringSubmatch(path2)
          names = pathMetadata.SubexpNames()
          for i, match := range matches {
                  if i != 0 {
                          fmt.Println(names[i], match)
                  }
          }
  }


Output
++++++

Output of above code:

.. code-block:: txt

  date 2016/01/05
  slug deploy-website-by-pelican-travis-ci-github-pages
  lang en
  date 2016/01/21
  slug everything-is-teaching-us-ajahn-chah
  lang zh

----

References:

.. [1] `python regular expression ?P <https://www.google.com/search?q=python+regular+expression+%3FP>`_

.. [2] `golang named regular expression <https://www.google.com/search?q=golang+named+regular+expression>`_

.. [3] `Go Playground - golang named path metadata <https://play.golang.org/p/cXgukkhTTu>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _regexp: https://golang.org/pkg/regexp/
.. _named group matches: https://golang.org/pkg/regexp/#Regexp.SubexpNames
