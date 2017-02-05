[web.py] Multiple Application with Same Context
###############################################

:date: 2015-03-09 20:18
:tags: web.py, Python, Google App Engine, Apache
:category: Python
:summary: Multiple web.py applications with the same context (web.ctx)
:adsu: yes


In web.py_ official site, there is a code example which shows multiple
sub applications [1]_. In the example, the context ``web.ctx`` are modified in
the sub application. This post shows multiple applications with the same context
``web.ctx``, and runs both on Apache_ with mod_wsgi_ and
`Google App Engine Python`_ (the same as [3]_).

The point is to overwrite the *_delegate_sub_application* in *web.applcation*
class. a *customApp* is declared:

.. code-block:: python

  class customApp(web.application):
    def _delegate_sub_application(self, dir, app):
      return app.handle_with_processors()

The *customApp* will be used for the mapping of applications.

.. adsu:: 2

Complete Source Code
++++++++++++++++++++

delegate requests app
`````````````````````

.. show_github_file:: siongui userpages content/code/webpy-multiple-application/delegate.py
.. adsu:: 3

.. note::

  web.py_ is not included in the `third-party libraries in GAE Python 2.7`_. To
  use web.py on GAE, please `download web.py from Github`_. Put the ``web``
  directory in web.py repo and this "Hello World" application in the same
  directory.

main app
````````

.. show_github_file:: siongui userpages content/code/webpy-multiple-application/mainweb.py

static app
``````````

.. show_github_file:: siongui userpages content/code/webpy-multiple-application/staticweb.py

sample GAE Python config
````````````````````````

.. show_github_file:: siongui userpages content/code/webpy-multiple-application/app.yaml

sample Apache config
````````````````````

.. show_github_file:: siongui userpages content/code/webpy-multiple-application/apache.conf

sample Makefile for development
```````````````````````````````

.. show_github_file:: siongui userpages content/code/webpy-multiple-application/Makefile



Tested on: ``Ubuntu Linux 14.10``, ``Google App Engine Python SDK 1.9.18``

----

References:

.. [1] `using subapplications (web.py) <http://webpy.org/cookbook/subapp>`_

.. [2] `Django style multiple apps with web.py (web.py) <http://webpy.org/multiple_apps>`_

.. [3] `[web.py] Web Application on Both Google App Engine and Apache <{filename}../04/webpy-gae-apache%en.rst>`_


.. _web.py: http://webpy.org/

.. _Apache: http://httpd.apache.org/

.. _mod_wsgi: https://code.google.com/p/modwsgi/

.. _Google App Engine Python: https://cloud.google.com/appengine/docs/python/

.. _third-party libraries in GAE Python 2.7: https://cloud.google.com/appengine/docs/python/tools/libraries27

.. _download web.py from Github: https://github.com/webpy/webpy
