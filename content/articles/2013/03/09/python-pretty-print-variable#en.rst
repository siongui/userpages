[Python] Pretty Print Variable (Dictionary, List, Object, Array)
################################################################

:tags: array, dictionary, list, object, pretty print, Python, variable
:category: Python
:summary: Pretty print variable in Python


The easiest way is through `json <http://docs.python.org/2/library/json.html>`_ module included in Python:

.. code-block:: python

    print(json.dumps(obj, sort_keys=True, indent=4, separators=(',', ': ')))

where *obj* is your variable, could be `dictionary <http://docs.python.org/2/library/stdtypes.html#mapping-types-dict>`_ (equivalent to object in JavaScript) or `list <http://docs.python.org/2/library/stdtypes.html#sequence-types-str-unicode-list-tuple-bytearray-buffer-xrange>`_ (equivalent to array in JavaScript).
