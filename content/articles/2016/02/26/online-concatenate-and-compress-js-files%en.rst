[Python] Online Concatenate and Compress JavaScript Files
#########################################################

:date: 2016-02-26T15:05+08:00
:tags: Python, JavaScript, Commandline
:category: Python
:summary: Online concatenate and compress JavaScript_ files via Python_ script.

Online concatenate and compress JavaScript_ files via Python_ script.

Assume we want to concatenate and compress the following js files into a single
js file.

.. code-block:: bash

  tongwen_core.js
  tongwen_table_ps2t.js
  tongwen_table_pt2s.js
  tongwen_table_s2t.js
  tongwen_table_t2s.js
  nanda.js

The ``tongwen_*.js`` are libraries and unchanged during our development. The
only js file that we code with is ``nanda.js``. We commit the above js files to
GitHub_, and write the following Python_ script to concatenate and compile the
js files via **online Closure Compiler** [1]_ [2]_ [3]_ [4]_.
(put this py script in the same directory as js files)

.. code-block:: python

  #!/usr/bin/env python
  # -*- coding: utf-8 -*- #

  import httplib, urllib

  allJS = ["nanda.js"]

  def combineJS():
      js_code = ""
      for name in allJS:
          with open(name, 'r') as f:
              js_code += f.read()
      return js_code


  def online_compile(js_code):
      # Define the parameters for the POST request and encode them in
      # a URL-safe format.

      params = urllib.urlencode([
          ('code_url', 'https://github.com/twnanda/twnanda/raw/master/theme/javascript/tongwen_core.js'),
          ('code_url', 'https://github.com/twnanda/twnanda/raw/master/theme/javascript/tongwen_table_ps2t.js'),
          ('code_url', 'https://github.com/twnanda/twnanda/raw/master/theme/javascript/tongwen_table_pt2s.js'),
          ('code_url', 'https://github.com/twnanda/twnanda/raw/master/theme/javascript/tongwen_table_s2t.js'),
          ('code_url', 'https://github.com/twnanda/twnanda/raw/master/theme/javascript/tongwen_table_t2s.js'),
          ('js_code', js_code),
          ('compilation_level', 'SIMPLE_OPTIMIZATIONS'),
          ('language', 'ECMASCRIPT5'),
          ('output_format', 'text'),
          ('output_info', 'compiled_code'),
      ])

      # Always use the following value for the Content-type header.
      headers = { "Content-type": "application/x-www-form-urlencoded" }
      conn = httplib.HTTPConnection('closure-compiler.appspot.com')
      conn.request('POST', '/compile', params, headers)
      response = conn.getresponse()
      print(response.read())
      conn.close()


  if __name__ == "__main__":
      online_compile(combineJS())


Note that *code_url* is the URL of our raw files on GitHub, which tells the
Closure Compiler to fetch files on GitHub instead of uploading them locally. The
only file we *POST* to Closure Compiler service API is ``nanda.js``, which we
code with in our development.

Save above py script as *compile.py* and put in the same directory as js files.
Run the script and save concatenated and compressed js as ``all.js`` by:

.. code-block:: bash

  $ python compile > all.js

----

Tested on: ``Ubuntu Linux 15.10``, ``Python 2.7.10``.

----

References:

.. [1] `Closure Compiler service API <https://www.google.com/search?q=Closure+Compiler+service+API>`_

.. [2] `Getting Started with the API  |  Closure Compiler  |  Google Developers <https://developers.google.com/closure/compiler/docs/gettingstarted_api>`_

.. [3] `Communicating with the Closure Compiler Service API  |  Closure Compiler  |  Google Developers <https://developers.google.com/closure/compiler/docs/api-tutorial1>`_

.. [4] `Compressing Files with the Closure Compiler Service API  |  Closure Compiler  |  Google Developers <https://developers.google.com/closure/compiler/docs/api-tutorial2>`_

.. _Python: https://www.python.org/
.. _JavaScript: https://www.google.com/search?q=javascript
.. _GitHub: https://github.com/
