Find pivot from a stream
########################

:date: 2017-09-22T17:42:00-07:00
:author: Shen-Fu Tsai
:tags: Math, Algorithm
:category: Algorithm
:summary: This is also an interesting result, although it's quite simple.
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  This is also an interesting result, although it&#39;s quite simple.<br/>
  <br/>
  You are given a stream of distinct numbers \(x[0], x[1], \ldots, x[n-1]\), and you are to find a pivot number from them, i.e. \(x[i]\) such that \(x[j]\lt x[i]\) iff \(j\lt i\). If there is no such number, report it. How to do it in \(O(n)\) time with \(O(1)\) space? Since it&#39;s a stream, each number can only be accessed once unless it&#39;s stored.<br/>
  <br/>
  Solution:<br/>
  Report the smallest possible pivot number: keep a candidate \(p\) initialized to \(x[0]\). It is invalidated whenever we see a number smaller than itself, which will be then initialized to the next number bigger than the max we have seen so far.<br/>
  <div>
  <br/></div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/09/find-pivot-from-stream.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
