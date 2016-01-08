===============
How to Develope
===============

.. image:: https://travis-ci.org/siongui/userpages.png?branch=master
    :target: https://travis-ci.org/siongui/userpages

.. See how to add travis ci image from https://raw.githubusercontent.com/demizer/go-rst/master/README.rst
   https://github.com/demizer/go-rst/commit/9651ab7b5acc997ea2751845af9f2d6efee825df

Development Tool: `pelican <http://blog.getpelican.com/>`_ (static site generator written in Python)

Development Environment: `Ubuntu 15.10 <http://releases.ubuntu.com/15.10/>`_


First-time Setup
----------------

1. Install `git <http://git-scm.com/>`_ and `pip <https://pypi.python.org/pypi/pip>`_:

.. code-block:: bash

    $ sudo apt-get install git
    $ sudo apt-get install python-pip

2. git clone source code:

.. code-block:: bash

    $ cd
    $ mkdir dev
    $ cd ~/dev/
    $ git clone https://github.com/siongui/userpages.git

3. Install Python tools:

.. code-block:: bash

    $ cd ~/dev/userpages/
    $ sudo pip install -r requirements.txt

4. Install pelican `i18n_subsites <https://github.com/getpelican/pelican-plugins/tree/master/i18n_subsites>`_ plugin and download `normalize.css <http://necolas.github.io/normalize.css/>`_:

.. code-block:: bash

    $ cd ~/dev/userpages/
    $ make download

5. Generate CSS file:

.. code-block:: bash

    $ cd ~/dev/userpages/
    $ make scss

6. Install language packages to add locale (English, Traditional Chinese, and Thai):

.. code-block:: bash

    $ sudo apt-get install language-pack-en
    $ sudo apt-get install language-pack-zh-hant
    $ sudo apt-get install language-pack-th


Daily Development
-----------------

.. code-block:: bash

    # start edit and develope
    $ cd ~/dev/userpages/
    # If something changes, re-generate the website:
    $ make html
    # start dev server
    $ make serve
    # open your browser and preview the website at http://localhost:8000/


UNLICENSE
---------

All works, including posts and code, of Siong-Ui Te are released in public domain.
Please see `UNLICENSE <http://unlicense.org/>`_.


References
----------

`python - Upgrading all packages with pip - Stack Overflow <http://stackoverflow.com/questions/2720014/upgrading-all-packages-with-pip>`_

`How do I add locale to ubuntu server? - Ask Ubuntu <http://askubuntu.com/questions/76013/how-do-i-add-locale-to-ubuntu-server>`_

`Online reStructuredText editor <http://rst.ninjs.org/>`_

edit on Github link:

  `pelican-edit-url <https://github.com/pmclanahan/pelican-edit-url>`_

reStructuredText:

  `reStructuredText Markup Specification <http://docutils.sourceforge.net/docs/ref/rst/restructuredtext.html>`_

  `reStructuredText简明教程 <http://jwch.sdut.edu.cn/book/rst.html>`_

  `轻量级标记语言 <http://www.worldhello.net/gotgithub/appendix/markups.html>`_

  `reStructuredText 简明教程 <http://wstudio.web.fc2.com/others/restructuredtext.html>`_

  rst2html:

    `How can I get rst2html.py to include the CSS for syntax highlighting? <http://stackoverflow.com/questions/9807604/how-can-i-get-rst2html-py-to-include-the-css-for-syntax-highlighting>`_

    `Hottest 'rst2html.py' Answers - Stack Overflow <http://stackoverflow.com/tags/rst2html.py/hot>`_

    `html4css1.css <http://sourceforge.net/p/docutils/code/HEAD/tree/trunk/docutils/docutils/writers/html4css1/html4css1.css>`_

    rst2html stylesheet:

      `Writing HTML (CSS) Stylesheets for Docutils <http://docutils.sourceforge.net/docs/howto/html-stylesheets.html>`_

    rst2html css:

      `Documentation: Create GitHub like styled html doc file with rst2html <https://gist.github.com/vergissberlin/6422a0fe146c8fc04d7f>`_

      `marianoguerra/rst2html5 <https://github.com/marianoguerra/rst2html5>`_

      `How to render reStructuredText documents with latest docutils on Ubuntu 12.04 LTS <http://www.van-tomas.de/blog/restructuredtext-docutils-ubuntu-12-04-lts/>`_

      `[rsST] 修改 rst2html highlight style <http://blog.float.tw/2013/07/rst2html-change-highlight-style.html>`_

      `Docutils使用方式 <http://www.openfoundry.org/tw/download/doc_download/417-docutils-teachingdoc>`_ (`Google cache <http://www.openfoundry.org/tw/download/doc_download/417-docutils-teachingdoc>`__)

  restructuredtext center text:

    `Best way to align center a paragraph with RestructuredText? <http://stackoverflow.com/questions/14819093/best-way-to-align-center-a-paragraph-with-restructuredtext>`_

Bootstrap image trigger modal:

  `Can I use an image to trigger a modal window in Bootstrap? <http://stackoverflow.com/questions/15423532/can-i-use-an-image-to-trigger-a-modal-window-in-bootstrap>`_

  `Bootstrap Image trigger modal example code <http://www.bootply.com/7wOLkC9AVX>`_

Image Hover:

  `iHover <http://gudh.github.io/ihover/dist/>`_ (`src <https://github.com/gudh/ihover>`_)

  `bootstrap image hover overlay with icon <http://stackoverflow.com/questions/26823237/bootstrap-image-hover-overlay-with-icon>`_
