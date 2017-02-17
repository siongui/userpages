Pelican reStructuredText Link to Section in Another Page
########################################################

:date: 2016-04-16T20:05+08:00
:tags: Python, reStructuredText, Pelican
:category: Python
:summary: Pelican_ static site generator - Link to section in another page via
          reStructuredText_.
:adsu: yes


Introduction
++++++++++++

Pelican_ static site generator - Link to section in another page via
reStructuredText_.


Assume you have a rst_ `list-table`_ as follows:

.. code-block:: rst

  .. list-table:: Frozen Delights!
     :widths: 15 10 30
     :header-rows: 1

     * - Treat
       - Quantity
       - Description
     * - Albatross
       - 2.99
       - On a stick!
     * - Crunchy Frog
       - 1.49
       - If we took the bones out, it wouldn't be
         crunchy, now would it?
     * - Gannet Ripple
       - 1.99
       - On a stick!

You want to link to the table from another page. How to do it?

.. adsu:: 2

Solution
++++++++

First add a reference label right before your table:

.. code-block:: rst

  .. _my-reference-label:

  .. list-table:: Frozen Delights!
     :widths: 15 10 30
     :header-rows: 1

     * - Treat
       - Quantity
       - Description
     * - Albatross
       - 2.99
       - On a stick!
     * - Crunchy Frog
       - 1.49
       - If we took the bones out, it wouldn't be
         crunchy, now would it?
     * - Gannet Ripple
       - 1.99
       - On a stick!

Please note that use lower-case name in your label. Because if not, it will
become lower-case after converted to HTML.

Then in another page, you can link to the table as follows:

.. code-block:: rst

  `link to table in another page <{filename}/path/to/your/file.rst#my-reference-label>`_

Please note that use lower-case name if your reference label.

.. adsu:: 3

----

Tested on: ``Ubuntu Linux 15.10``, ``Python 2.7.10``, `Pelican 3.6.3`_.

----

References:

.. [1] `rst link to section - Google search <https://www.google.com/search?q=rst+link+to+section>`_

.. [2] `rst link to another page - Google search <https://www.google.com/search?q=rst+link+to+another+page>`_

.. [3] `python sphinx - RestructuredText: Adding a cross-reference to a subheading or anchor in another page - Stack Overflow <http://stackoverflow.com/questions/15394347/restructuredtext-adding-a-cross-reference-to-a-subheading-or-anchor-in-another>`_

.. [4] `舊文部份移植：文章選讀 · twnanda/twnanda@5529ded · GitHub <https://github.com/twnanda/twnanda/commit/5529ded532e76229f57f1fd84f134b726f5b8c8e>`_


.. _Python: https://www.python.org/
.. _reStructuredText: https://www.google.com/search?q=reStructuredText
.. _rst: https://www.google.com/search?q=reStructuredText
.. _list-table: http://docutils.sourceforge.net/docs/ref/rst/directives.html#list-table
.. _Pelican: http://blog.getpelican.com/
.. _Pelican 3.6.3: http://docs.getpelican.com/en/3.6.3/
