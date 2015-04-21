The infamous Grasshopper problem
################################

:date: 2015-04-01T18:09-08:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary: The infamous Grasshopper problem


The infamous Grasshopper problem
++++++++++++++++++++++++++++++++

Let :math:`a_1`, :math:`a_2`, ... , :math:`a_n` be distinct positive integers
and let :math:`M` be a set of :math:`n − 1` positive integers not containing
:math:`s = a_1 + a_2 + ... + a_n`. A grasshopper is to jump along the real axis,
starting at the point :math:`0` and making :math:`n` jumps to the right with
lengths :math:`a_1, a_2, ... , a_n` in some order. Prove that the order can be
chosen in such a way that the grasshopper never lands on any point in :math:`M`.


Proof
+++++

First note that if the statement is still equivalent if :math:`|M| < n`.

Let the set :math:`M = \{ b_1, b_2, ... , b_{n-1} \}`, where
:math:`0 < b_1 < b_2 < ... < b_{n-1} < s`. We will prove by induction on
:math:`n`. :math:`n = 1` and :math:`n = 2` are trivial. Suppose the statement
holds for all :math:`n <= m`. We want to prove that it holds for
:math:`n = m + 1 >= 3`.

Lemma
+++++

If there's any :math:`a_i > b_1` that doesn't belong to :math:`M`, then we're
done here. This is because placing :math:`a_i` at :math:`[0, a_i]` covers at
least one point :math:`b_1` in :math:`M`, and there are now :math:`n - 1`
segments and less than :math:`n - 1` points left, so by induction the statement
holds.

So suppose there isn't such :math:`a_i`. There are 3 possibilities:

  (a) :math:`a_n < b_1`

  (b) :math:`a_n = b_1`

  (c) :math:`a_n = b_j > b_1`

We will deal with them in reverse order.


:math:`a_n = b_j > b_1`
```````````````````````

Now :math:`a_n = b_j` is in :math:`M`, so we proceed from :math:`a_{n-1}`,
:math:`a_{n-2}`, ... to :math:`a_1`. We stop when we see an :math:`a_i` that's
not in :math:`M`.  Let's say we stop at :math:`a_{n-k}` with :math:`k > 0`. Note
we will stop at some point because in :math:`A` there can be at most
:math:`|M| = n - 1` numbers in :math:`M`.

  (i) :math:`a_{n-k} > b_1`. We're done because of the lemma.

  (ii) :math:`a_{n-k} < b_1`. So :math:`|A \cap M| = k > 0`, i.e.
       :math:`A \cap M = \{a_n, a_{n-1}, ... , a_{n-k+1}\}` This is the most
       interesting case to me.

How many points in :math:`M` are larger than :math:`a_n`? At most
:math:`n - 1 - k` because :math:`k` of them are in the intersection of :math:`M`
and :math:`A`. How many :math:`a_i` in :math:`A` are less than :math:`b_1`?
:math:`n - k`. Therefore, there is at least an :math:`a_i < b_1` such that
:math:`a_i + a_n` is not in :math:`M`. If we put :math:`a_i` at :math:`[0, a_i]`
and :math:`a_n` at :math:`[a_i, a_i + a_n]`, we're sure these two segments have
no end points in :math:`M` because :math:`a_i < b_1` and :math:`a_i + a_n` is
not in :math:`M`, yet they cover at least 2 points in :math:`M`, i.e.
:math:`b_1` and :math:`a_n = b_j`. By induction, the remaining
:math:`n - 2 >= 1` segments can be arranged such that no end points fall into
:math:`M`.

:math:`a_n = b_1`
`````````````````

Place :math:`a_n` at :math:`[0, a_n]`. Let :math:`A' = A - \{ a_n \}` and
:math:`M' = M - \{ b_1 \}`. By induction we can arrange :math:`A'` in such a way
that no points in :math:`M'` coincide with an end point in :math:`A'`. Now only
:math:`b_1` touches end points of :math:`a_n` and :math:`a_j`. Since
:math:`a_j < a_n`, swapping :math:`a_j` and :math:`a_n` resolves it.

:math:`a_n < b_1`
`````````````````

This is similar. Let's put :math:`a_n` at the :math:`[0, a_n]`. There are
:math:`n - 1` segments but :math:`n - 1` points left, so let's pretend
:math:`b_1` doesn't exist. That is, consider :math:`M' = M - \{ b_1 \}`. Now by
induction we can arrange the remaining :math:`n - 1` segments such that no point
in :math:`M'` is touched by any end point. Only :math:`b_1` might coincide with
an end point of :math:`A - \{ a_n \}`. Let's assume the current arrangement is

:math:`a_n, a_{s_1}, a_{s_2}, ... a_{s_{n-1}}`

where :math:`s_1, ... , s_{n-1}` is a permutation of :math:`\{1, ... , n - 1\}`.
Suppose the end points of :math:`a_{s_j}, a_{s_{j+1}}` is :math:`b_1`. Then we
could rearrange it as

:math:`a_{s_1}, a_{s_2}, ... , a_{s_{j+1}}, a_n, a_{s_{j+2}}, a_{s_{j+3}} ...`

Because :math:`a_n > a_{s_{j+1}}`, :math:`b_1` is no longer touched by any end
point so we're done.

Q.E.D.

----

`post <http://oathbystyx.blogspot.com/2015/04/the-infamous-grasshopper-problem.html>`_
by
`Shen-Fu Tsai <https://plus.google.com/102515651050568228591>`_
