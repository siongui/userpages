[Math] International Mathematical Olympiad (IMO) 2014 Problem 6
###############################################################

:date: 2015-12-17T15:02-08:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary: `International Mathematical Olympiad`_ (IMO_) 2014 Problem 6


A set of lines in the plane is in general position if no two are parallel and no
three pass through the same point. A set of lines in general position cuts the
plane into regions, some of which have finite area; we call these its finite
regions. Prove that for all sufficiently large :math:`n`, in any set of
:math:`n` lines in general position it is possible to color at least
:math:`n^{0.5}` of the lines blue in such a way that none of its finite regions
has a completely blue boundary.

==================================================================

It looks like a problem where we prove asymptotic bound in computation theory,
but interestingly the statement holds for all :math:`n`.

==================================================================

Proof:

We describe an algorithm that colors at least :math:`n^{0.5}` lines blue while
keeping the condition satisfied.

Let :math:`S` be the set of lines that we couldn't color blue. As we proceed, we
pick a line not yet colored and not in :math:`S`. It is then colored blue, and
then we add non-blue lines touching regions that are therefore bounded by all
but one blue lines to :math:`S`. The algorithm halts as all lines are either
blue or in :math:`S`. We call these regions almost blue-bounded regions.

The key observation is that the second blue line adds at most 2 lines to
:math:`S`, the third 4, ..., and the :math:`(k-1)th` blue line adds at most
:math:`2(k-2)` lines to :math:`S`. So we could color the :math:`k-th` line blue
if and only if :math:`2[1+2+...+(k-2)] + (k-1) < n`, or :math:`k-1 < n^{0.5}`
Clearly the optimal :math:`k >= n^{0.5}`

The devil is in the detail of proving this observation.  After we color the
`k-th` line say, `L`, it intersects with other blue lines at `k-1` points, which
forms `k-2` segments. Each such segment `E` adds at most two lines to `S`:

(1) If :math:`L` doesn't intersect with any line inside :math:`E`, then at most
    the two regions touching :math:`E` become almost blue-bounded.
(2) If :math:`L` intersects with exactly one line :math:`M` inside :math:`E`,
    then at most :math:`M` is added to `S`.
(3) If :math:`L` intersects with lines :math:`M_1, M_2, ..., M_d` inside
    :math:`E`, then at most :math:`M_1` and :math:`M_d` get added to :math:`S`.

So these :math:`k-2` segments at most add :math:`2(k-2)` lines to :math:`S`.
Now, beyond the first or the :math:`(k-2)th` segment, if :math:`L` doesn't
intersect with any other line then we're done.  Otherwise, at most the line that
intersects with :math:`L` at a point closest to the segment gets added to
:math:`S`, so at most 2 more lines are added to :math:`S`, giving a total of
:math:`2(k-1)`.

Q.E.D.


----

`post <http://oathbystyx.blogspot.com/2015/12/imo-2014-problem-6.html>`_
by
`Shen-Fu Tsai <{filename}/pages/sftsai.rst>`_

----

References:

.. [1] `International Mathematical Olympiad <https://www.imo-official.org/>`__


.. _International Mathematical Olympiad: https://www.imo-official.org/
.. _IMO: https://www.imo-official.org/
