[Algorithm] Swap and Sum
########################

:date: 2015-01-15T13:04-08:00
:author: Shen-Fu Tsai
:tags: Algorithm
:category: Algorithm
:summary: Swap and Sum

There are two keys to solving the problem `swap and sum`_ :

1. Treap, which enables O(log N) time for swapping two segments and summing over
   one segment, because we could attach an addition field 'sum' to each node of
   Treap. However in this problem there is O(N) element swaps per request so the
   issue isn't resolved yet. 

2. Coordinate transform. This will deal with O(N) element swap. Transform the
   array into all even-indexed elements followed by all odd-indexed elements,
   i.e.  a[0], a[2], a[4], ... , a[1], a[3], ...

With this, any O(N) swap becomes swapping two segments of same length in the
transformed array, and every summation splits to two which is still managable.

----

`post <http://oathbystyx.blogspot.com/2015/01/swap-and-sum.html>`_
by
`Shen-Fu Tsai <{filename}/pages/sftsai.rst>`_


.. _swap and sum: https://www.hackerrank.com/contests/w13/challenges/swaps-and-sum
