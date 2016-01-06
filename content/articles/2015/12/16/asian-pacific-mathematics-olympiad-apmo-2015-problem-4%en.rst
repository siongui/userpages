[Math] Asian Pacific Mathematics Olympiad (APMO) 2015 Problem 4
###############################################################

:date: 2015-12-16T16:47-08:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary: `Asian Pacific Mathematics Olympiad`_ (APMO_) 2015_ Problem 4


Given :math:`n` blue lines and :math:`n` red lines in the plane such that no two
lines are parallel to each other, show that there is a circle that intersects
with the set of blue lines in exactly :math:`2n-1` points, and intersects with
the set of red lines in exactly :math:`2n-1` points.


Proof:
``````

We will prove a more specific condition, that there exists a circle such that
for each color, is tangential to one line and intersects with the rest at two
points each. The intuition is to find a red and blue line that "bundle" the rest
of :math:`2n-2` lines, then a circle big enough and tangential to both this red
and blue line indeed intersects the rest of :math:`2n-2` lines at two points
each.

To find such two lines, it suffices to take the blue and red lines that span the
largest angle between them. Even these two lines may not be unique, it is not
hard to show that a sufficiently big circle tangential to them intersects each
of other lines at two points.

Q.E.D.

----

`post <http://oathbystyx.blogspot.com/2015/12/apmo-2015-problem-4.html>`_
by
`Shen-Fu Tsai <{filename}/pages/sftsai.rst>`_

----

References:

.. [1] `Asian Pacific Mathematics Olympiad <https://cms.math.ca/Competitions/APMO/>`_


.. _Asian Pacific Mathematics Olympiad: https://cms.math.ca/Competitions/APMO/
.. _APMO: https://cms.math.ca/Competitions/APMO/
.. _2015: https://cms.math.ca/Competitions/APMO/exam/apmo2015.pdf
