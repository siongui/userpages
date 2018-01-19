Count spread in O(n)
####################

:date: 2017-10-12T18:25:00-07:00
:author: Shen-Fu Tsai
:tags: Math, Algorithm
:category: Algorithm
:summary: This is not hard but still interesting. A spread of an array of
          numbers is the difference between its maximum and minimum.
          Given an array...
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  <br/>
  This is not hard but still interesting.<br/>
  <br/>
  A spread of an array of numbers is the difference between its maximum and minimum. Given an array of numbers \(x_1,\ldots,x_n\), compute the sum of spreads of all of its continuous subarrays in \(O(n)\) time.<br/>
  <br/>
  Solution:<br/>
  If we could compute the sum of maximum of all continuous subarrays then it&#39;s easy. To do this, suppose we know this sum for \(x_1,\ldots,x_{n-1}\). When \(x_n\) comes in, we want to add to this sum the sum of maximums of \((x_1,\ldots,x_n),(x_2,\ldots,x_n),\ldots,(x_{n-1},x_n),(x_n)\), or \(y_1+y_2+\ldots+y_n\), where \(y_k\) is the maximum between \(x_k\) and the last number. Clearly \(y\) never goes up, so we could use a stack of triple \((value, right, area)\) to store info we need. \((value, right, area)\) is such that for \(right_{k-1} + 1\leq i\leq right_k\), \(y_i=value_k\), and the cumulative area of \(y\) curve from the beginning to \(right_k\) is \(area_k\).<br/>
  <br/>
  After \(x_n\) comes in, \((value, right, area)\) is updated as this. Pop out all entries where \(value\lt x_n\). Now if the top of the stack has \(value=x_n\), update its \(area\) and \(right\) accordingly. Otherwise push a new entry with \(value=x_n\) and \(right=n\). Its \(area\) should be the preceding \(area\) plus \(x_n\times\delta\) where \(\delta\) is \(n\) minus the preceding \(right\).<br/>
  <br/>
  After \((value, right, area)\) is updated, add the last \(area\) to the sum. The \(pop\) takes amortized \(O(1)\) per element, and the rest takes a fixed \(O(1)\).<br/>
  <div>
  <br/></div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/10/count-spread-in-on.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
