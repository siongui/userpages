Personal Blog
=============

Usage
-----

.. code-block:: bash

  # install virtualenv (use Ubuntu 13.10 as example)
  $ sudo apt-get install python-pip
  $ sudo pip install virtualenv

  # # create isolated Python environment
  $ virtualenv dev
  # enter isolated Python environment
  $ cd dev
  $ source bin/activate
  # git clone this repository.
  $ git clone https://github.com/siongui/userpages.git
  # install packages for building the blog
  $ pip install -r requirements.txt


Non-Standard Pelican Config
---------------------------

External Libraries URL
~~~~~~~~~~~~~~~~~~~~~~

* `FONT_AWESOME_URL`: for example, ``//netdna.bootstrapcdn.com/font-awesome/4.0.0/css/font-awesome.min.css``

Github Ribbon
~~~~~~~~~~~~~

* `GITHUB_REPO_URL`: for example, ``https://github.com/siongui/userpages``

HTML Meta Info
~~~~~~~~~~~~~~

* `META_KEYWORDS`: for example, ``Python, webapp``
* `META_DESCRIPTION`: for example, ``Template for Python Web Application``


UNLICENSE
---------

Released in public domain. Please see `UNLICENSE <http://unlicense.org/>`_.
