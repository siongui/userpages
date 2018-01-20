2009 USAMO Problem 3
####################

:date: 2017-04-25T22:23:00-07:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary: I found a different solution to the second part of this problem.
          We define a chessboard polygon to be a polygon whose edges are situate...
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  I found a different solution to the second part of this problem.<br/>
  <br/>
  <br/>
  <div style="text-align: justify;">
  We define a chessboard polygon to be a polygon whose edges are situated along lines of the form \(x = a\) or \(y = b\), where \(a\) and \(b\) are integers. These lines divide the interior into unit squares, which are shaded alternately grey and white so that adjacent squares have different colors. To tile a chessboard polygon by dominoes is to exactly cover the polygon by non-overlapping \(1\times 2\) rectangles. Finally, a tasteful tiling is one which avoids the two configurations of dominoes shown on the left below. Two tilings of a \(3\times 4\)<span style="text-align: center;"> rectangle are shown; the first one is tasteful, while the second is not, due to the vertical dominoes in the upper right corner.</span></div>
  <span style="text-align: center;"><br/></span>
  <br/>
  <div class="separator" style="clear: both; text-align: center;">
  <a href="https://2.bp.blogspot.com/-N5ooucX1YRQ/WQAtpc2vDXI/AAAAAAAAWmw/US44Vlh-A2M1p8Fx7GKDSQa_-YA8x0PvQCEw/s1600/Screen%2BShot%2B2017-04-25%2Bat%2B9.53.13%2BPM.png" imageanchor="1" style="margin-left: 1em; margin-right: 1em;"><img border="0" height="90" src="https://2.bp.blogspot.com/-N5ooucX1YRQ/WQAtpc2vDXI/AAAAAAAAWmw/US44Vlh-A2M1p8Fx7GKDSQa_-YA8x0PvQCEw/s400/Screen%2BShot%2B2017-04-25%2Bat%2B9.53.13%2BPM.png" width="400"/></a></div>
  <br/>
  <div style="text-align: justify;">
  a) Prove that if a chessboard polygon can be tiled by dominoes, then it can be done so tastefully.</div>
  b) Prove that such a tasteful tiling is unique.<br/>
  <br/>
  ===============================================<br/>
  <br/>
  Sol:<br/>
  <br/>
  a)<br/>
  <div style="text-align: justify;">
  There must be an upper right corner of the polygon \(P\). If it&#39;s white, cover it with a vertical domino and then tile the rest by induction. Otherwise cover it with a horizontal. It is clear that this domino doesn&#39;t form distasteful tiling with the rest.</div>
  <br/>
  b)<br/>
  <div style="text-align: justify;">
  Suppose there are two different tasteful tilings, then they form a chain surrounding a non-empty chessboard polygon \(R\) with a tasteful tiling because of induction. We show that walking counterclockwise along the perimeter of such surrounded \(R\) we can always find a domino with B following W, which we call bad domino, and making one of the tilings distasteful. This is proved by induction. Given any surrounded tasteful tiling with at least a bad domino, we will prove that adding any other domino does not decrease the number of bad dominoes. Suppose the bad domino is W above B somewhere locally rightmost in \(R\). The new domino covers either W or B to its right, but not both as that&#39;d be distasteful. However no matter what the new domino will be bad, which concludes the proof.</div>
  <div>
  <br/></div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/04/2009-usamo-problem-3.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
