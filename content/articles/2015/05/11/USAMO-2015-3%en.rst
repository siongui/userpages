[Math] United States of America Mathematical Olympiad (USAMO) 2015 Problem 3
############################################################################

:date: 2015-05-11T19:00-08:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary: United States of America Mathematical Olympiad (USAMO_) 2015 Problem 3

As high school math competition gets harder and harder, it becomes notable for
me to independently solve solve a 3/6 problem at IMO_. Given USAMO_ appears, in
my opinion, to be at least as hard as IMO_, I should write this down.

Problem: Let :math:`n` be a positive integer. We color each subset of :math:`N=\{1,2,...,n\}` with
blue or red, including empty set. Given a subset :math:`A` of :math:`N`, let :math:`f(A)` be the
number of blue subsets of :math:`A`. Find :math:`T(n)`, the number of colorings such that for any
subsets :math:`A` and :math:`B` of :math:`N`.

:math:`f(A)f(B)=f(A\cap B)f(A\cup B)`	(*)

At first I wasn't sure whether it's a combinatoric or functional problem, though there is never a
need, and actually bad, for every problem to belong to a specific category because that gives people
some non-intuitive clue to the approach. After spending quite some time working on it as a
functional problem, I didn't seem to get anywhere. I have to admit that I then shamelessly wrote a
simple code to compute some values of :math:`T(n)` and then guessed the pattern and prove it, which
is a great advantage, and never possible, in real competition.

Anyways, the key observation is this: what will happen if the empty set is red and
:math:`f(empthset)=0`? Then one of :math:`f(\{1\}),\dots,f(\{n\})` is zero. Let's
say :math:`f(\{1\})=0`, then one of :math:`f(\{1,2\}),f(\{1,3\}),\dots...,f(\{1, n\})` is zero.
Let's say :math:`f(\{1,2\})=0`, then...

Eventually, with f(\emptyset)=0 we arrive at a situation that without loss of generality, for any
subset :math:`A` of :math:`\{1,2,...,n-1\}` its color is red and :math:`f(A)=0`, and all sets that
may have non-zero f() value are those including n. Moreover, interestingly we could treat these
:math:`2^{n-1}` sets containing :math:`n` as another identical problem on
:math:`\{1,2,3,...,n-1\}`, as long as :math:`n` is temporarily ignored. Now either
:math:`f(\{n\})=1`, or :math:`f(\{n\})=0` and maybe some set containing :math:`n` is not 0, and so
on. So if :math:`G(n)` is the number of coloring with
:math:`f(\emptyset)=1`, we might have

:math:`T(n)=C(n,0)G(n)+C(n,1)G(n-1)+C(n,2)G(n-2)+...+C(n,n)G(0)+1`

The rough argument is to have a minimum set :math:`A` with :math:`f(A)=1`. For :math:`|A|=0` the
count is :math:`C(n,0)G(n)`, for :math:`|A|=1` the count is :math:`C(n,1)G(n-1)`, and so on. 1
counts for the coloring with :math:`f(A)=0` for every :math:`A`.

Again, I then shamelessly changed my code to figure out :math:`G(n)=2^n`, and so

:math:`T(n)=C(n,0)2^n+C(n,1)2^{n-1}+C(n,2)2^{n-2}+\dots+C(n,n)2^0+1=(1+2)^n+1=3^n+1`

Now the rigorous proof will strengthen the arguments and show :math:`G(n)=2^n`. As always, it
appears much less intuitive when we make it formal.

Proof
+++++

Throughout the proof we ignore the trivial coloring with all subsets labeled red, and will add it
back at the last step. We're also only concerned with coloring that makes (*) hold. We proceed with
a series of lemmas, some of which self-evident.

Lemma 1
-------
:math:`f(A)>0` iff :math:`A` contains a blue subset.



Lemma 2
-------
There is an unique minimum blue subset :math:`S`.

To prove it, consider two different minimum blue subsets :math:`S` and :math:`S'` with same size.
Then :math:`f(S\cap S') > 0`. Since :math:`S\cap S'` is smaller than :math:`S` and :math:`S'`, by
Lemma 1 this is a contradiction.
:math:`\Diamond`



