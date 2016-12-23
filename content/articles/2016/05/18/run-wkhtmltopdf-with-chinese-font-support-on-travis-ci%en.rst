Run wkhtmltopdf With Chinese Font Support on Travis CI
######################################################

:date: 2016-05-18T04:55+08:00
:tags: Bash, Commandline, find command, List Files in Directory, wkhtmltopdf,
       Continuous Integration, Ubuntu Linux, apt command
:category: Bash
:summary: Run wkhtmltopdf_ on `Travis CI`_, which convert HTML_ files with
          chinese characters to PDF_.


Run wkhtmltopdf_ on `Travis CI`_, which convert HTML_ files with chinese
characters to PDF_.

We need to overcome the following two problems:

1. error message: ``wkhtmltopdf: cannot connect to X server``

   The solution comes from `this answer on Stack Overflow`_.
   Add the following lines in your ``.travis.yml``

   .. code-block:: yaml

     sudo: required
     dist: trusty
     before_install:
     - sudo apt-get -qq update
     - sudo apt-get install -y xfonts-75dpi
     - wget http://download.gna.org/wkhtmltopdf/0.12/0.12.2/wkhtmltox-0.12.2_linux-trusty-amd64.deb
     - sudo dpkg -i wkhtmltox-0.12.2_linux-trusty-amd64.deb


2. wkhtmltopdf Chinese character support

   The solution comes from `this blog post`_.
   Add the following lines in your ``.travis.yml``

   .. code-block:: yaml

     before_install:
     - sudo apt-get install -y language-pack-zh-hant fonts-wqy-microhei ttf-wqy-microhei fonts-wqy-zenhei ttf-wqy-zenhei
     install:
     - fc-cache -f -v


Combine the above two solution together, your final ``.travis.yml`` will looks
like:

.. code-block:: yaml

  sudo: required
  dist: trusty
  before_install:
  - sudo apt-get -qq update
  - sudo apt-get install -y language-pack-zh-hant fonts-wqy-microhei ttf-wqy-microhei fonts-wqy-zenhei ttf-wqy-zenhei
  - sudo apt-get install -y xfonts-75dpi
  - wget http://download.gna.org/wkhtmltopdf/0.12/0.12.2/wkhtmltox-0.12.2_linux-trusty-amd64.deb
  - sudo dpkg -i wkhtmltox-0.12.2_linux-trusty-amd64.deb
  install:
  - fc-cache -f -v

----

Example Bash_ script to run wkhtmltopdf_:

.. code-block:: sh

  #!/bin/bash

  # $1 is the directory in which files to be processed
  # $2 is the css file
  for path in $(find $1 -name "*.html")
  do
    echo -e "\033[92mProcessing ${path}\033[0m"
    wkhtmltopdf ${path} --disable-javascript --user-style-sheet $2 "${path%.html}.pdf"
  done

----

Tested on: ``Travis CI - The Trusty beta Build Environment``, ``http://download.gna.org/wkhtmltopdf/0.12/0.12.2/wkhtmltox-0.12.2_linux-trusty-amd64.deb``

----

References:

.. [1] `wkhtmltopdf: cannot connect to X server - Google search <https://www.google.com/search?q=wkhtmltopdf:+cannot+connect+to+X+server>`_

       `xserver - wkhtmltopdf: cannot connect to X server - Stack Overflow <http://stackoverflow.com/questions/9604625/wkhtmltopdf-cannot-connect-to-x-server>`_

.. [2] `wkhtmltopdf chinese font - Google search <https://www.google.com/search?q=wkhtmltopdf+chinese+font>`_

       `pdf - wkhtmltopdf and chinese characters - Stack Overflow <http://stackoverflow.com/questions/25833954/wkhtmltopdf-and-chinese-characters>`_

       `wkHTMLToPDF chinese character support on Linux based systems | Clement Nedelcu's Development Journal <https://cnedelcu.blogspot.com/2015/04/wkhtmltopdf-chinese-character-support.html>`_

.. [3] `Installing Dependencies - Travis CI <https://docs.travis-ci.com/user/installing-dependencies/>`_

       `The Build Environment - Travis CI <https://docs.travis-ci.com/user/ci-environment/>`_

       `The Trusty beta Build Environment - Travis CI <https://docs.travis-ci.com/user/trusty-ci-environment/>`_

.. [4] `[Bash] HTML to PDF via wkhtmltopdf <{filename}../17/bash-html-to-pdf-via-wkhtmltopdf%en.rst>`_


.. _Bash: https://www.google.com/search?q=Bash
.. _HTML: https://www.google.com/search?q=HTML
.. _PDF: https://www.google.com/search?q=PDF
.. _wkhtmltopdf: http://wkhtmltopdf.org/
.. _install wkhtmltopdf: https://www.google.com/search?q=install+wkhtmltopdf
.. _Travis CI: https://travis-ci.org/
.. _this answer on Stack Overflow: http://stackoverflow.com/a/34947479
.. _this blog post: https://cnedelcu.blogspot.com/2015/04/wkhtmltopdf-chinese-character-support.html
