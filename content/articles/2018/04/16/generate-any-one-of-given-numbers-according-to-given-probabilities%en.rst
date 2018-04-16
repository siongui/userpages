Generate Any One of Given Numbers According to Given Probabilities
##################################################################

:date: 2018-04-16T23:49+08:00
:tags: Go, Golang, Random Number, Go time Package, Algorithm
:category: Algorithm
:summary: Given N numbers, generate any one of the given numbers according to
          given probabilities.
:og_image: http://codetheory.in/wp-content/uploads/2013/04/random_post_img.png
:adsu: yes


I saw from *Techie Delight* [1]_

*Write an algorithm to generate any one of the given N numbers according to
given probabilities*

For example,

| input        {1,  2,  3,  4,  5}
| probability  {30, 10, 20, 15, 25}  // sum should be 100

The solution should return *1* with *30%* probability, *2* with *10%*, and so
on.

See [1]_ for algorithm details. The following is the algorithm implemented in
Go.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/EuWDstu59sw>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"math/rand"
  	"time"
  )

  var r *rand.Rand // Rand for this package.

  func Random(input []int, prob []int) int {
  	// check sanity
  	if len(input) != len(prob) {
  		panic("length of input and prob is not equal")
  	}
  	sum := 0
  	for _, p := range prob {
  		sum += p
  	}
  	if sum != 100 {
  		panic("sum of prob must be 100")
  	}

  	probSum := make([]int, len(input))
  	for i, p := range prob {
  		if i == 0 {
  			probSum[0] = p
  		} else {
  			probSum[i] = probSum[i-1] + p
  		}
  	}

  	ri := r.Intn(100)

  	for i, ps := range probSum {
  		if ri < ps {
  			return input[i]
  		}
  	}

  	panic("should not run here")
  }

  func init() {
  	r = rand.New(rand.NewSource(time.Now().UnixNano()))
  }

  func main() {
  	input := []int{1, 2, 3, 4, 5}
  	prob := []int{30, 10, 20, 15, 25}

  	m := make(map[int]int)
  	for _, n := range input {
  		m[n] = 0
  	}
  	for i := 0; i < 1000000; i++ {
  		m[Random(input, prob)] += 1
  	}
  	for k, v := range m {
  		print(k)
  		print(" : ")
  		println(v)
  	}
  }

.. adsu:: 2

One of the outputs run on my desktop:

.. code-block:: txt

  1 : 299989
  2 : 99405
  3 : 200718
  4 : 149958
  5 : 249930

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10.1``.

References:

.. [1] `Generate Random Input from an Array according to given Probabilities - Techie Delight <http://www.techiedelight.com/generate-random-input-array-according-given-probabilities/>`_
.. [2] `[Golang] Generate Random String From [a-z0-9] <{filename}/articles/2015/04/13/go-generate-random-string%en.rst>`_
.. [3] `[Golang] All Permutations of Given String With Distinct Characters <{filename}/articles/2017/03/11/go-all-permutations-of-given-string-with-all-distinct-characters%en.rst>`_

.. _Go Playground: https://play.golang.org/
