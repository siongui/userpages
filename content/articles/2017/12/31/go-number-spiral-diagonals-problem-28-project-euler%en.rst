[Golang] Number spiral diagonals - Problem 28 - Project Euler
#############################################################

:date: 2018-10-24T01:51+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler
:category: Go
:summary: Go solution to Number spiral diagonals
          - Problem 28 - Project Euler.
:og_image: https://1.bp.blogspot.com/-70_M0zB9n8A/Vyt9HwCT-0I/AAAAAAAAB3s/ur6VXgOxz3QoISbXdAoDsZDoOuIMqgs-QCLcB/s1600/project%2Beuler%2Bproblem%2B28%2Bwith%2Banswer.png
:adsu: yes

**Problem**: [1]_

.. raw:: html

  <p>Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:</p>
  <p style="text-align:center;font-family:'courier new';"><span style="color:#ff0000;font-family:'courier new';"><b>21</b></span> 22 23 24 <span style="color:#ff0000;font-family:'courier new';"><b>25</b></span><br />
  20  <span style="color:#ff0000;font-family:'courier new';"><b>7</b></span>  8  <span style="color:#ff0000;font-family:'courier new';"><b>9</b></span> 10<br />
  19  6  <span style="color:#ff0000;font-family:'courier new';"><b>1</b></span>  2 11<br />
  18  <span style="color:#ff0000;font-family:'courier new';"><b>5</b></span>  4  <span style="color:#ff0000;font-family:'courier new';"><b>3</b></span> 12<br /><span style="color:#ff0000;font-family:'courier new';"><b>17</b></span> 16 15 14 <span style="color:#ff0000;font-family:'courier new';"><b>13</b></span></p>
  <p>It can be verified that the sum of the numbers on the diagonals is 101.</p>
  <p>What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?</p>

**Solution**:

  669171001

  First take a look at the four numbers on the diagonals of N x N square.
  for example, *3 5 7 9* on **3 x 3** square. The difference is 2 on the
  numbers. So the difference is *n-1* on **N x N** square. We use this
  observation to calculate all numbers on the diagonals and hence the sum of
  them.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/-yKs-IMUoTy>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func main() {
  	sum := 1
  	nextNumberInDiagoal := 1
  	for j := 2; j < 1001; j += 2 {
  		for i := 0; i < 4; i++ {
  			nextNumberInDiagoal += j
  			sum += nextNumberInDiagoal
  			//fmt.Println(nextNumberInDiagoal)
  		}
  	}
  	fmt.Println(sum)
  }

.. adsu:: 2

----

Test on:

- `Go Playground`_

References:

.. [1] `Number spiral diagonals - Problem 28 - Project Euler <https://projecteuler.net/problem=28>`_
.. [2] `Two-dimensional slices - Effective Go <https://golang.org/doc/effective_go.html#two_dimensional_slices>`_

.. _Go Playground: https://play.golang.org/
