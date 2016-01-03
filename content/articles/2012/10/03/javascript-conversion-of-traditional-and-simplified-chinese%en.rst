[JavaScript] Conversion of Traditional and Simplified Chinese
#############################################################

:date: 2012-10-03 20:37
:tags: JavaScript, String Manipulation, Conversion of Traditional and Simplified Chinese
:category: JavaScript
:summary: JavaScript library `New Tong Wen Tang` for conversion of Traditional
          and Simplified Chinese.


There is a easy-to-use library for conversion of Traditional and Simplified
Chinese called **New Tong Wen Tang** [1]_. There are total 5 JavaScript files in
the library (see following code snippet):

.. code-block:: html

  <script language="JavaScript" src="http://tongwen.openfoundry.org/src/web/tongwen_core.js"></script>
  <script language="JavaScript" src="http://tongwen.openfoundry.org/src/web/tongwen_table_s2t.js"></script>
  <script language="JavaScript" src="http://tongwen.openfoundry.org/src/web/tongwen_table_t2s.js"></script>
  <script language="JavaScript" src="http://tongwen.openfoundry.org/src/web/tongwen_table_ps2t.js"></script>
  <script language="JavaScript" src="http://tongwen.openfoundry.org/src/web/tongwen_table_pt2s.js"></script>

**tongwen_core.js** is mandatory and must be included. **tongwen_table_s2t.js**
and **tongwen_table_ps2t.js** are tables for conversion from *Simplified to
Traditional Chinese*. **tongwen_table_t2s.js** and **tongwen_table_pt2s.js** are
tables for conversion from *Traditional to Simplified Chinese*. Include these
tables according to your requirements.

Usage
+++++

To convert entire HTML document from Traditional to Simplified Chinese:

.. code-block:: javascript

  TongWen.trans2Trad(document);

To convert entire HTML document from Simplified to Traditional Chinese:

.. code-block:: javascript

  TongWen.trans2Simp(document);

To convert a string *str* from Traditional to Simplified Chinese:

.. code-block:: javascript

  str = TongWen.convert(str, TongWen.flagSimp);

To convert a string *str* from Simplified to Traditional Chinese:

.. code-block:: javascript

  str = TongWen.convert(str, TongWen.flagTrad);

----

References:

.. [1] `新同文堂 for Web <http://tongwen.openfoundry.org/web.htm>`_

.. [2] `[Golang] Conversion of Traditional and Simplified Chinese <{filename}../../../2016/01/03/go-conversion-of-traditional-and-simplified-chinese%en.rst>`_

.. [3] `[Python] Conversion of Traditional and Simplified Chinese <{filename}../../../2016/01/04/python-conversion-of-traditional-and-simplified-chinese%en.rst>`_
