[Golang] Check If a String Exists in File
#########################################

:date: 2017-06-05T08:30+08:00
:tags: Go, Golang, Regular Expression, String Manipulation
:category: Go
:summary: Check if a string exists in file via Go regexp package.
:og_image: https://cdn-images-1.medium.com/max/800/1*6iIv3citGMm7eQQrrlukcA.jpeg
:adsu: yes

Check if a string exists in file via Go regexp_ package.

.. code-block:: go

  import (
  	"io/ioutil"
  	"regexp"
  )

  func IsExist(str, filepath string) bool {
  	b, err := ioutil.ReadFile(filepath)
  	if err != nil {
  		panic(err)
  	}

  	isExist, err := regexp.Match(str, b)
  	if err != nil {
  		panic(err)
  	}
  	return isExist
  }

----

References:

.. [1] `func Match - regexp - The Go Programming Language <https://golang.org/pkg/regexp/#Match>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _regexp: https://golang.org/pkg/regexp/
.. _named group matches: https://golang.org/pkg/regexp/#Regexp.SubexpNames
