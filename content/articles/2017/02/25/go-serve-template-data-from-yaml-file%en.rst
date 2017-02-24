[Golang] Serve Template Data From YAML File
###########################################

:date: 2017-02-25T03:30+08:00
:tags: Go, Golang, Golang template, File Input/Output, JSON, YAML
:category: Go
:summary:  Extract template data from human readable YAML_ file and create
           textual output in Go_.
:og_image: http://www.unixstickers.com/image/cache/data/stickers/golang/Go-brown-side.sh-600x600.png
:adsu: yes


Extract template data from human readable YAML_ file and create textual/HTML
output by Golang_ `text/template`_ or `html/template`_ package.

Assume we have the following data stored in YAML_ file:

.. show_github_file:: siongui userpages content/code/go/yaml-template-data/tmpldata.yaml

We want to apply the above data to the template of `text/template`_ or
`html/template`_ package to create textual/HTML output.

First use Google [2]_ to find Go YAML parser. I found `github.com/ghodss/yaml`_
[3]_ is a good choice because we can define a struct once and use it for both
JSON_ and YAML. Install the YAML parser by:

.. code-block:: bash

  $ go get -u github.com/ghodss/yaml

Then we will read the data from YAML file, parse the YAML data, and store the
result in a pre-defined struct, i.e., Unmarshal_ the YAML data. This is what the
following code does:

.. adsu:: 2
.. show_github_file:: siongui userpages content/code/go/yaml-template-data/yml.go
.. adsu:: 3

Now the data is stored in the struct, and we apply the data to Go template and
create output by following code:

.. show_github_file:: siongui userpages content/code/go/yaml-template-data/yml_test.go
.. adsu:: 4

The output is ass follows:

.. code-block:: txt

  === RUN   TestServeFromYAML

  sitename: Theory and Practice
  open graph url: https://siongui.github.io/
  open graph locale: en_US
  --- PASS: TestServeFromYAML (0.00s)
  PASS

----

Tested on:

- Ubuntu Linux 16.10
- Go 1.7.5

----

References
++++++++++

.. [1] `serve template data from yaml · siongui/wat-pah-photiyan@2b48ef6 · GitHub <https://github.com/siongui/wat-pah-photiyan/commit/2b48ef6be45c66cb9299a6badaba5d964ebc8134>`_
.. [2] | `golang yaml - Google search <https://www.google.com/search?q=golang+yaml>`_
       | `golang yaml - DuckDuckGo search <https://duckduckgo.com/?q=golang+yaml>`_
       | `golang yaml - Ecosia search <https://www.ecosia.org/search?q=golang+yaml>`_
       | `golang yaml - Qwant search <https://www.qwant.com/?q=golang+yaml>`_
       | `golang yaml - Bing search <https://www.bing.com/search?q=golang+yaml>`_
       | `golang yaml - Yahoo search <https://search.yahoo.com/search?p=golang+yaml>`_
       | `golang yaml - Baidu search <https://www.baidu.com/s?wd=golang+yaml>`_
       | `golang yaml - Yandex search <https://www.yandex.com/search/?text=golang+yaml>`_
.. adsu:: 5
.. [3] `GitHub - ghodss/yaml: A better way to marshal and unmarshal YAML in Golang <https://github.com/ghodss/yaml>`_
.. [4] `text/template - The Go Programming Language <https://golang.org/pkg/text/template/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _html/template: https://golang.org/pkg/html/template/
.. _text/template: https://golang.org/pkg/text/template/
.. _Execute: https://golang.org/pkg/text/template/#Template.Execute
.. _YAML: https://www.google.com/search?q=YAML
.. _JSON: https://www.google.com/search?q=JSON
.. _Unmarshal: https://godoc.org/github.com/ghodss/yaml#Unmarshal
.. _github.com/ghodss/yaml: https://github.com/ghodss/yaml
