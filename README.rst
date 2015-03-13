===============
How to Develope
===============

Development Tool: `pelican <http://blog.getpelican.com/>`_ (static site generator written in Python)

Development Environment: `Ubuntu 14.10 <http://releases.ubuntu.com/14.10/>`_


First-time Setup
----------------

1. Install `git <http://git-scm.com/>`_, `pip <https://pypi.python.org/pypi/pip>`_, and `virtualenv <http://docs.python-guide.org/en/latest/dev/virtualenvs/>`_:

.. code-block:: bash

    $ sudo apt-get install git
    $ sudo apt-get install python-pip
    $ sudo pip install virtualenv

2. Create and enter Python virtual environment:

.. code-block:: bash

    $ cd
    $ mkdir dev
    $ virtualenv ~/dev/
    $ cd ~/dev/
    $ source bin/activate

3. git clone source code:

.. code-block:: bash

    $ cd ~/dev/
    $ git clone https://github.com/siongui/userpages.git

4. Install Python tools:

.. code-block:: bash

    $ cd ~/dev/userpages/
    $ sudo pip install -r requirements.txt

5. Install pelican `i18n_subsites <https://github.com/getpelican/pelican-plugins/tree/master/i18n_subsites>`_ plugin and download `normalize.css <http://necolas.github.io/normalize.css/>`_:

.. code-block:: bash

    $ cd ~/dev/userpages/
    $ make download

6. Generate CSS file:

.. code-block:: bash

    $ cd ~/dev/userpages/
    $ make scss


Daily Development
-----------------

.. code-block:: bash

    $ cd ~/dev/
    $ source bin/activate
    $ cd userpages/
    # start edit and develope
    # If something changes
    $ make html
    # start dev server at http://localhost:8000/
    $ make serve


UNLICENSE
---------

Released in public domain. Please see `UNLICENSE <http://unlicense.org/>`_.


References
----------

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
