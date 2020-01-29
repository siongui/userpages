[Golang] Check if Environment Variable Exist
############################################

:date: 2020-01-29T08:07+08:00
:tags: Go, Golang, Environment Variable
:category: Go
:summary: Check if an environment variable exists using *os.LookupEnv* in Go.
:og_image: https://i.imgur.com/1eNgrht.jpg
:adsu: yes


Use os.LookupEnv_ in Go standard library to check if an environment variable
exists. the following is an method to help you check:

.. code-block:: go

  import (
  	"os"
  )

  func isEnvExist(key string) bool {
  	if _, ok := os.LookupEnv(key); ok {
  		return true
  	}
  	return false
  }

You can run example code by the following link:

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/9DfCKY8r_Qt>`__
   :class: align-center

----

References:

.. [1] | `golang check if environment variable exists - Google search <https://www.google.com/search?q=golang+check+if+environment+variable+exists>`_
       | `golang check if environment variable exists - DuckDuckGo search <https://duckduckgo.com/?q=golang+check+if+environment+variable+exists>`_
       | `golang check if environment variable exists - Ecosia search <https://www.ecosia.org/search?q=golang+check+if+environment+variable+exists>`_
       | `golang check if environment variable exists - Qwant search <https://www.qwant.com/?q=golang+check+if+environment+variable+exists>`_
       | `golang check if environment variable exists - Bing search <https://www.bing.com/search?q=golang+check+if+environment+variable+exists>`_
       | `golang check if environment variable exists - Yahoo search <https://search.yahoo.com/search?p=golang+check+if+environment+variable+exists>`_
       | `golang check if environment variable exists - Baidu search <https://www.baidu.com/s?wd=golang+check+if+environment+variable+exists>`_
       | `golang check if environment variable exists - Yandex search <https://www.yandex.com/search/?text=golang+check+if+environment+variable+exists>`_

.. _Go: https://golang.org/
.. _os.LookupEnv: https://golang.org/pkg/os/#LookupEnv
