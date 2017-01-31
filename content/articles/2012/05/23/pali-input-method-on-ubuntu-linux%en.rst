Pāli Input Method on Ubuntu Linux
#################################

:tags: Pāli Input Method, Ubuntu Linux
:category: Pāli
:summary: Type romanized Pāli on your Ubuntu Linux.
:adsu: yes


The blog post [1]_ discusses how to create an input method for romanized
`Pāli`_, but it doesn't give references for tweaking the mim file, i.e.,
tweaking the input method. Also, the blog post [1]_ seems to come from the
thread [2]_.

The blog post is good, but I need another feature that when I type the key
combination "alt + a", ā will show up. The same for other vowels. After some
investigation, removing Sanskrit_ part, and making necessary modifications, the
following is my *mim* file (tested on `Ubuntu 12.04`_).

.. adsu:: 2
.. show_github_file:: siongui userpages content/articles/2012/05/23/pali-translit.mim
.. adsu:: 3

----

References:

.. [1] `How to use iBus to input Romanized Pali and Sanskrit in Ubuntu <http://thanhsiang.org/faqing/node/109>`_
.. [2] `Sanskrit input (Devanagari and transliteration) <http://ubuntuforums.org/showthread.php?t=646207>`_
.. [3] `Data format of the m17n database <http://www.nongnu.org/m17n/manual-en/m17nDBFormat.html>`_
.. [4] `ITRANS <http://en.wikipedia.org/wiki/ITRANS>`_

.. _Pāli: https://en.wikipedia.org/wiki/Pali
.. _Sanskrit: https://en.wikipedia.org/wiki/Sanskrit
.. _Ubuntu 12.04: http://releases.ubuntu.com/12.04/
