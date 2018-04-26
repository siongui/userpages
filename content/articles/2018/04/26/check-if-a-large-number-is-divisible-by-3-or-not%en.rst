Check If A Large Number Is Divisible By 3 Or Not
################################################

:date: 2018-04-26T23:46+08:00
:tags: Go, Golang, Math, Algorithm, Type Casting, Type Conversion
:category: Algorithm
:summary: Check if a large number is divisible by 3 or not in Go. This exercise
          is good example for type casting between *int* and *string* in Go.
:og_image: http://www.mathwarehouse.com/arithmetic/numbers/images/divisibility-rules/divisibility-by-6-rule.png
:adsu: yes


I saw this exercise from GeeksforGeeks [1]_, and it is a good example for type
casting/conversion between *int* and *string* in Go.

The question is: *Given a number, check if it is divisible by 3 or not.*

The answer is: *If a number is divisible by 3 if the sum of its digits is
divisible by 3.* For example, 12345 is divisible by 3, because 1+2+3+4+5=15, and
15 is divisible by 3.

Since Go is strongly-typed, we assume the number can be passed as either *int*
or *string* type.


Number as *int*
+++++++++++++++

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/E1tdbBjw5cs>`__
   :class: align-center

.. code-block:: go

  import (
  	"strconv"
  )

  func IsIntDivisibleBy3(n int) bool {
  	digits := strconv.Itoa(n)
  	sumOfDigits := 0
  	for _, digit := range digits {
  		d, _ := strconv.Atoi(string(digit))
  		sumOfDigits += d
  	}

  	return (sumOfDigits % 3) == 0
  }


Example:

.. code-block:: go

  import (
  	"testing"
  )

  func TestIsIntDivisibleBy3(t *testing.T) {
  	if IsIntDivisibleBy3(123456758933312) != false {
  		t.Error(123456758933312)
  	}
  	if IsIntDivisibleBy3(769452) != true {
  		t.Error(769452)
  	}
  }


Number as *string*
++++++++++++++++++

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/1vVUqFCmWYR>`__
   :class: align-center

.. code-block:: go

  import (
  	"strconv"
  )

  func IsStrDivisibleBy3(n string) (bool, error) {
  	sumOfDigits := 0
  	for _, digit := range n {
  		d, err := strconv.Atoi(string(digit))
  		if err != nil {
  			return false, err
  		}
  		sumOfDigits += d
  	}

  	return (sumOfDigits % 3) == 0, nil
  }


Example:

.. code-block:: go

  import (
  	"testing"
  )

  func TestIsStrDivisibleBy3(t *testing.T) {
  	result, err := IsStrDivisibleBy3("3635883959606670431112222")
  	if err != nil {
  		t.Error(err)
  	}
  	if result != true {
  		t.Error("3635883959606670431112222")
  	}

  	result, err = IsStrDivisibleBy3("123456758933312")
  	if err != nil {
  		t.Error(err)
  	}
  	if result != false {
  		t.Error("123456758933312")
  	}

  	result, err = IsStrDivisibleBy3("769452")
  	if err != nil {
  		t.Error(err)
  	}
  	if result != true {
  		t.Error("769452")
  	}
  }

.. adsu:: 2

Tested on:

- ``Ubuntu Linux 17.10``, ``Go 1.10.1``
- `Go Playground`_

----

References:

.. [1] `Check if a large number is divisible by 3 or not - GeeksforGeeks <https://www.geeksforgeeks.org/check-large-number-divisible-3-not/>`_

.. _Go Playground: https://play.golang.org/
