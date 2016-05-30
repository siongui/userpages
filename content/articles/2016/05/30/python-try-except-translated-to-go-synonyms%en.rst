Python try except Translated to Golang Synonyms
###############################################

:date: 2016-05-30T21:13+08:00
:tags: Python, Go, Golang
:category: Python
:summary: Synonyms - Python_ try-except translated to Go_ programming language.

Synonyms - Python_ try-except translated to Golang_
(Go_ programming language).

Python_:

.. code-block:: python

  try:
      one()
      two()
      three()
  except:
      handler()

Go_:

.. code-block:: go

  func main() {
  	defer func() {
  		e := recover()
  		if e != nil {
  			handler()
  		}
  	}()
  	one()
  	two()
  	three()
  }

I believe the original version comes from [3]_, but the post seems to be
missing.

----

References:

.. [1] `Go 语言的错误处理机制是一个优秀的设计吗？ - 异常处理 - 知乎 <https://www.zhihu.com/question/27158146>`_

       .. image:: https://pic3.zhimg.com/c5908ed64e95112aa1e82b47cf30a67a_r.jpg
          :alt: Python try except Translated to Golang Synonyms

.. [2] `HTML5note/go笔记.txt at master · summerworm/HTML5note · GitHub <https://github.com/summerworm/HTML5note/blob/master/go%E7%AC%94%E8%AE%B0.txt#L2821>`_

.. [3] `科技博 on Twitter: "Golang Go语言中的try ... except ...统一的异常处理: 写python的时候，有时候懒得一个一个错误去分别处理，直接使用一个整体来try catch来捕获所有异常。 http://tinyurl.com/3ruyctw <https://twitter.com/kejibo/status/70987551355310080>`_

.. [4] `golang recover - Google search <https://www.google.com/search?q=golang+recover>`_

       `exception handling - “Catching” panics in Go lang - Stack Overflow <http://stackoverflow.com/questions/25025467/catching-panics-in-go-lang>`_

       `return - Go: returning from defer - Stack Overflow <http://stackoverflow.com/questions/19934641/go-returning-from-defer>`_


.. _Python: https://www.python.org/
.. _Go: https://golang.org/
.. _Golang: https://golang.org/
