Online number partitioning
##########################

:date: 2017-09-18T10:12:00-07:00
:author: Shen-Fu Tsai
:tags: Math, Algorithm
:category: Algorithm
:summary: A simple yet surprisingly interesting result.
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  A simple yet surprisingly interesting result.<br/>
  <br/>
  You have a bunch of numbers each in the form of \(2^{-k}\) where \(k&gt;0\), and they sum up to \(1\). You want to partition them into two groups of equal sum, i.e. each group sums up to \(1/2\). It has to be done online: for each number you have to decide which group to throw into, and are not allowed to revisit it again. How?<br/>
  <br/>
  Solution:<br/>
  For each number, if adding it to group \(A\) doesn&#39;t make it exceed \(1/2\), then do it. Otherwise add it to group \(B\). It&#39;ll just work. Correctness is not hard to prove.<br/>
  <div>
  <br/></div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/09/online-number-partitioning.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
