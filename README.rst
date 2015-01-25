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

5. Install pelican `i18n_subsites <https://github.com/getpelican/pelican-plugins/tree/master/i18n_subsites>`_ plugin:

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

`Can I use an image to trigger a modal window in Bootstrap? <http://stackoverflow.com/questions/15423532/can-i-use-an-image-to-trigger-a-modal-window-in-bootstrap>`_

`Bootstrap Image trigger modal example code <http://www.bootply.com/7wOLkC9AVX>`_

`iHover <http://gudh.github.io/ihover/dist/>`_ (`src <https://github.com/gudh/ihover>`_)

`bootstrap image hover overlay with icon <http://stackoverflow.com/questions/26823237/bootstrap-image-hover-overlay-with-icon>`_
