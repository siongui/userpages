Radon's theorem
###############

:date: 2017-12-21T18:33:00-08:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary: I came across Radon's theorem the other day, and found this proof. It
          must have been discovered before.
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  I came across Radon&#39;s theorem the other day, and found this proof. It must have been discovered before.<br/>
  <br/>
  Radon&#39;s theorem: Any \(d+2\) points in \(R^d\) can be partition into two disjoint sets whose convex hulls intersect.<br/>
  <br/>
  Proof:<br/>
  Let these points be \(A=\{a_1,a_2,\ldots,a_{d+2}\}\). Consider the point \(x=\frac{1}{d+2}\sum_{i=1}^{d+2}a_i\). By definition \(x\) is in convex hull of \(A\).<br/>
  <br/>
  By Caratheodory&#39;s theorem, \(x\) can be represented as convex combination of at most \(d+1\) points in \(A\). We then have an equation whose left hand side and right hand side do not cancel each other because one side has \(d+2\) non-zero terms and the other has no more than \(d+1\). Moreover after rearrangement we get a point in \(R^d\) which is the convex combination of two disjoint sets of \(A\).<br/>
  <br/>
  Q.E.D.<br/>
  <div>
  <br/></div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/12/radons-theorem.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
