IMO 2016 Problem 6
##################

:date: 2017-01-02T10:46:00-08:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary: There are n>=2 line segments in the plane such that every two segments cross and no three segments meet at a point. Geoff has to choose an ...
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  There are n&gt;=2 line segments in the plane such that every two segments cross and no three segments meet at a point. Geoff has to choose an endpoint of each segment and place a frog on it facing the other endpoint. Then he will clap his hands n-1 times. Every time he claps, each frog will immediately jump forward to the next intersection point on its segment. Frogs never change the direction of their jumps. Geoff wishes to place the frogs in such a way that no two of them will every occupy the same intersection point at the same time.<br/>
  <br/>
  (a) Prove that Geoff can always fulfill his wish if n is odd.<br/>
  <br/>
  (b) Prove that Geoff can never fulfill his wish if n is even.<br/>
  <br/>
  ===================================================================<br/>
  <br/>
  Before solving (b), I got the slight hint words &#34;clockwise&#34;, and a bigger hint to draw a huge circle that contains all intersections and label the segments in the order of their intersection with the circle. However it still took me a really long time to got the proof, before which I tried quite a few useless directions, including a false conjecture that there is always an intersection which ranks exactly in the middle on both segments.<br/>
  <br/>
  Proof:<br/>
  <br/>
  (a)<br/>
  <br/>
  When n is odd, the number of intersections on either side of any intersection on the same segment have different parity. Below, we will define f(end point) as 1 if a frog starts at the end point, and 0 otherwise. We also denote the number of intersections between two intersections A and B on the same segment by c(AB). <br/>
  <br/>
  It turns out we could assign the frogs in such a way that the numberings of any intersection for the two segments have different parity, i.e. a stronger statement than the required.<br/>
  <br/>
  Without loss of generality, let f(A)=1 in the figure. We then decide the locations of all other n-1 frogs accordingly, and we have to prove that there won&#39;t be any conflict. Take any two other segments CD and EF, as shown in the figure. Based on how we assign f(), we have (all additions below are modulo 2):<br/>
  <br/>
  <br/>
  <div class="separator" style="clear: both; text-align: center;">
  </div>
  <div class="separator" style="clear: both; text-align: center;">
  <a href="https://4.bp.blogspot.com/-dbT8LcAOtq8/WGqfnQxRQwI/AAAAAAAAVLs/D7b5I_tT6WY42XTJgRE2ZaZj-5qCUFApQCLcB/s1600/imo%2B2016%2Bp6-a%2B%25281%2529.png" imageanchor="1" style="margin-left: 1em; margin-right: 1em;"><img border="0" height="300" src="https://4.bp.blogspot.com/-dbT8LcAOtq8/WGqfnQxRQwI/AAAAAAAAVLs/D7b5I_tT6WY42XTJgRE2ZaZj-5qCUFApQCLcB/s400/imo%2B2016%2Bp6-a%2B%25281%2529.png" width="400"/></a></div>
  <br/>
  f(C)=c(AG)+c(CG)<br/>
  f(E)=c(AG)+c(GH)+1+c(EH)<br/>
  <br/>
  or<br/>
  <br/>
  f(C)+f(E)=c(CG)+c(GH)+c(EH)+1<br/>
  <br/>
  It suffices to show that there&#39;s no conflict between segments CD and EF, for which we need:<br/>
  <br/>
  f(C)+f(E)=c(CG)+1+c(GI)+c(EH)+1+c(HI)+1=1+c(CG)+c(GI)+c(EH)+c(HI)<br/>
  <br/>
  So all we need to prove is<br/>
  <br/>
  c(GH)+c(GI)+c(HI)=0<br/>
  <br/>
  which is trivial.<br/>
  <br/>
  (b)<br/>
  We number the segments clockwise S_{1}, S_{2}, ..., S_{n}. Consider a huge circle that contains all intersections and intersects with the segments at P_{1}, P_{2}, ... P_{2n}. The very crucial observation is that the intersection of every two consecutive segments always ranks exactly the same when counting from consecutive intersections with the circle! So, if the frog starts from P_{1}, then its counterparts start from P_{3}, P_{5}, ... P_{n-1}, P_{n+1}, .... However P_{n+1} and P_{1} are both end points of S_{1}, contradiction.

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/01/imo-2016-problem-6-a.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
