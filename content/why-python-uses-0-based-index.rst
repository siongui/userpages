Why Python uses 0-based indexing?
#################################

:date: 2013-10-26 19:27
:lang: en
:tags: Python
:category: Python
:slug: why-python-uses-0-based-indexing
:author: Siong-Ui Te
:summary: Guido van Rossum (father of Python) explains why Python uses 0-based indexing.


Here is the link_ of Guido van Rossum's explanation. The following is the post by him on Google+:

::

  I was asked on Twitter why Python uses 0-based indexing, with a link to a new (fascinating) post on the subject (http://exple.tive.org/blarg/2013/10/22/citation-needed/). I recall thinking about it a lot; ABC, one of Python's predecessors, used 1-based indexing, while C, the other big influence, used 0-based. My first few programming languages (Algol, Fortran, Pascal) used 1-based or variable-based. I think that one of the issues that helped me decide was slice notation.

  Let's first look at use cases. Probably the most common use cases for slicing are "get the first n items" and "get the next n items starting at i" (the first is a special case of that for i == the first index). It would be nice if both of these could be expressed as without awkward +1 or -1 compensations.

  Using 0-based indexing, half-open intervals, and suitable defaults (as Python ended up having), they are beautiful: a[:n] and a[i:i+n]; the former is long for a[0:n].

  Using 1-based indexing, if you want a[:n] to mean the first n elements, you either have to use closed intervals or you can use a slice notation that uses start and length as the slice parameters. Using half-open intervals just isn't very elegant when combined with 1-based indexing. Using closed intervals, you'd have to write a[i:i+n-1] for the n items starting at i. So perhaps using the slice length would be more elegant with 1-based indexing? Then you could write a[i:n]. And this is in fact what ABC did -- it used a different notation so you could write a@i|n.(See http://homepages.cwi.nl/~steven/abc/qr.html#EXPRESSIONS.)

  But how does the index:length convention work out for other use cases? TBH this is where my memory gets fuzzy, but I think I was swayed by the elegance of half-open intervals. Especially the invariant that when two slices are adjacent, the first slice's end index is the second slice's start index is just too beautiful to ignore. For example, suppose you split a string into three parts at indices i and j -- the parts would be a[:i], a[i:j], and a[j:].

  So that's why Python uses 0-based indexing.


.. _link: https://plus.google.com/115212051037621986145/posts/YTUxbXYZyfi
