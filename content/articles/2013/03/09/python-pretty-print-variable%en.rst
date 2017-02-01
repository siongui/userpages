[Python] Pretty Print Variable (Dictionary, List, Object, Array)
################################################################

:tags: pretty print, Python
:category: Python
:summary: Pretty print variable in Python_
:adsu: yes


The easiest way is through json_ module included in Python_:

.. code-block:: python

    print(json.dumps(obj, sort_keys=True, indent=4, separators=(',', ': ')))

where *obj* is your variable, could be dictionary_ (equivalent to object in
JavaScript_) or list_ (equivalent to array in JavaScript).

.. adsu:: 2

.. _Python: https://www.python.org/
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _json: https://docs.python.org/2/library/json.html
.. _dictionary: https://docs.python.org/2/library/stdtypes.html#mapping-types-dict
.. _list: https://docs.python.org/2/library/stdtypes.html#sequence-types-str-unicode-list-tuple-bytearray-buffer-xrange
