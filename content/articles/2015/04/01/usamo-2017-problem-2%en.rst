USAMO 2017 Problem 2
####################

:date: 2017-04-20T22:36:00-07:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary:  Let \(m_1, m_2, \ldots, m_n\) be a collection of \(n\) positive integers, not necessarily distinct. For any sequence of integers \(A = (a_1...
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  Let \(m_1, m_2, \ldots, m_n\) be a collection of \(n\) positive integers, not necessarily distinct. For any sequence of integers \(A = (a_1, \ldots, a_n)\) and any permutation \(w = w_1, \ldots, w_n\) of \(m_1, \ldots, m_n\), define an \(A\)-inversion of \(w\) to be a pair of entries \(w_i, w_j\) with \(i &lt; j\) for which one of the following conditions holds:<br/>
  <br/>
  \(a_i \ge w_i &gt; w_j\),<br/>
  \(w_j &gt; a_i \ge w_i\),<br/>
  \(w_i &gt; w_j &gt; a_i\).<br/>
  <br/>
  Show that, for any two sequences of integers \(A = (a_1, \ldots, a_n)\) and \(B = (b_1, \ldots, b_n)\), and for any positive integer \(k\), the number of permutations of \(m_1, \ldots, m_n\) having exactly \(k\) \(A\)-inversions is equal to the number of permutations of \(m_1, \ldots, m_n\) having exactly \(k\) \(B\)-inversions.<br/>
  <br/>
  ==========================================================<br/>
  <br/>
  Proof:<br/>
  <br/>
  It suffices to show that swapping any \(a_i\) with \(a_{i+1}\) does not alter the distribution of number of \(A\)-inversions, as then any sequence \(A\) could be transformed to any other sequence \(B\) with a combination of such swapping and update of the last element \(a_n\).<br/>
  <div>
  <br/></div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/04/usamo-2017-problem-2.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
