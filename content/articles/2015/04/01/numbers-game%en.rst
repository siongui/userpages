Numbers game
############

:date: 2017-09-06T11:56:00-07:00
:author: Shen-Fu Tsai
:tags: Math, Algorithm
:category: Math
:summary: Given a sequence of \(n\) integers \(a_1,\ldots,a_n\), we map it to ...
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  Given a sequence of \(n\) integers \(a_1,\ldots,a_n\), we map it to \(|a_1-a_2|,\ldots,|a_{n-1}-a_n|,|a_n-a_1|\). We repeat this process again and again until all numbers become zero. For what \(n\) is this process guaranteed to stop?<br/>
  <br/>
  Solution:<br/>
  <br/>
  Interestingly parity (not me!) plays a central role here. It turns out all sequences terminate if and only if all \(0-1\) sequences do. Why? If all \(0-1\) sequences terminate then any integer sequence eventually becomes an all-even sequence, and it doesn&#39;t harm to divide every number by \(2\), and the process repeats. Note that the maximum number in the sequence never goes up, and goes down by half each time they are divided by \(2\), so this process clearly will not go on forever than hence must stop.<br/>
  <br/>
  So for what \(n\) could all \(0-1\) sequences terminate? It is not hard to show that if \(k\) qualifies, then \(2k\) does too, so all powers of two qualify. Clearly an odd \(n\) does not qualify, because unless initially all numbers in the sequence are equal, the second to the last step is obtain the alternating sequence \(0,1,\ldots,0,1\), which doesn&#39;t exist for odd \(n\).<br/>
  <br/>
  What about an even \(n\) not a power of two? Here comes the most fun part. If \(n=(2k+1)2^m\), divide the sequence into \(2k+1\) blocks, each of length \(2^m\). Let each block starts with either \((0,\ldots,0)\) or \((1,0,\ldots,0)\). With the same argument used to show that \(2k\) qualifies if \(k\) does, we can show that after a cycle a block becomes \((1,0,\ldots,0)\) if either itself or the subsequent block is \((1,0,\ldots,0)\), and \((0,\ldots,0)\) if it is identical to the subsequent block. So, its behavior is identical to \(n=2k+1\) if \((1,0,\ldots,0)\) is considered \(1\) and \((0,\ldots,0)\) is \(0\). Since an odd \(n\) does not qualify, neither does \(n=(2k+1)2^m\).<br/>
  <div>
  <br/></div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/09/numbers-game.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
