[Golang] Coin sums - Problem 31 - Project Euler
###############################################

:date: 2018-10-07T22:28+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler
:category: Go
:summary: Go solution to Coin sums
          - Problem 31 - Project Euler.
:og_image: https://2.bp.blogspot.com/-jLTT8G2AFaw/VzS32ov82jI/AAAAAAAAB58/U9fMlRAtSvgZoOfzhfRE_VFSiHQ_AR4dgCLcB/s1600/project%2Beuler%2Bproblem%2B31%2Bwith%2Banswer.png
:adsu: yes

**Problem**: [1]_

  In England the currency is made up of pound, £, and pence, p, and there are
  eight coins in general circulation:

    1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).

  It is possible to make £2 in the following way:

    1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p

  How many different ways can £2 be made using any number of coins?


**Solution**:

  73682

  My solution is ugly, which uses the frute-force method.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/x1YKduEw-Ck>`__
   :class: align-center

.. code-block:: go

  package main
  
  import (
  	"fmt"
  )
  
  func main() {
  	numberOfWays := 0
  
  	// loop for £2 (200p)
  	for a := 0; a <= 1; a++ {
  		sumA := a * 200
  		if sumA >= 200 {
  			if sumA == 200 {
  				fmt.Println("£2 = £2 *", a)
  				numberOfWays++
  			}
  			break
  		}
  
  		// loop for £1 (100p)
  		for b := 0; b <= 2; b++ {
  			sumB := sumA + b*100
  			if sumB >= 200 {
  				if sumB == 200 {
  					fmt.Println("£2 = £2 *", a, "+ £1 *", b)
  					numberOfWays++
  				}
  				break
  			}
  
  			// loop for 50p
  			for c := 0; c <= 4; c++ {
  				sumC := sumB + c*50
  				if sumC >= 200 {
  					if sumC == 200 {
  						fmt.Println("£2 = £2 *", a, "+ £1 *", b, "+ 50p *", c)
  						numberOfWays++
  					}
  					break
  				}
  
  				// loop for 20p
  				for d := 0; d <= 10; d++ {
  					sumD := sumC + d*20
  					if sumD >= 200 {
  						if sumD == 200 {
  							fmt.Println("£2 = £2 *", a, "+ £1 *", b, "+ 50p *", c, "+ 20p *", d)
  							numberOfWays++
  						}
  						break
  					}
  
  					// loop for 10p
  					for e := 0; e <= 20; e++ {
  						sumE := sumD + e*10
  						if sumE >= 200 {
  							if sumE == 200 {
  								fmt.Println("£2 = £2 *", a, "+ £1 *", b, "+ 50p *", c, "+ 20p *", d, "+ 10p *", e)
  								numberOfWays++
  							}
  							break
  						}
  
  						// loop for 5p
  						for f := 0; f <= 40; f++ {
  							sumF := sumE + f*5
  							if sumF >= 200 {
  								if sumF == 200 {
  									fmt.Println("£2 = £2 *", a, "+ £1 *", b, "+ 50p *", c, "+ 20p *", d, "+ 10p *", e, "+ 5p *", f)
  									numberOfWays++
  								}
  								break
  							}
  
  							// loop for 2p
  							for g := 0; g <= 100; g++ {
  								sumG := sumF + g*2
  								if sumG >= 200 {
  									if sumG == 200 {
  										fmt.Println("£2 = £2 *", a, "+ £1 *", b, "+ 50p *", c, "+ 20p *", d, "+ 10p *", e, "+ 5p *", f, "+ 2p *", g)
  										numberOfWays++
  									}
  									break
  								}
  
  								// loop for 1p
  								for h := 0; h <= 200; h++ {
  									sumH := sumG + h*1
  									if sumH >= 200 {
  										if sumH == 200 {
  											fmt.Println("£2 = £2 *", a, "+ £1 *", b, "+ 50p *", c, "+ 20p *", d, "+ 10p *", e, "+ 5p *", f, "+ 2p *", g, "+ 1p *", h)
  											numberOfWays++
  										}
  										break
  									}
  								}
  							}
  						}
  					}
  				}
  			}
  		}
  	}
  
  	fmt.Println("number of ways:", numberOfWays)
  }

.. adsu:: 2

----

Test on:

- `Go Playground`_

References:

.. [1] `Coin sums - Problem 31 - Project Euler <https://projecteuler.net/problem=31>`_

.. _Go Playground: https://play.golang.org/
