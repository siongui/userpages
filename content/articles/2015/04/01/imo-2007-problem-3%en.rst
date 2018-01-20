IMO 2007 Problem 3
##################

:date: 2017-02-08T15:32:00-08:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary: In a mathematical competition some competitors are friends. Friendship is always mutual. Call a group of competitors a clique if each two of...
:og_image: https://lh4.googleusercontent.com/-aNB-12LGE20/AAAAAAAAAAI/AAAAAAAAZN0/9DWnaMSkb_M/photo.jpg

.. raw:: html

  In a mathematical competition some competitors are friends. Friendship is always mutual. Call a group of competitors a clique if each two of them are friends. The number of members of a clique is called its size. Given that, in this competition, the largest size of a clique is even, prove that the competitors can be arranged in two rooms such that the largest size of a clique contained in one room is the same as the largest size of a clique contained in the other room.<br/>
  <br/>
  =========================================<br/>
  <br/>
  Proof:<br/>
  <br/>
  Suppose not. Let graphs A and B be G and an empty graph, respectively. We move vertex one by one from A to B unless a move is invalid, meaning that after the move \(w(A)\)<w is="" left.="" move="" no="" p="" process="" terminates="" the="" there="" valid="" when=""><br/>
  Now we claim that w(A)=w(B)+1=d+1=|V(A)|. The first equalities are obvious. If there are more than d+1 vertices in A, then pick a maximum (d+1)-clique and move a vertex not in it to B.<br/>
  <br/>
  Move a vertex v from A to B. Now w(A)=w(B)-1=d=|V(A)|. How many (d+1)-cliques in B does v belong to?<br/>
  <br/>
  If v is in only one (d+1)-clique C, then by assumption moving every vertex in C back to A creates a (d+1)-clique in A, meaning that C, A, and v together is a clique of size 2d+1. Since the maximum clique in G has even size, it is at least of size 2d+2, a contradiction that the moving process above end up with maximum cliques of size d+1 and d in A and B, respectively.<br/>
  <br/>
  If v is in two (d+1)-cliques C1 and C2, we arbitrarily pick vertices a and b from C1-C2 and C2-C1, respectively. Moving a or b to A does not increase w(A), but moving both does. So a and b are adjacent. Because a and b are arbitrarily chosen, the union of C1 and C2 is a clique bigger than C1 and C2, a contradiction. Notice here we do not need A to be a d-clique, but only has to contain a d-clique.<br/>
  <br/>
  If v is in a set of k&gt;=3 (d+1)-cliques C, we start a process of moving vertex one by one in C to A. A vertex u is moved to A if it does not belong to all cliques in C, and all cliques in C that contain u is removed from the set C right away. By the previous argument C could not be left with two (d+1)-cliques. Nor can moving a vertex not in all cliques in C results in an empty C. The process terminates when moving any remaining vertex from C to A leaves only one clique in C.<br/>
  <br/>
  Now let&#39;s study C with k&gt;1 (d+1)-cliques. Every vertex, except for the vertices common to all cliques in C, belongs to exactly k-1 cliques in C, meaning that any two vertices belong to at least one common clique and are therefore adjacent. So all vertices in C form a larger clique, contradiction. This concludes our proof.<br/>
  <br/>
  Q.E.D.</w>

  <script src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

----

`post <https://oathbystyx.blogspot.com/2017/02/imo-2007-problem-3.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_
