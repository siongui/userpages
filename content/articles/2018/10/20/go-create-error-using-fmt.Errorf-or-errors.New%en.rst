[Golang] Create Error Using fmt.Errorf or errors.New
####################################################

:date: 2018-10-20T02:58+08:00
:tags: Go, Golang, Error Handling
:category: Go
:summary: Create new or custom error using *fmt.Errorf* or *errors.New* in Go.
:og_image: http://cdn.codesamplez.com/wp-content/uploads/2016/01/golang-error-handling.jpg
:adsu: yes


In Go_, we have two easy methods to create new or custom error - either
errors.New_ or fmt.Errorf_ in Go standard library. According to discussion in
reddit [2]_, these two methods are identical except fmt.Errorf_ allows you to
format error message string. We will illustrate the difference by the example of
handing HTTP response status code.


errors.New_
+++++++++++

.. code-block:: go

  if resp.StatusCode != 200 {
  	err = errors.New("response status code != 200")
  	return
  }


fmt.Errorf_
+++++++++++

.. code-block:: go

  if resp.StatusCode != 200 {
  	err = fmt.Errorf("response status code: %d", resp.StatusCode)
  	return
  }


As you can see from the above two code snippets, errors.New_ returns only a
fixed string of error message, while fmt.Errorf_ returns a variable string of
error message which tells the staus code of the HTTP response. This example
illustrates the difference of errors.New_ and fmt.Errorf_.

----

.. adsu:: 2

References:

.. [1] `golang fmt.errorf at DuckDuckGo <https://duckduckgo.com/?q=golang+fmt.errorf>`_
.. [2] `fmt.Errorf() or errors.New()? : golang <https://old.reddit.com/r/golang/comments/6ffrie/fmterrorf_or_errorsnew/>`_

.. _Go: https://golang.org/
.. _errors.New: https://golang.org/pkg/errors/#New
.. _fmt.Errorf: https://golang.org/pkg/fmt/#Errorf
