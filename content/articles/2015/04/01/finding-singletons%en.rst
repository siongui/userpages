Finding singletons
##################

:date: 2017-09-18T16:44:00-07:00
:author: Shen-Fu Tsai
:tags: Math, Algorithm
:category: Algorithm
:summary: For \(i\in\{1,2,3\}\), there are \(2n+i\) integers consisting of \(n+i\) unique ones.
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  For \(i\in\{1,2,3\}\), there are \(2n+i\) integers consisting of \(n+i\) unique ones. \(n\) of them appear twice, and \(i\) of them just once. Find these singletons in \(O(n)\) time with \(O(1)\) space.<br/>
  <br/>
  Solution with \(O(\log n)\) space:<br/>
  Too bad, that finding median takes \(O(\log n)\) space! Let \(algorithm_i\) be the algorithm that tackles \(i\) singletons. \(algorithm_1\) is to just XOR all numbers.<br/>
  <br/>
  \(algorithm_2\): find median \(m\); if it&#39;s a singleton, apply \(algorithm_1\) to the rest; otherwise, count how many numbers are below and above \(m\). If both are odd, apply \(algorithm_1\) to each set separately, otherwise XOR all numbers in each set separately to find out which has two singletons, and then apply \(algorithm_2\) to that recursively.<br/>
  <br/>
  \(algorithm_3\): find median \(m\); if it&#39;s a singleton, apply \(algorithm_2\) to the rest; otherwise, count how many numbers are below and above \(m\) -- one set has even size and the other has odd size.XOR all numbers in the even-sized set to check if it contains singleton. If it does, apply \(algorithm_2\) to it and apply \(algorithm_1\) to the other; otherwise apply \(algorithm_3\) recursively to the other.<br/>
  <br/>
  <br/>
  Real solution:<br/>
  \(algorithm_2\): XOR all numbers to get \(s\). From \(s\) find a bit \(b\) that is \(1\) in \(s\). Now there must be odd numbers with bit \(b\) equal to \(0\) and odd numbers with bit \(b\) equal to \(1\). XOR each set separately to get these two singletons.<br/>
  <br/>
  \(algorithm_3\): If \(s\) is not zero, find a similar bit \(b\). There are odd numbers with \(b\) set to \(1\) and even numbers with \(b\) set to \(0\). Figure out if the second set has singleton by XOR, and act accordingly. If \(s=0\), find a bit \(b\) that not all numbers agree. There are even numbers with \(b\) set to \(1\) and odd numbers with \(b\) set to \(0\). Decide if the first set has two singletons by XOR and then proceed.<br/>
  <div>
  <br/></div>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/09/finding-singletons.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
