Robot
#####

:date: 2014-09-23T18:34-08:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary: Robot


In this post we present a coding problem and our solution which is different
from the author's. Our solution doesn't involve tree, and only a stack is used.

https://www.hackerrank.com/contests/w10/challenges/robot

The original problem was phrased in psuedo code. In my wording it becomes:

A robot is to travel from station 0 to station 1, then to station 2, station 3,
and so on, to station N - 1. It enters station 0 with energy 0 and no money. At
each station i, the robot has two options: either it resets its energy to
:math:`P_i \geq 0`, or it takes money :math:`V_i \geq 0`. As it enters next
station, its energy deceases by 1. The only restriction is that its energy
should never becomes negative. The goal is to maximize the total amount of money
gathered. :math:`1 \leq N \leq 500000`, :math:`0 \leq V_i \leq 10^9`,
:math:`0 \leq P_i \leq 100000`. It is guaranteed there's a way that energy is
always non-negative.

A standard DP with state (station, energy) takes O[N * max_energy]. Let D(i, e)
be the maximum amount of money to be gained when entering station i with energy
e. Clearly


  D(i, e)

  = max{

    max money gained if reset energy at station i,

    max money gained if take :math:`V_i` at station i

  }

  = max{

    :math:`P_i` == 0 ? -1 : D(i + 1, :math:`P_i` - 1),

    e == 0 ? -1 : :math:`V_i` + D(i + 1, e - 1)

  }


As the complexity is too high, we have to try reducing it. Let's observe an
example and see what we got.

N = 5

P: 3 0 1 3 1
V: 2 1 3 6 9

===========================================================================

Our DP starts from i = 4. Because there's no concern of energy after entering
the last station, D(4, e) are

9 9 9 9 9 9 9....

===========================================================================

For i = 3, either we take energy 3 and enter i = 4 with energy 2 which will end
up with 9 dollars, or we take 6 dollars if we enter with positive energy and end
up with 6 + 9 = 15 dollars. So we're comparing these two rows:

  9  9  9  9  9  9  9 ....

    15 15 15 15 15 15 ....

which becomes

  9 15 15 15 15 15 15 ...

===========================================================================

For i = 2, either we take energy 1 and enter i = 3 with energy 0 which will end
up with 9 dollars, or we take 3 dollars if we enter with positive energy and end
up with 3 more dollars than i=3. So we're comparing these two rows:

  9  9  9  9  9  9  9 ...

    12 18 18 18 18 18 ...

which becomes

  9 12 18 18 18 18 18 ...

===========================================================================

As we see now, the pattern of update is to compare 2 rows. The first is obtained
by expanding D(i + 1, :math:`P_i` - 1) to a row. The second is by right shifting
the row of D(i + 1, e) by 1 and add :math:`V_i`. Adding :math:`V_i` is expensive
because the row can be long, so why not store
:math:`V_i + V_{i + 1} +...+ V_{N - 1}` - earning in the table? This is exactly
the amount of money we don't manage to get from station i to the end. This is
the first observation.

The second observation is that the number of distinct values increases by at
most 1, and is bounded by :math:`N - i`. So what if we just store a vector or
stack of segments with size no more than :math:`O(N)`? One issue follows
immediately: because of the right shift we have to update the end points of
these segment in :math:`O(N)` for each i. But since the right shift always
occurs, what about storing end point - :math:`(N - i)` as the coordinate of the
segment? If a segment is not "interrupted", this value remains constant across
stations and needs no update. In this way no :math:`O(N)` update is necessary.
In each iteration, the main task is to find the value of
D(i + 1, :math:`P_i` - 1) because it's not immediately available -- we have to
perform :math:`O(log(N))` binary search of :math:`P_i - 1` on the end points of
the segment. Once found, updating the stack takes amortized cost of
:math:`O(1)`.

In brief, with these two coordinate transformations this problem is solved in
:math:`O(N*log(N))` time and :math:`O(N)` memory without tree structure. The
code is available at
https://github.com/paritystsai8/coding_problem/blob/master/robot.cpp

----

`post <http://oathbystyx.blogspot.com/2014/09/robot.html>`_
by
`Shen-Fu Tsai <{filename}/pages/sftsai.rst>`_
