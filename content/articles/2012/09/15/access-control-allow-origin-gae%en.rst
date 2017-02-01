Access-Control-Allow-Origin in HTTP Header on Google App Engine
###############################################################

:date: 2012-09-15 23:18
:modified: 2015-02-22 10:53
:tags: HTTP Header, Google App Engine, CORS
:category: Google App Engine
:summary: Config Access-Control-Allow-Origin in HTTP headers on Google App Engine.
:adsu: yes


To add *Access-Control-Allow-Origin* in HTTP headers on `Google App Engine`_,
many search results said that we CAN NOT add it. But actually we CAN add it now.
The following is an example of how I add *Access-Control-Allow-Origin* in my
*app.yaml* file (`Google App Engine for Python`_, the same for
`Google App Engine for Go`_).

.. code-block:: yaml

  - url: /translation
    static_dir: static/translation
    http_headers:
      Access-Control-Allow-Origin: "*"

For Python & Go config details, please refer to reference [1]_ and [2]_.

.. adsu:: 2

For `Google App Engine for Java`_ user, the following is an example from the
official doc:

.. code-block:: xml

  <static-files>
    <include path="/my_static-files" >
      <http-header name="Access-Control-Allow-Origin" value="http://example.org" />
    </include>
  </static-files>

For Java config details, please refer to reference [3]_.

----

.. adsu:: 3

References:

.. [1] `CORS Support - Configuring with app.yaml - Python <https://cloud.google.com/appengine/docs/python/config/appconfig#cors_support>`_

.. [2] `CORS Support - Configuring with app.yaml - Go <https://cloud.google.com/appengine/docs/go/config/appconfig#cors_support>`_

.. [3] `Configuring appengine-web.xml - Java <https://cloud.google.com/appengine/docs/java/config/appconfig#Java_appengine_web_xml_Including_and_excluding_files>`_

.. [4] `Cross-origin resource sharing <http://en.wikipedia.org/wiki/Cross-origin_resource_sharing>`_


.. _Google App Engine: https://cloud.google.com/appengine/docs

.. _Google App Engine for Python: https://cloud.google.com/appengine/docs/python/

.. _Google App Engine for Go: https://cloud.google.com/appengine/docs/go/

.. _Google App Engine for Java: https://cloud.google.com/appengine/docs/java/
