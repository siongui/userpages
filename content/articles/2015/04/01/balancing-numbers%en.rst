Balancing numbers
#################

:date: 2017-09-07T17:23:00-07:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary: There are 2n+1 coins each associated with a weight. When we remove
          any coin, we can split the rest into two piles each with n coins...
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  There are \(2n+1\) coins each associated with a weight. When we remove any coin, we can split the rest into two piles each with \(n\) coins such that the sum of one pile equals the sum of the other. Prove that all coins have the same weight.<br/>
  <br/>
  Proof:<br/>
  First I got this big hint to consider only integers, for which it is simpler: every number should have the same parity (again?!), and since adding a constant to all numbers or dividing them all by \(2\) when they are even does not remove the property, we can keep doing these to them. It is clear that initially all numbers are the same.<br/>
  <br/>
  So why could there be no non-integer not all the same that hold the property? This has something to do with linear algebra. Essentially we are seeking the null space of a certain linear operator \(A\), and we want to prove it&#39;s \(0\). \(A\) consists of integer coefficients, and if its null space is not \(0\), then it certainly contains non-zero points with rational and therefore integer coordinates, a contradiction.<br/>
  <div>
  <br/></div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/09/balancing-numbers.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
