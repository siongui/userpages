JavaScript basename()
#####################

:date: 2012-10-02 01:13
:modified: 2015-02-20 23:17
:tags: html, String Manipulation, JavaScript
:category: JavaScript
:summary: JavaScript equivalent of basename program.


JavaScript basename_
++++++++++++++++++++

.. code-block:: javascript

  /**
   * JavaScript version of basename
   * @param {string} path Example: a/b/c/d
   * @return {string} Example: d
   */
  basename = function(path) {
    var array = path.split('/');
    return array[array.length - 1];
  };

.. warning::

  The above code may not be robust for some cases like ``abc`` or ``/a/b/c/``

----

References:

.. [1] `[JavaScript] String startswith, endswith and contains Implementation <{filename}../../09/27/javascript-string-startswith-endswith-contains%en.rst>`_

.. [2] `Python basename <https://docs.python.org/2/library/os.path.html#os.path.basename>`_


.. _basename: http://en.wikipedia.org/wiki/Basename
