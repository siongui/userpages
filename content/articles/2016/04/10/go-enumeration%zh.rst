Go語言 列舉 型態
################

:date: 2016-04-10T22:27+08:00
:tags: Web開發, Go語言
:category: Go語言
:summary: Go_ 語言 `列舉`_ 型態 寫法
:adsu: yes


Go_ 語言 `列舉`_ 型態 寫法
（來自官網 Iota_ ）：

.. code-block:: go

  const ( // iota is reset to 0
  	c0 = iota  // c0 == 0
  	c1 = iota  // c1 == 1
  	c2 = iota  // c2 == 2
  )

  const ( // iota is reset to 0
  	a = 1 << iota  // a == 1
  	b = 1 << iota  // b == 2
  	c = 3          // c == 3  (iota is not used but still incremented)
  	d = 1 << iota  // d == 8
  )

  const ( // iota is reset to 0
  	u         = iota * 42  // u == 0     (untyped integer constant)
  	v float64 = iota * 42  // v == 42.0  (float64 constant)
  	w         = iota * 42  // w == 84    (untyped integer constant)
  )

  const x = iota  // x == 0  (iota has been reset)
  const y = iota  // y == 0  (iota has been reset)

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `golang enumeration - Google search <https://www.google.com/search?q=golang+enumeration>`_

       `What is an idiomatic way of representing enums in Go? - Stack Overflow <http://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go>`_

.. [2] `程式 列舉 - Google search <https://www.google.com/search?q=%E7%A8%8B%E5%BC%8F+%E5%88%97%E8%88%89>`_

.. [3] `Enumerated type - Wikipedia, the free encyclopedia <https://en.wikipedia.org/wiki/Enumerated_type>`_

       `枚舉 - 維基百科，自由的百科全書 <https://zh.wikipedia.org/zh-tw/%E6%9E%9A%E4%B8%BE>`_

.. _Go: https://golang.org/
.. _列舉: https://en.wikipedia.org/wiki/Enumerated_type
.. _Iota: https://golang.org/ref/spec#Iota
