[Python] Convert Footnote to reStructuredText Format
####################################################

:date: 2016-02-15T07:30+08:00
:tags: Python, Regular Expression, reStructuredText
:category: Python
:summary: Python_ `Regular Expression`_ to convert footnotes in file to
          reStructuredText_ format.


Problem
+++++++

1. Convert ``(01-001)``, ``(01-002)``, ... to ``[01-001]_``, ``[01-002]_``, ...

2. Convert ``〔註01-001〕``, ``〔註01-002〕``, ... to ``.. [01-001]``, ``.. [01-002]``, ...


Solution
++++++++

.. code-block:: python

  #!/usr/bin/env python
  # -*- coding:utf-8 -*-

  import re
  import os

  def findFootnote(filePath):
    with open(filePath, 'r') as f:
      for line in f:
        p = re.compile(r'\(01-[\d]{3}\)')
        notes = p.findall(line)
        if len(notes) > 0:
          print(notes)

  def replaceFootnote(filePath, output):
    with open(filePath, 'r') as f:
      with open(output, 'w') as fo:
        for line in f:
          result = re.sub(r'\(01-([\d]{3})\)', r' [01-\1]_ ', line)
          fo.write(result)

  def rstripFile(filePath, output):
    with open(filePath, 'r') as f:
      with open(output, 'w') as fo:
        for line in f:
          fo.write(line.rstrip() + '\n')

  def replaceFootnote2(filePath, output):
    with open(filePath, 'r') as f:
      with open(output, 'w') as fo:
        for line in f:
          result = re.sub(r'〔註01-([\d]{3})〕', r'.. [01-\1]', line)
          fo.write(result)

  if __name__ == '__main__':
    #findFootnote("../content/articles/2016/02/14/visuddhimagga-chap01%zh.rst")
    path = "../content/articles/2016/02/14/visuddhimagga-chap01%zh.rst"
    #replaceFootnote(path, os.path.basename(path))
    #rstripFile(path, os.path.basename(path))
    replaceFootnote2(path, os.path.basename(path))


----

Tested on: ``Ubuntu Linux 15.10``, ``Python 2.7.10``.

----

References:

.. [1] `Python Regular Expressions  |  Google for Education  |  Google Developers <https://developers.google.com/edu/python/regular-expressions>`_

.. [2] `Regex replace (in Python) - a simpler way? - Stack Overflow <http://stackoverflow.com/questions/490597/regex-replace-in-python-a-simpler-way>`_

.. _Python: https://www.python.org/
.. _Regular Expression: https://developers.google.com/edu/python/regular-expressions
.. _reStructuredText: https://www.google.com/search?q=reStructuredText
