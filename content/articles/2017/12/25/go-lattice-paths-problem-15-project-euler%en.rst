[Golang] Lattice paths - Problem 15 - Project Euler
###################################################

:date: 2017-12-25T03:49+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler
:category: Go
:summary: Go solution to Lattice paths
          - Problem 15 - Project Euler.
:og_image: https://projecteuler.net/project/images/p015.gif
:adsu: yes


**Problem**: [1]_

  Starting in the top left corner of a 2×2 grid, and only being able to move to
  the right and down, there are exactly 6 routes to the bottom right corner.

  .. image:: https://projecteuler.net/project/images/p015.gif
     :align: center
     :alt: Lattice paths

  How many such routes are there through a 20×20 grid?

**Solution**:

  Observe the case of a 2x2 grid. We can found that to move from top left to
  bottom right, we need to move 2 down steps and 2 right steps. The total number
  of routes in 2x2 grid is hence equivalent to total number of combinations of 2
  down steps and 2 right steps in a row, i.e., :math:`\binom{4}{2}`
  = :math:`\frac{4!}{2! 2!}` = 6.

  So, the routes in 20x20 grid = :math:`\binom{20}{10}` =
  :math:`\frac{20!}{10! 10!}` = 184756

.. adsu:: 2

----

References:

.. [1] `Lattice paths - Problem 15 - Project Euler <https://projecteuler.net/problem=15>`_
