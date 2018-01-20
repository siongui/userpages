APMO 2017 Problem 1
###################

:date: 2017-05-20T08:47:00-07:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary: We call a \(5\)-tuple of integers arrangeable if its elements can be labeled \(a, b, c, d, e\) in some order so that \(a-b+c-d+e=29\). Deter...
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  We call a \(5\)-tuple of integers arrangeable if its elements can be labeled \(a, b, c, d, e\) in some order so that \(a-b+c-d+e=29\). Determine all \(2017\)-tuples of integers \(n_1, n_2, \ldots , n_{2017}\) such that if we place them in a circle in clockwise order, then any \(5\)-tuple of numbers in consecutive positions on the circle is arrangeable.<br/>
  <br/>
  =============================<br/>
  <br/>
  <br/>
  Proof:<br/>
  <br/>
  My proof is very crooked. An obvious solution is all numbers being \(29\). If we can show regardless of the plus and minus signs the linear system is not underdetermined, then that will be the only solution. Given any qualifying linear system with matrix \(A\), we want to show \(det(A)\neq 0\). Since its parity doesn&#39;t change if we flip any plus or minus sign, it suffices to show \(det(A)\) is odd when the signs in each row are \(1, -1, 1, -1, 1\). We notice \(AB=C\), where \(A\) has two consecutive ones in each row, and \(C\) has \(1, 0, 0, 0, 0, 1\) each each row. Since we can show that \(det(B)=det(C)\), we have \(det(A)=1\).<br/>
  Q.E.D.<br/>
  <div>
  <br/></div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/05/apmo-2017-problem-1.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
