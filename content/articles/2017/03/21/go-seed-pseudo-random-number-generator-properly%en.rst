[Golang] Seed Pseudorandom Number Generator (PRNG) Properly
###########################################################

:date: 2017-03-21T15:08+08:00
:tags: Go, Golang, Random Number
:category: Go
:summary: How to seed the pseudorandom number generator (PRNG) properly in Go
          programming language.
:adsu: yes


In my `previous post`_, I use the `pseudo random number generator`_ (PRNG) of Go_
standard `math/rand`_ package to generate random string as follows:

.. code-block:: go

  package mylib

  import (
  	"math/rand"
  	"time"
  )

  func RandomString(strlen int) string {
  	// seed the PRNG
  	rand.Seed(time.Now().UTC().UnixNano())

  	// generate randsom string with rand.Intn method
  }

As you can see, I seed the PRNG on each use because I have no idea what is the
proper way to use the PRNG.

`@dchapes`_ left a comment_ to tell me something I do not know before:

- Do **not** re-seed PRNGs on each use.
- Do **not** re-seed the global PRNG in a package.

So I did some googling [1]_ [2]_, and at the same time `@shurcooL`_ left another
comment to show me how to `properly seed the PRNG`_:

  If you're using global ``rand.Rand``, don't re-seed it. But if you want to
  set your own seed, create your own local instance of ``rand.Rand``.

  .. code-block:: go

    var r *rand.Rand // Rand for this package.

    func init() {
    	r = rand.New(rand.NewSource(time.Now().UnixNano()))
    }

    func RandomString(strlen int) string {
    	// use existing r, no need to re-seed it
    	r.Intn(len(chars))
    }

  You can also `simplify this into one line`_:

  .. code-block:: go

    // r is a source of random numbers used in this package.
    var r = rand.New(rand.NewSource(time.Now().UnixNano()))

----

According to above comments, the most proper way to seed the PRNG in my case is
as follows:

.. code-block:: go

  package mylib

  import (
  	"math/rand"
  	"time"
  )

  // r is a source of random numbers used in this package.
  var r = rand.New(rand.NewSource(time.Now().UnixNano()))

  func RandomString(strlen int) string {
  	// generate randsom string with r.Intn method
  }

You can see my post for complete example of using PRNG to generate random
string. [3]_

----

Tested on: ``Ubuntu Linux 16.10``, ``Go 1.8``.

----

References:

.. [1] | `Do not reseed the global PRNG in a package - Google search <https://www.google.com/search?q=Do+not+reseed+the+global+PRNG+in+a+package>`_
       | `Do not reseed the global PRNG in a package - DuckDuckGo search <https://duckduckgo.com/?q=Do+not+reseed+the+global+PRNG+in+a+package>`_
       | `Do not reseed the global PRNG in a package - Ecosia search <https://www.ecosia.org/search?q=Do+not+reseed+the+global+PRNG+in+a+package>`_
       | `Do not reseed the global PRNG in a package - Qwant search <https://www.qwant.com/?q=Do+not+reseed+the+global+PRNG+in+a+package>`_
       | `Do not reseed the global PRNG in a package - Bing search <https://www.bing.com/search?q=Do+not+reseed+the+global+PRNG+in+a+package>`_
       | `Do not reseed the global PRNG in a package - Yahoo search <https://search.yahoo.com/search?p=Do+not+reseed+the+global+PRNG+in+a+package>`_
       | `Do not reseed the global PRNG in a package - Baidu search <https://www.baidu.com/s?wd=Do+not+reseed+the+global+PRNG+in+a+package>`_
       | `Do not reseed the global PRNG in a package - Yandex search <https://www.yandex.com/search/?text=Do+not+reseed+the+global+PRNG+in+a+package>`_

.. [2] `go - Golang random number generator how to seed properly - Stack Overflow <http://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly>`_
.. [3] `[Golang] Generate Random String From [a-z0-9] <{filename}../../../2015/04/13/go-generate-random-string%en.rst>`_

.. _previous post: {filename}../../../2015/04/13/go-generate-random-string%en.rst
.. _pseudo random number generator: https://www.google.com/search?q=pseudo+random+number+generator
.. _Go: https://golang.org/
.. _math/rand: https://golang.org/pkg/math/rand/
.. _Seed: https://golang.org/pkg/math/rand/#Seed
.. _time: https://golang.org/pkg/time/
.. _Go Playground: https://play.golang.org/
.. _@dchapes: https://github.com/dchapes
.. _comment: https://github.com/siongui/userpages/commit/77cd55346752ccaa2efa44b9084e97af81b664dd#commitcomment-21400225
.. _@shurcooL: https://github.com/shurcooL
.. _properly seed the PRNG: https://github.com/siongui/userpages/commit/77cd55346752ccaa2efa44b9084e97af81b664dd#commitcomment-21401369
.. _simplify this into one line: https://github.com/siongui/userpages/commit/6dc58d0b28a615ae19c43900358bcd258c1faac6#commitcomment-21402539
