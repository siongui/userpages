Maximum number of pairwise non-acute vectors in R^n
###################################################

:date: 2017-05-21T13:10:00-07:00
:author: Shen-Fu Tsai
:tags: Math, Algorithm
:category: Math
:summary:  It is said to be well-known, but perhaps I've never seen it before:
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  It is said to be well-known, but perhaps I&#39;ve never seen it before: there are at most \(2n\) non-zero vectors in \(R^n\) such that the inner product of any two of them is not positive.<br/>
  <br/>
  Proof:<br/>
  We could always change the basis such that one of the vectors is \((\delta,0,\ldots,0)\) where \(\delta&gt;0\). Obviously there could not be any other vector with positive first coordinate, i.e. the rest of vectors are grouped into \(A\) and \(B\), where \(A\) consists of vectors with negative first coordinate, and \(B\) zero first coordinate. Express each vector in \((x,v)\) where \(x\in R\). The \(v\)s in \(A\) do not have duplicate, because that would violate the assumption. Similary the \(v\)s in \(B\) are distinct. Moreover \(v\)s from \(A\) and \(B\) are distinct as well. This means that these \(v\)s in \(A\cup B\) are all different and possess the property of pairwise non-positive inner product, except \(A\) could have one zero vector. By induction this finishes the proof.<br/>
  Q.E.D.<br/>
  <br/>
  The induction proof also leads to another interesting and intuitive result: for \(n&gt;1\) optimality happens only with \(n\) mutual orthogonal vectors and their scaled negatives.<br/>
  <div>
  <br/></div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/05/maximum-number-of-pairwise-non-acute.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
