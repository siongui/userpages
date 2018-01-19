Descartes' rules of signs
#########################

:date: 2018-01-16T18:02:00-08:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary: Hard to believe that I don't remember seeing Descartes' rules of signs
          before -- you'd think this is impossible given the time I spent on ...
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  Hard to believe that I don&#39;t remember seeing <a href="https://en.wikipedia.org/wiki/Descartes%27_rule_of_signs" target="_blank">Descartes&#39; rules of signs</a> before -- you&#39;d think this is impossible given the time I spent on those kinds of things in middle school. Anyways, it says:<br/>
  <br/>
  If the terms of a single-variable polynomial with real coefficients are ordered by descending variable exponent, then the number of positive roots of the polynomial is either equal to the number of sign differences between consecutive nonzero coefficients, or is less than it by an even number. Multiple roots of the same value are counted separately.<br/>
  <br/>
  Proof:<br/>
  <br/>
  By induction on the number of positive roots. All zero coefficients are ignored here. The base case is a polynomial \(f(x)\) without any positive root. Clearly the first and last cofficients must have the same sign, as otherwise \(f(x)\) has a positive root between \(0\) and \(\infty\). So the base case is resolved.<br/>
  <br/>
  Now, let \(g(x)=f(x)(x-a)\) with \(a&gt;0\). It is not hard to observe that the number of sign changes increases at least \(1\), with parity altered.<br/>
  Q.E.D.<br/>
  <div>
  <br/></div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2018/01/descartes-rules-of-signs.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
