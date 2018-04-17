Generate Fair Results From A Biased Coin
########################################

:date: 2018-04-17T23:27+08:00
:tags: Go, Golang, Random Number, Go time Package, Algorithm
:category: Algorithm
:summary: Given a biased coin with the probability of p to be head on each toss,
          where 0 < p < 1 and p ≠ 0.5, generate fair results from the biased
          coin.
:og_image: http://www.learner.org/courses/learningmath/data/images/session8/8c1.gif
:adsu: yes


Also saw from *Techie Delight* [1]_

  *Given a biased coin with the probability of p to be head on each toss, where
  0 < p < 1 and p ≠ 0.5, generate fair results from the biased coin.*

Wikipedia [2]_ also explains the solution to this problem. The following texts
comes from wiki:

  1. Toss the coin twice.
  2. If the results match, start over, forgetting both results.
  3. If the results differ, use the first result, forgetting the second.

This works because if we toss twice, the probability of HEAD-TAIL and TAIL-HEAD
are both *p(1-p)*. So we can ignore the case of HEAD-HEAD and TAIL-TAIL and use
only HEAD-TAIL and TAIL-HEAD, which are fair results coming from the biased
coin.

The following solution is implemented in Go.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/ITO5XAZu0pX>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"math/rand"
  	"time"
  )

  var r *rand.Rand // Rand for this package.

  type RESULT int

  const (
  	HEAD RESULT = iota
  	TAIL
  )

  func BiasedCoin() RESULT {
  	p := 30 // p = 0.3 to be HEAD on each toss
  	if r.Intn(100) < p {
  		return HEAD
  	} else {
  		return TAIL
  	}
  }

  func Fair() RESULT {
  	for {
  		r1 := BiasedCoin()
  		r2 := BiasedCoin()
  		if r1 != r2 {
  			return r1
  		}
  	}
  }

  func init() {
  	r = rand.New(rand.NewSource(time.Now().UnixNano()))
  }

  func main() {
  	m := make(map[RESULT]int)
  	m[HEAD] = 0
  	m[TAIL] = 0

  	for i := 0; i < 100000; i++ {
  		m[Fair()] += 1
  	}

  	for r, count := range m {
  		print(r)
  		print(" : ")
  		println(count)
  	}
  }


One of the outputs run on my desktop:

.. code-block:: txt

  0 : 50021
  1 : 49979

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10.1``.

References:

.. [1] `Generate Fair Results from a Biased Coin - Techie Delight <http://www.techiedelight.com/generate-fair-results-biased-coin/>`_
.. [2] `Fair results from a biased coin - Wikipedia <https://en.wikipedia.org/wiki/Fair_coin#Fair_results_from_a_biased_coin>`_
.. [3] `Generate Any One of Given Numbers According to Given Probabilities <{filename}/articles/2018/04/16/generate-any-one-of-given-numbers-according-to-given-probabilities%en.rst>`_

.. _Go Playground: https://play.golang.org/
