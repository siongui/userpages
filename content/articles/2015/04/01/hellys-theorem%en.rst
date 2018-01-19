Helly's theorem
###############

:date: 2018-01-04T16:48:00-08:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary: This is the last of a small series of similar and basic results in
          convex geometry.
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  This is the last of a small series of similar and basic results in convex geometry.<br/>
  <br/>
  Helly&#39;s theorem: Let \(X_1,\ldots,X_n\) be a finite collection of convex subsets of \(R^d\). If the intersection of every \(d+1\) of these sets is nonempty, then the whole collection has a nonempty intersection.<br/>
  <br/>
  Proof:<br/>
  Below we show that if every \(k\geq d+1\) sets have nonempty intersection, then so do every \(k+1\) sets.<br/>
  <br/>
  It suffices to consider only \(k+1\geq d+2\) convex sets \(X_1,\ldots,X_{k+1}\).<br/>
  <br/>
  Let \(C_i=\bigcap_{v\neq i}X_v\)<br/>
  <br/>
  Clearly if any two distinct \(C_i\) and \(C_j\) overlap then we are done, so assume they don&#39;t and pick a representative point \(p_i\) for \(C_i\) for each \(i\). Now apply Radon&#39;s theorem to \(p_1,\ldots,p_{k+1}\), i.e. without loss of generality there is a point \(x\) that is convex combination of \(p_1,\ldots,p_m\), and also convex combination of \(p_{m+1},\ldots,p_{k+1}\). For any \(i\in[1,m]\), the latter ensures \(x\in X_i\), and for any \(i\in[m+1,k+1]\), the former implies the same thing.<br/>
  Q.E.D.<br/>
  <div>
  <br/></div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2018/01/hellys-theorem.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
