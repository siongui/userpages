No crossing!
############

:date: 2017-09-05T16:04:00-07:00
:author: Shen-Fu Tsai
:tags: Math, Algorithm
:category: Math
:summary: On the 2D plane there are \(n\) blue and \(n\) red points, no three of
          them are co-linear. Then we can always pair a blue point with a red ...
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  On the 2D plane there are \(n\) blue and \(n\) red points, no three of them are co-linear. Then we can always pair a blue point with a red one and draw \(n\) segments between them, one from each pair, such that no two segments cross with each other.<br/>
  <br/>
  My proof:<br/>
  Induction is handy here. If the convex hull has points of both colors, we could isolate a red and an adjacent blue one with the rest, and induction will work. Otherwise, say all points on the convex hull are red and we shoot rays from a fixed red points \(u\) on the convex hull. Each ray divides other \(2n-1\) points into two groups, and we count the number of blue points minus the number of red points in each group. At some point these two numbers are \((-1,2)\) and at the other \((2,-1)\). Then one of the rays has one of the numbers \(0\), and the induction will work again.<br/>
  <br/>
  Another proof:<br/>
  When there is intersection, we can always replace the two intersecting segments by another two that don&#39;t intersect and reduce the total lengths of segment. The process cannot go on infinitely because the total length is bounded from below. Therefore there will be a time when we could not do this replacement anymore, i.e. no segment intersect with another.<br/>
  <div>
  <br/></div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/09/no-crossing_5.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
