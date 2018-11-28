[Golang] JSON Unmarshal When Field Name is Determined at Runtime
################################################################

:date: 2018-05-13T23:31+08:00
:tags: Go, Golang, JSON
:category: Go
:summary: Unmarshal JSON data in which the field name is determined at runtime
          in Go.
:og_image: https://i.ytimg.com/vi/XkTwDPzeF9k/maxresdefault.jpg
:adsu: yes


Today I got a problem about unmarshalling JSON data. One of the field names in
the JSON data can only be known at runtime. The JSON data looks like:

.. code-block:: json

  {
    "reels": {
      "highlight:1234567890": {
        ...
      }
    }
  }

The field name *highlight:1234567890* is the id of the query, and is known only
at runtime. But to unmarshal JSON using json.Unmarshal_ in Go standard
`encoding/json`_ package, you have to fill the JSON tag or name in the struct at
compile-time. It seems impossible to decode the JSON data in this case, so I
tried to search for the answers [1]_. And the answer of the question [2]_ gives
me the direction of how to unmarshal JSON in such cases.

The trick is **replace runtime field name with fixed string**!

The field name given at runtime is very unique, so it is very easy to replace
the field name with another fixed string using strings.Replace_ in Go standard
strings_ package.

Im my case, the field name *highlight:1234567890* is stored in the variable
*id*, and we can replace it with another fixed string *reels_media* as follows:

.. code-block:: go

  strings.Replace(jsonData, id, "reels_media", 1)

Now in our struct used to unmarshal JSON, we can declare as follows:

.. code-block:: go

  type highlightMedia struct {
  	Reels struct {
  		ReelsMedia struct{
			...
  		} `json:"reels_media"`
  	} `json:"reels"`
  }

Now we can use json.Unmarshal_ as usual without any problem!

The other possible way to solve this problem is to retag at runtime, but it
looks not simple, so I gave up. There is a third-party package retag_ which
helps you do this, but I like to use only Go standard library if possible, so I
did not give it a try.

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 18.04``, ``Go 1.10.2``.

References:

.. [1] | `golang runtime determine json tag - Google search <https://www.google.com/search?q=golang+runtime+determine+json+tag>`_
       | `golang runtime determine json tag - DuckDuckGo search <https://duckduckgo.com/?q=golang+runtime+determine+json+tag>`_
       | `golang runtime determine json tag - Ecosia search <https://www.ecosia.org/search?q=golang+runtime+determine+json+tag>`_
       | `golang runtime determine json tag - Qwant search <https://www.qwant.com/?q=golang+runtime+determine+json+tag>`_
       | `golang runtime determine json tag - Bing search <https://www.bing.com/search?q=golang+runtime+determine+json+tag>`_
       | `golang runtime determine json tag - Yahoo search <https://search.yahoo.com/search?p=golang+runtime+determine+json+tag>`_
       | `golang runtime determine json tag - Baidu search <https://www.baidu.com/s?wd=golang+runtime+determine+json+tag>`_
       | `golang runtime determine json tag - Yandex search <https://www.yandex.com/search/?text=golang+runtime+determine+json+tag>`_
.. [2] `go - JSON Unmarshal when structure is determined at runtime - Stack Overflow <https://stackoverflow.com/questions/40145706/json-unmarshal-when-structure-is-determined-at-runtime>`_
.. [3] `go - Override Golang JSON tag values at runtime - Stack Overflow <https://stackoverflow.com/questions/45032514/override-golang-json-tag-values-at-runtime>`_
.. [4] `change code to adapt to change of the format of reels_media story hig… · siongui/instago@6bfcfff · GitHub <https://github.com/siongui/instago/commit/6bfcfff26dbde5fae27cec5d729b0d3847f54561>`_
.. [5] `Noob question regarding runes : golang <https://old.reddit.com/r/golang/comments/a0zbwj/noob_question_regarding_runes/>`_

.. _encoding/json: https://golang.org/pkg/encoding/json/
.. _json.Unmarshal: https://golang.org/pkg/encoding/json/#Unmarshal
.. _strings.Replace: https://golang.org/pkg/strings/#Replace
.. _strings: https://golang.org/pkg/strings/
.. _retag: https://github.com/sevlyar/retag
.. _Go Playground: https://play.golang.org/
