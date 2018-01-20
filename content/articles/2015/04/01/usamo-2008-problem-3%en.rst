USAMO 2008 Problem 3
####################

:date: 2017-04-20T18:31:00-07:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary: The problem is rephrased.
          Q: A diamond of size n consists of 2n rows of nodes with lengths from top to bottom being 1, 3, 5, ..., 2n-1...
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  The problem is rephrased.<br/>
  <div class="navbar section" id="navbar" style="-webkit-text-stroke-width: 0px; color: black; font-family: Times; font-size: medium; font-style: normal; font-variant-caps: normal; font-variant-ligatures: normal; font-weight: normal; letter-spacing: normal; orphans: 2; text-align: start; text-decoration-color: initial; text-decoration-style: initial; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px;">
  <div class="post-body entry-content">
  <div style="margin: 0px;">
  <br/></div>
  <div style="margin: 0px;">
  Q: A diamond of size n consists of 2n rows of nodes with lengths from top to bottom being 1, 3, 5, ..., 2n-1, 2n-1, 2n-3, ..., 3, 1. The middle nodes of all rows are in the same column. For instance the following &#39;+&#39; show a diamond of size 3:</div>
  <div style="margin: 0px;">
  <br/></div>
  <div style="margin: 0px;">
  ##+##</div>
  <div style="margin: 0px;">
  #+++#</div>
  <div style="margin: 0px;">
  +++++</div>
  <div style="margin: 0px;">
  +++++</div>
  <div style="margin: 0px;">
  #+++#</div>
  <div style="margin: 0px;">
  ##+##</div>
  <div style="margin: 0px;">
  <br/></div>
  <div style="margin: 0px;">
  We want to partition each diamond into acyclic paths such that each node is contained in exactly one path. A path is a sequence of distinct nodes in which neighboring nodes are horizontally or vertically adjacent. Prove that for diamond of size n the number of paths is at least n.</div>
  <div style="margin: 0px;">
  <br/></div>
  <div style="margin: 0px;">
  Proof:</div>
  <div style="margin: 0px;">
  We&#39;ll prove by induction. The case for n=1 is trivial. Now as an example consider a diamond of size 4. We draw part of it below where &#39;X&#39; denote nodes of diamond we don&#39;t quite care:</div>
  <div style="margin: 0px;">
  <br/></div>
  <div style="margin: 0px;">
  A</div>
  <div style="margin: 0px;">
  BC</div>
  <div style="margin: 0px;">
  XDE</div>
  <div style="margin: 0px;">
  XXFG</div>
  <div style="margin: 0px;">
  XXIH</div>
  <div style="margin: 0px;">
  XLJ</div>
  <div style="margin: 0px;">
  NP</div>
  <div style="margin: 0px;">
  Q</div>
  <div style="margin: 0px;">
  <br/></div>
  <div style="margin: 0px;">
  A useful argument we name OA (optimality argument) is that if a node has only one neighbor, connecting it to its neighbor doesn&#39;t affect optimality. Suppose x has only one neighbor y and is not connected to y. If in some optimal partitioning y is an endpoint of some path, certainly adding x to that path reduces number of paths; if y is in the middle of some path, splitting that path at y and connecting x with y gives the same number of paths.</div>
  <div style="margin: 0px;">
  <br/></div>
  <div style="margin: 0px;">
  With OA, we connect AB. If C is connected with B then regard ABC as a super node C&#39; with only one neighbor D; if C is not connected with B then C has only one neighbor D. In either case we are allowed to connect C with D using OA. Similarly, either E or super node CDE or super node ABCDE has only one neighbor F and thus we connect E with F. By symmetry we connect N with Q, L with P, I with J.</div>
  <div style="margin: 0px;">
  <br/></div>
  <div style="margin: 0px;">
  Next we argue that at least one pair of FG and IH can be connected. If not, G and H should be connected for optimality and is regarded as a super node; if EFIJ is a sub-path then EFGHIJ reduces number of paths by 1; if EFIJ is not a sub-path then we can split EFX... and attach GH to F without increasing number of paths. In either case, connecting FG or IH does not affect optimality.</div>
  <div style="margin: 0px;">
  <br/></div>
  <div style="margin: 0px;">
  Suppose FG is connected. Because from G we can&#39;t go beyond Q and never reach D, we connect DE and hence BC. Now ABCDEFG is a super node with one neighbor H so we connect GH, IH, JL, NP in order. This leaves us a path that can&#39;t be further extended and the rest is a diamond of size 3.</div>
  <div style="margin: 0px;">
  <br/></div>
  <div style="margin: 0px;">
  Q. E. D.</div>
  </div>
  </div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/04/usamo-2008-problem-3.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
