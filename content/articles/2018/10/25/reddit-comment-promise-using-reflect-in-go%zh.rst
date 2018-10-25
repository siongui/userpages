Reddit上關於利用Go語言反射機制實作Promise的評論
###############################################

:date: 2018-10-25T22:52+08:00
:tags: Go語言, JavaScript
:category: Go語言
:summary: Reddit上關於利用Go語言反射機制實作Promise的評論。
:og_image: https://ws1.sinaimg.cn/large/687148dbly1fo6mhg1xw3j20mf0c10v1.jpg
:adsu: yes


在 Reddit [1]_ 上看到的有趣評論，有位網友利用 Golang 的 reflect_ 實作了 Promise_
的機制。有位網友解釋了為何不要在Go語言裡使用Promise。因為 Promise_ 是單線程語言
(single-threaded language，例如JavaScript)無法多工所以創造出來的機制。
Go語言本身就支援多工，透過利用 channel_ 的機制即可實現 Promise_ 可以做的事，
而且是以更清楚易懂的方式來實現。

有興趣更深入了解者，可看參考裡的評論（英文）。

.. adsu:: 2

----

參考：

.. [1] `Promises Using Reflect : golang <https://old.reddit.com/r/golang/comments/9r50bb/promises_using_reflect/>`_

.. _reflect: https://golang.org/pkg/reflect/
.. _Promise: http://liubin.org/promises-book/
.. _channel: https://tour.golang.org/concurrency/2
