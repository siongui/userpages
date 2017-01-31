Special Format Conversion in Python
###################################

:tags: Python, String Manipulation
:category: Python
:summary: `%2c%e3%80%90%e5%bd%a2%e3%80%91%e6%97%a0%e7%a2%8d%e7%9a%84%e3%80%82` <=> `,【形】无碍的。`
:adsu: yes


The folloing string:

.. code-block:: bash

  %2c%e3%80%90%e5%bd%a2%e3%80%91%e6%97%a0%e7%a2%8d%e7%9a%84%e3%80%82

Pass to the following function

.. code-block:: python

  def HexStringToString(hexString):
    # convert hex string to utf8 string
    # example: "%2c%e3%80" -> "\x2C\xE3\x80"
    bytes = []
    hexStr = ''.join( hexString.split("%") )
    for i in range(0, len(hexStr), 2):
      bytes.append( chr( int (hexStr[i:i+2], 16 ) ) )

    # decode as utf8
    string = ''.join( bytes ).decode("utf-8")

    return string

Will become

.. code-block:: bash

  ,【形】无碍的。

.. adsu:: 2

Pass again to the following function

.. code-block:: python

  def StringToHexString(string):
    # convert string to hex string
    string = string.encode('utf8')
    return ''.join( [ "%02X " % ord( x ) for x in string ] ).strip().replace(' ','%')

Will become back

.. code-block:: bash

  2C%E3%80%90%E5%BD%A2%E3%80%91%E6%97%A0%E7%A2%8D%E7%9A%84%E3%80%82

.. adsu:: 3

----

References:

.. [1] `Byte to Hex and Hex to Byte String Conversion (Python recipe) <http://code.activestate.com/recipes/510399-byte-to-hex-and-hex-to-byte-string-conversion/>`_
.. [2] `Convert bytes to a Python string <http://stackoverflow.com/questions/606191/convert-bytes-to-a-python-string>`_
