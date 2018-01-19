Caratheodory's theorem
######################

:date: 2017-12-22T21:09:00-08:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary: Previous I used Caratheodory's theorem to prove Radon's theorem, now
          I'm going to prove the former.
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  Previous I used Caratheodory&#39;s theorem to prove Radon&#39;s theorem, now I&#39;m going to prove the former.<br/>
  <br/>
  Caratheodory&#39;s theorem: if a point \(x\) of \(R^d\) lies in the convex hull of a set \(P=\{p_1,\ldots,p_{n}\}\), then \(x\) can be written as the convex combination of at most \(d+1\) points in \(P\).<br/>
  <br/>
  Proof:<br/>
  It suffices to prove it for \(|P|=d+2\), and then for \(|P|=d+3\) we can reduce the convex combination of \(p_1,\ldots,p_{d+2}\) to that of some \(d+1\) points in \(\{p_1,\ldots,p_{d+2}\}\) which together with \(p_{d+3}\) can be reduced again to convex combination of some \(d+1\) points, and so on.<br/>
  <br/>
  Let \(P&#39;=\{p_1,\ldots,p_{d+1}\}\). Given \(x\in conv(P)\), if \(x\in conv(P&#39;)\) then we are done. If not, note that it is nevertheless the convex combination of \(p_{d+2}\) and \(z\in conv(P&#39;)\). So \(x\) belongs to segment \(zp_{d+2}\). Consider point y, one of the intersections of \(zp_{d+2}\) and boundary of \(conv(P&#39;)\). Point \(x\) must lie between \(y\) and \(p_{d+2}\), so it is the convex combination of the two as well. As \(y\) is at the boundary of \(conv(P&#39;)\), it could be represented as convex combination of \(P&#39;\) with at least one zero coefficient, i.e. at most \(d\) points with positive coefficient. So \(x\) can be written as convex combination of \(P\) with at most \(d+1\) positive coefficients.<br/>
  Q.E.D.<br/>
  <div>
  <br/></div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/12/caratheodorys-theorem.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
