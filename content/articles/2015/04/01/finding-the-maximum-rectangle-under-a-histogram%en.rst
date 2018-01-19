Finding the maximum rectangle under a histogram
###############################################

:date: 2017-09-28T17:31:00-07:00
:author: Shen-Fu Tsai
:tags: Math, Algorithm
:category: Algorithm
:summary: I heard this is classic, but turns out not too hard.
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  I heard this is classic, but turns out not too hard.<br/>
  <br/>
  You&#39;re given non-negative numbers \(h[0],\ldots,h[n-1]\) representing a histogram. Find the maximum area of rectangle beneath it, i.e. \(\max_{i\leq j}(j-i+1)\min_{i\leq k\leq j}h[k]\) in \(O(n)\) time.<br/>
  <br/>
  Solution:<br/>
  Scanning \(h\) from left to right, we maintain a stack of \((height, last, cur)\). \(height\) is height, \(last\) is the last \(x\)-coordinate that is equal or above \(height\), i.e. the leftmost \(x\)-coordinate that covers \(height\), and \(cur\) is current \(x\)-coordinate.<br/>
  <br/>
  The update logic is this. Suppose we&#39;re dealing with \(h[i]\). For each top element with \(height&gt;h[i]\), replace the current maximum area with \(height\times (i-last)\) if the latter is larger. Remove the top element no matter what. Then, if (a) the top element has the same \(height\) as \(h[i]\), update its \(cur\) with \(i\); (b) stack is empty, insert \((h[i], 0, i)\); (c) top element has \(cur=l\), insert \((h[i], l+1, i)\).<br/>
  <br/>
  When we&#39;re done scanning, for each \((height, last, cur)\) in the stack update maximum area with \((n-last)height\) if it&#39;s larger.<br/>
  <div>
  <br/></div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/09/finding-maximum-rectangle-under.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