Lemma 3
--------------
:math:`f(A)>0` iff :math:`A` contains the minimum blue subset :math:`S`.

If :math:`f(A)>0` then :math:`B=S\cap A` also has :math:`f(B)>0`, meaning :math:`B` has a blue
subset. Now if :math:`A` doesn't contain :math:`S`, then :math:`|B|<|S|`, a contradiction.
:math:`\Diamond`



Lemma 4
-------
Let :math:`T(n)` be the answer to the original problem, and :math:`G(n)` be the number
of colorings with :math:`f(\emptyset)=1`. Then :math:`T(n)=C(n,0)G(n)+C(n,1)G(n-1)+\dots+C(n,n)G(0)+1`

Let the minimum blue subset be :math:`S`. By Lemma 3 all subsets not containing :math:`S` should be
red. We can check that (1) the problem of coloring subsets containing S is equivalent to the original problem on set :math:`N-S` when we map each :math:`X` to :math:`X-S`. (2) (*) still holds whether
:math:`A` and :math:`B` contains :math:`S` or not. So by iterating through all possible :math:`S`
we get the :math:`T(n)` as described.
:math:`\Diamond`



Finally, it suffices to show that :math:`G(n)=2^n`. This is the most interesting part to me.

The idea is to prove that each coloring of singletons :math:`\{1\}, \{2\}, \{3\},\ldots,\{n\}`
leads to one and only one possible colorings of the rest, and therefore :math:`G(n)=2^n`. 

To show there's only one possible coloring given a coloring of all singletons, first observe that
:math:`f(\{i\})` is either 1 or 2. So :math:`f(doubleton)` is either 1, 2, or 4. By induction on
:math:`|A|` we can see that :math:`f(A)` can take at most one value, which is a power of 2, because
:math:`f(A)=f(\{x\})f(A-\{x\})/f(\emptyset)=f(\{x\})f(A-\{x\})` where :math:`x` belongs to
:math:`A`. Now the last remaining piece is to show such :math:`f` makes (*) hold all the time. 

For any :math:`A`, because the only possible :math:`f(A)` is constructed as product of other
:math:`f()` which are power of 2, :math:`L(A)=\log_2(f(A))` is always integer, and
:math:`L(A)=L(\{x\})+L(A-\{x\})`. Therefore :math:`L(A)` is the number of blue singletons contained
by :math:`A`! Now how are non-singletons colored so that :math:`f(A)` is also the number of blue
subsets owned by :math:`A`? We can color :math:`A` blue iff all singletons of :math:`A` are blue. Then
:math:`f'(A)`, the function defined based on this coloring, is the number of subsets of :math:`A`
containing only singletons that are blue, i.e. :math:`f'(A)=2^{L(A)}=f(A)`, so this only possible
:math:`f(A)` is also achievable.

Q.E.D.

An example of coloring with :math:`f(\emptyset)=1` when :math:`n=4`. Notation: color/f/L

:math:`\emptyset: B/1/0`

:math:`\{1\}: B/2/1,\{2\}: R/1/0,\{3\}: R/1/0,\{4\}: B/2/1`

:math:`\{1,2\}: R/2/1,\{1,3\}: R/2/1,\{1,4\}: B/4/2,\{2,3\}: R/1/0,\{2,4\}: R/2/1,\{3,4\}: R/2/1`

:math:`\{1,2,3\}: R/2/1,\{1,2,4\}: R/4/2,{1,3,4}: R/4/2,\{2,3,4\}: R/2/1`

:math:`\{1,2,3,4\}: R/4/2`

----

`post <http://oathbystyx.blogspot.com/2015/05/2015-usamo-3.html>`_
by
`Shen-Fu Tsai <{filename}/pages/sftsai.rst>`_

----

References:

.. [1] `USAMO Problems and Solutions <http://www.artofproblemsolving.com/wiki/index.php/USAMO_Problems_and_Solutions>`_

.. [2] `International Mathematical Olympiad <https://www.imo-official.org/>`__


.. _International Mathematical Olympiad: https://www.imo-official.org/
.. _IMO: https://www.imo-official.org/
.. _USAMO: http://www.maa.org/math-competitions
