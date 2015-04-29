[Algorithm] Insane DFS
######################

:date: 2015-01-15T11:07-08:00
:author: Shen-Fu Tsai
:tags: Algorithm
:category: Algorithm
:summary: Insane DFS


The `hardest problem`_ in the latest weekly HackerRank challenge looks like a
problem of graph theory, but only at the superficial level. It is not hard to
tell that the "suitable array" defined by DFS traversal throughout a rooted tree
is equivalent to all and only arrays a[] satisfying these constraints:

1. a[0] = 0, a[1] = 1 if N > 1
2. For n > 0, a[n] :math:`\leq` a[n - 1] + 1

Let's start with DP. Let D[x][y] be the number of arrays a[0 ... x] with a[x]
:math:`\geq` y. Then, the number of arrays a[0 ... x] such that a[x] == y is
exactly D[x - 1][y - 1] because a[x - 1] must be no less than y - 1. Thus we
have relation

D[x][y] = D[x - 1][y - 1] + D[x][y + 1] (*)

with some boundary conditions, including D[x][0] = D[x][1]. For simplicity we
can append a[N] = 1 in order to find the final answer. Suppose a[:math:`x_0`]
and a[:math:`x_1`] are the only fixed numbers in [:math:`x_0`, :math:`x_1`], the
core task is to compute f(:math:`x_0`, :math:`y_0`, :math:`x_1` - 1, :math:`y_1`
- 1) = D[:math:`x_1` - 1][a[:math:`x_1`] - 1] from D[:math:`x_0`][0] =
D[:math:`x_0`][1] = ... = D[:math:`x_0`][a[:math:`x_0`]] = 1, and the final
answer will be the product of all such f(:math:`x_0, y_0, x_1, y_1`) modulo M.

With initial observations, I almost felt f(:math:`x_0, y_0, x_1, y_1`) has
binomial form C(a, b), although it doesn't because of boundary condition D[x][0]
= D[x][1]. After some calculations, I noticed the nice relation

f(:math:`x_0, y_0, x_1, y_1`) = C(2W - H, W) - C(2W - H, W - :math:`y_1` - 1)

where W = :math:`x_1 - x_0`, H = :math:`y_1 - y_0`. The proof is actually
mathematically very interesting.

Imagine you're on an infinite xy-plane and stand at cell (0, :math:`y_0`).
Relation (*) says you can only go in two directions (1, 1) or (0, -1). Denote
them by step A and B. Then D[:math:`x_1`][:math:`y_1`] is the number of shortest
path you may take to get to cell (:math:`x_1 - x_0, y_0`), and the constraint
D[x][0] = D[x][1] implies you are not allowed to go below x-axis. Now,
C(2W - H, W) is the number of shortest paths you could take without this
contraint. This is because to go from (0, :math:`y_0`) to
(:math:`x_1 - x_0, y_1`), the number of A step you need is exactly
:math:`x_1 - x_0` = W since only A step makes x-direction move. The number of
required B step to count for y-direction move is hence
:math:`y_0 - y_1` + W = W - H. We could also transform the xy-plane to
x'y'-plane with x' = x and y' = y - x. There A step become horizontal step
(1, 0) and B step is vertical downward (0, -1), i.e. everything is square and
not tilted.

We now have to show C(2W - H, W - :math:`y_1` - 1) is number of shortest paths
that goes under x-axis in xy-plane. Going under x-axis is the same as stepping
into zone x' + y' < 0 since y = x + y' = x' + y'. What is the number of such
shortest paths? For illustration we draw an example of :math:`y_0` = 2,
W - H = 6, :math:`x_1 - x_0` = 13 in x'y'-plane below, where the goal is to go
from upper left corner to lower right corner by stepping in at least one '*'
cell, i.e. all '*' cells have x' + y' < 0.

::

  (0, 2)
        oooooooooooooo
        oooooooooooooo
  (0, 0)oooooooooooooo
        *ooooooooooooo
        **oooooooooooo
        ***ooooooooooo
        ****oooooooooo (13, -4)

Here's the magic: we can 'fold' the plane along x' + y' = 0 to get this:

::

  (0, 2)
        oooooooooooooo
        oooooooooooooo
  (0, 0)oooooooooooooo
        *ooo----------
        **oo----------
        ***o----------
        ****---------- (13, -4)
        ++++
        ++++
        ++++
        ++++
        ++++
        ++++
        ++++
        ++++
        ++++
        ++++ (3, -14)

Fold all the '-' cells to the '+' cells shown above.  There's then a 1-1 mapping
between the original shortest paths from (0, 2) to (13, -4) touching at least
one '*' cell and the shortest paths from (0, 2) to (3, -14), and so the number
of shortest paths is C(19, 3). To generalize, with some counting we can thus
prove that it's C(2W - H, W - H + 1 - (:math:`y_0` + 1) - 1) =
C(2W - H, W - :math:`y_1` - 1)

Overall, this solution takes O(N) computation of binomial coefficient C(a, b).
Assuming each C(a, b) requires O(b)log(M) where M is :math:`10^9 + 7` and
b = O(W), each boils down to O(W)log(M). The total time complexity is then
O(N)log(M) and space complexity O(1).

----

`post <http://oathbystyx.blogspot.com/2015/01/insane-dfs.html>`_
by
`Shen-Fu Tsai <{filename}/pages/sftsai.rst>`_


.. _hardest problem: https://www.hackerrank.com/contests/w13/challenges/insane-dfs
