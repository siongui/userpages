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

embed Github repo file:

  `Gistfy - Gist your GitHub and Bitbucket repositories <http://www.gistfy.com/>`_

  `gist-it.appspot.com - Embed files from a github repository like a gist <http://gist-it.appspot.com/>`_

edit on Github link:

  `pelican-edit-url <https://github.com/pmclanahan/pelican-edit-url>`_

CSS layout:

  `vertical-align with bootstrap 3 <http://stackoverflow.com/questions/20547819/vertical-align-with-bootstrap-3>`_

  `How do you keep parents of floated elements from collapsing? <http://stackoverflow.com/questions/218760/how-do-you-keep-parents-of-floated-elements-from-collapsing>`_

CSS toggle element:

  `Is it possible to toggle div visibility with CSS? <http://www.reddit.com/r/css/comments/1f1nmm/is_it_possible_to_toggle_div_visibility_with_css/>`_

  `HTML5 Window Toggle Events In Pure CSS3 <http://demosthenes.info/blog/506/HTML5-Window-Toggle-Events-In-Pure-CSS3>`_

mobile responsive design:

  `Bootstrap 3 Slide in Menu / Navbar on Mobile <http://stackoverflow.com/questions/20863288/bootstrap-3-slide-in-menu-navbar-on-mobile>`_

  `Bootstrap Tutorial – Creating a Responsive Navbar (Video) <http://bootstrapbay.com/blog/bootstrap-tutorial-navbar/>`_

  `How to Create Off Canvas Layouts with Susy <http://www.zell-weekeat.com/off-canvas-layouts-susy/>`_

  `Off The Beaten Canvas: Exploring The Potential Of The Off-Canvas Pattern <http://www.smashingmagazine.com/2014/02/24/off-the-beaten-canvas-exploring-the-potential-of-the-off-canvas-pattern/>`_

  `Implementing Off-Canvas Navigation For A Responsive Website <http://www.smashingmagazine.com/2013/01/15/off-canvas-navigation-for-responsive-website/>`_

Bootstrap image trigger modal:

  `Can I use an image to trigger a modal window in Bootstrap? <http://stackoverflow.com/questions/15423532/can-i-use-an-image-to-trigger-a-modal-window-in-bootstrap>`_

  `Bootstrap Image trigger modal example code <http://www.bootply.com/7wOLkC9AVX>`_

Image Hover:

  `iHover <http://gudh.github.io/ihover/dist/>`_ (`src <https://github.com/gudh/ihover>`_)

  `bootstrap image hover overlay with icon <http://stackoverflow.com/questions/26823237/bootstrap-image-hover-overlay-with-icon>`_
