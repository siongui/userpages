[web.py] Web Application on Both Google App Engine and Apache
#############################################################

:date: 2015-03-04 23:07
:modified: 2015-03-05 09:33
:tags: web.py, Python, Google App Engine, Apache
:category: Python
:summary: Web application (web.py framework) template runs on both Google App
          Engine Python and Apache with mod_wsgi.


This post shows a web.py application template which runs on both Apache_ with
mod_wsgi_ and `Google App Engine Python`_.

"Hello World" web.py_ application
+++++++++++++++++++++++++++++++++

.. show_github_file:: siongui userpages content/code/webpy-gae-apache/mainweb.py

.. note::

  web.py_ is not included in the `third-party libraries in GAE Python 2.7`_. To
  use web.py on GAE, please `download web.py from Github`_. Put the ``web``
  directory in web.py repo and this "Hello World" application in the same
  directory.

Sample GAE Python config
++++++++++++++++++++++++

.. show_github_file:: siongui userpages content/code/webpy-gae-apache/app.yaml

Sample Apache config
++++++++++++++++++++

.. show_github_file:: siongui userpages content/code/webpy-gae-apache/apache.conf

Development
+++++++++++

.. show_github_file:: siongui userpages content/code/webpy-gae-apache/Makefile

Modify the path of ``GAE_PY_SDK`` in *Makefile* to your path of GAE Python SDK.

Test run the web.py_ application locally:

.. code-block:: bash

  # open your terminal and run
  $ make local
  # OR
  $ make
  # open browser with URL: http://localhost:8080/

Test run the web.py_ application on local GAE Python environment:

.. code-block:: bash

  # open your terminal and run
  $ make devserver
  # open browser with URL: http://localhost:8080/


Tested on: ``Ubuntu Linux 14.10``, ``Google App Engine Python SDK 1.9.18``

----

References:

.. [1] `web.py <http://webpy.org/>`_

.. [2] `Webpy + Apache with mod_wsgi on Ubuntu (web.py) <http://webpy.org/cookbook/mod_wsgi-apache-ubuntu>`_

.. [3] `Webpy + Google App Engine (web.py) <http://webpy.org/cookbook/google_app_engine>`_

.. [4] `[web.py] Multiple Application with Same Context <{filename}../09/webpy-multiple-app-with-same-context%en.rst>`_


.. _web.py: http://webpy.org/

.. _Apache: http://httpd.apache.org/

.. _mod_wsgi: https://code.google.com/p/modwsgi/

.. _Google App Engine Python: https://cloud.google.com/appengine/docs/python/

.. _third-party libraries in GAE Python 2.7: https://cloud.google.com/appengine/docs/python/tools/libraries27

.. _download web.py from Github: https://github.com/webpy/webpy
