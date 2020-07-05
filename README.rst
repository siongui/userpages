===============
How to Develope
===============

.. image:: https://travis-ci.org/siongui/userpages.svg?branch=master
    :target: https://travis-ci.org/siongui/userpages

.. image:: https://gitlab.com/siongui/userpages/badges/master/pipeline.svg
    :target: https://gitlab.com/siongui/userpages/-/commits/master

.. See how to add travis ci image from https://raw.githubusercontent.com/demizer/go-rst/master/README.rst
   https://github.com/demizer/go-rst/commit/9651ab7b5acc997ea2751845af9f2d6efee825df

Development Tool: Pelican_ (static site generator written in Python)

Development Environment: `Ubuntu 20.04`_


First-time Setup
----------------

1. On a fresh/clean installation of Ubuntu, update system first. Otherwise will
   get unable to locate package error.
   See `this SO answer <https://stackoverflow.com/a/58072486>`__.

   .. code-block:: bash

     $ sudo apt-get update

2. Install git_ and pip_:

   .. code-block:: bash

     $ sudo apt-get install git
     $ sudo apt-get install python3-pip

   From the `answer in Ask Ubuntu <https://askubuntu.com/a/1031733>`_,
   we can use python-is-python3 and prevent Python 2 from being installed
   on Ubuntu 20.04

   .. code-block:: bash

     $ sudo apt-get install python-is-python3
     $ sudo apt-mark hold python2 python2-minimal python2.7 python2.7-minimal libpython2-stdlib libpython2.7-minimal libpython2.7-stdlib

3. Install language packages to add locale (English, Traditional Chinese, and
   Thai):

   .. code-block:: bash

     $ sudo apt-get install language-pack-en
     $ sudo apt-get install language-pack-zh-hant
     $ sudo apt-get install language-pack-th

   Or you can install languages in "Settings" -> "Region & Language", which
   installs more related packages such as fonts for languages.

4. git clone source code:

   .. code-block:: bash

     $ cd
     $ mkdir dev
     $ cd ~/dev/
     $ git clone https://github.com/siongui/userpages.git --depth=1
     # or clone with full depth
     $ git clone https://github.com/siongui/userpages.git

5. Install Python tools:

   .. code-block:: bash

     $ cd ~/dev/userpages/
     $ pip3 install -r requirements.txt

   Note that in `.travis.yml <.travis.yml>`_, pip is actually pip3 if bionic and
   python 3.8 is set in Travis CI config.

6. Install Pelican `i18n_subsites`_ plugin and download `normalize.css`_:

   .. code-block:: bash

     $ cd ~/dev/userpages/
     $ make download

7. Generate CSS file:

   .. code-block:: bash

     $ cd ~/dev/userpages/
     $ make scss


Daily Development
-----------------

.. code-block:: bash

    # start edit and develope
    $ cd ~/dev/userpages/
    # re-generate the website and start dev server
    $ make
    # open your browser and preview the website at http://localhost:8000/


Auto-deploy by `Travis CI`_
---------------------------

See `GitHub Pages Deployment - Travis CI`_.

First save your `personal access token`_ in `repository settings`_.

For User Pages, the following is sample config:

.. code-block:: yaml

  deploy:
    provider: pages
    repo: USERNAME/USERNAME.github.io
    target_branch: master
    skip_cleanup: true
    github_token: $GITHUB_TOKEN
    local_dir: output
    on:
      branch: master

For Project Pages, the following is sample config:

.. code-block:: yaml

  deploy:
    provider: pages
    skip_cleanup: true
    github_token: $GITHUB_TOKEN
    local_dir: output
    on:
      branch: master


Deploy to `GitLab Pages`_ via `GitLab CI/CD`_
---------------------------------------------

See `.gitlab-ci.yml <.gitlab-ci.yml>`_.

1. Use `Ubuntu image in Docker Hub <https://hub.docker.com/_/ubuntu>`_.
   The *ubuntu:latest* tag points to the "latest LTS".

2. Need to update Ubuntu first to install Ubuntu packages.
   See `this SO answer <https://stackoverflow.com/a/58072486>`__.

3. `Travis CI`_ can deploy to any repository, but need `personal access token`_
   to do so.
   `GitLab CI/CD`_ can deploy to the repository where CI/CD runs without
   credentials.


UNLICENSE
---------

All works, including posts and code, of Siong-Ui Te are released in public domain.
Please see UNLICENSE_.


References
----------

.. [1] `pelican-edit-url <https://github.com/pmclanahan/pelican-edit-url>`_
       inspires the *Edit on Github* link.

.. [2] | JINJA_FILTERS in `Settings — Pelican documentation <http://docs.getpelican.com/en/latest/settings.html>`_
       | `Jinja custom filters documentation <http://jinja.pocoo.org/docs/dev/api/#custom-filters>`_

.. [3] | Home Screen Icon on Android/iPhone & PWA support
       | `website icon on android home screen - Google search <https://www.google.com/search?q=website+icon+on+android+home+screen>`_
       | `Tutorial: Home Screen Icons | Responsive Web Design Training Tutorial | Webucator <https://www.webucator.com/tutorial/developing-mobile-websites/home-screen-icons.cfm>`_
       | `pwa manifest - Google search <https://www.google.com/search?q=pwa+manifest>`_
       | `WebPageTest - Website Performance and Optimization Test <https://www.webpagetest.org/>`_


.. _Pelican: https://blog.getpelican.com/
.. _Ubuntu 20.04: http://releases.ubuntu.com/20.04/
.. _git: https://git-scm.com/
.. _pip: https://pypi.python.org/pypi/pip
.. _i18n_subsites: https://github.com/getpelican/pelican-plugins/tree/master/i18n_subsites
.. _normalize.css: https://necolas.github.io/normalize.css/
.. _Travis CI: https://travis-ci.org/
.. _GitHub Pages Deployment - Travis CI: https://docs.travis-ci.com/user/deployment/pages/
.. _personal access token: https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token
.. _repository settings: https://docs.travis-ci.com/user/environment-variables#defining-variables-in-repository-settings
.. _GitLab Pages: https://docs.gitlab.com/ee/user/project/pages/
.. _GitLab CI/CD: https://docs.gitlab.com/ee/ci/
.. _Google Adsense: https://www.google.com/search?q=Google+AdSense
.. _UNLICENSE: https://unlicense.org/
