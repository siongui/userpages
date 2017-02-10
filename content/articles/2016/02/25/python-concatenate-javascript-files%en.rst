[Python] Concatenate JavaScript Files
#####################################

:date: 2016-02-25T13:46+08:00
:tags: Python, JavaScript, Commandline, List Files in Directory,
       File Input/Output
:category: Python
:summary: Concatenate JavaScript_ files via Python_ script.
:adsu: yes

Concatenate JavaScript_ files via Python_ script.

.. code-block:: python

  #!/usr/bin/env python
  # -*- coding: utf-8 -*- #

  import os

  def concatJS(dirpath):
      combinedJS = ""
      for filename in os.listdir(dirpath):
          if filename.endswith(".js"):
              filepath = os.path.join(dirpath, filename)
              print("combining " + filepath + " ...")
              with open(filepath, "r") as f:
                  combinedJS += f.read()

      return combinedJS

  if __name__ == "__main__":
      print(concatJS("."))

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 15.10``, ``Python 2.7.10``.

----

References:

.. [1] `python list file in directory <https://www.google.com/search?q=python+list+file+in+directory>`_

.. _Python: https://www.python.org/
.. _JavaScript: https://www.google.com/search?q=javascript
