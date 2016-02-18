Deploy Website by Pelican, Travis CI, and GitHub Pages
######################################################

:date: 2016-01-05T03:53+08:00
:tags: Continuous Integration, Pelican
:category: Web Development
:summary: Automatic deployment of websites, by Pelican_, `Travis CI`_, and
          `Github Pages`_


I use two separate repositories for the source and the built website. One is
``userpages`` (source of the website) and the other is ``siongui.github.io``
(built website).

Build Website
+++++++++++++

First add ``.travis.yml`` [11]_ in the root of your source repo. My
configuration is as follows:

.. code-block:: yaml

  language: python
  python:
  - '2.7'
  branches:
    only:
    - master
  addons:
    apt:
      packages:
      - language-pack-en
      - language-pack-zh-hant
      - language-pack-th
  install:
  - pip install -r requirements.txt
  - make download
  script:
  - make publish

My website support threes languages (English, Traditional Chinese, Thai), and I
need three locales (en_US.UTF8, zh_TW.UTF-8, th_TH.UTF-8) to build each
subsites. That's why three language-packs are installed.

Deploy Website to `Github Pages`_
+++++++++++++++++++++++++++++++++

Generate access token
`````````````````````

Next, we need a access token [12]_ to authorize Travis CI to push build to
`Github Pages`_. Generate a `personal access token`_ in Personal Settings (only
*public_repo* privilege is enough).

Encrypt token
`````````````

Then we will encrypt this token and put in our ``.travis.yml``: We need
`Travis CI command line client`_ to encrypt the token. In Ubuntu Linux 15.10,
you can install the command line client by:

.. code-block:: bash

  $ sudo apt-get install ruby ruby-dev
  $ sudo gem install travis

After finish, encrypt your token by:

.. code-block:: bash

  $ travis encrypt GH_TOKEN=your_token

Auto-add the token to your ``.travis.yml``:

.. code-block:: bash

  $ travis encrypt GH_TOKEN=your_token --add

Then you will see something like this in your ``.travis.yml``:

.. code-block:: yaml

  env:
    global:
      secure: Ph34LjX71tay9yVAYcg0RiwcLCeu/FthMkUtUoNf8pnI+OOL1MkW305jEf6F2oqObnRDuMvowppUfd/4LzqkSBat3CMr9htQzIQ+PNB5mIhx9JTYjkmz31iVoIBEPOFadSwvUDwKbcoJfAzmgrsX5wC0MadeD2iyUQMQ1BekT8k+HXqAp5WtmqjD0KcnFnlhXFzWl+fDiMoadVH85wUn+z827Tx+aH2XP1C9q8BhM8KeJXjTqacSzrLxOs+ElX603s1o2gSG9pxe+FPhABIVGPx/44V67zthwxs5j4d+WtXCWliK+ah+H/gQ2HxWydKLwfKoUI5BeowhUx9QHmodxjaWomL9ciqSPa0b9TpbeyzypzErecoZrFC8RZeXl1B2TDANROXGo5nTSKziDs1/Ajj83bNLoJ9K1B1yELpkHHi1mEoXb68Z6d9pit9uYsrA78Atkn/kAVrfS8/+fPuX+oeGk4MFF63epMxJm+otqWW4yjUTDg9vqVOe9eOgoD5sKWve5ZX8+ELIPHwg7xA5h+03nMqr3Bhe0YcmxIX9MHeZh49aub+Dm5H6q/rXWi8pwftW304BN/dO5gDk5EOd6F7ufOGvIM9Hj6iS1PYpJATw558Bkf/RX17kNhpAPhy9S1eswR34ET0ogvfjsLvMpQa/LvKcj8JTWj7OD+ilMoE=

Push build to GitHub Pages
``````````````````````````

I add the following command in my ``.travis.yml`` to push the build to GitHub
Pages:

.. code-block:: yaml

  after_success:
  - git config --global user.email "travis@travis-ci.org"
  - git config --global user.name "Travis"
  - git config --global push.default simple
  - echo "Clone siongui.github.io..."
  - git clone --depth 1 https://${GH_TOKEN}@github.com/siongui/siongui.github.io.git > /dev/null
  - cd siongui.github.io
  - git rm -rf * > /dev/null
  - cp -r  ../output/* .
  - git add .
  - git status
  - git commit -m "Pushed by Travis CI"
  - git push --quiet

You can modify the above commands to fit your need. My development machine is
``Ubuntu Linux 15.10``. For more information of how to deploy the website, see
[5]_, [6]_, [7]_, [8]_, and [9]_.

----

References:

.. [1] Google Search: `pelican travis <https://www.google.com/search?q=pelican+travis>`__

.. [2] `Pelican Static Site Generator, Powered by Python <http://blog.getpelican.com/>`_

.. [3] `Travis CI - Test and Deploy Your Code with Confidence <https://travis-ci.org/>`_

.. [4] `GitHub Pages <https://pages.github.com/>`_

.. [5] `Publish your Pelican blog on Github pages via Travis-CI <http://blog.mathieu-leplatre.info/publish-your-pelican-blog-on-github-pages-via-travis-ci.html>`_

.. [6] `How to automatically build your Pelican blog and publish it to Github Pages | Andrea Zonca's blog <http://zonca.github.io/2013/09/automatically-build-pelican-and-publish-to-github-pages.html>`_

.. [7] `用 Travis-CI 生成 Github Pages 博客 - Farseerfc的小窩 <https://farseerfc.me/travis-push-to-github-pages-blog.html>`_

.. [8] `A static site with Pelican, Grunt, Travis & Github Pages <http://iamemmanouil.com/blog/static-site-pelican-grunt-travis-github-pages/>`_

.. [9] `Deploying my blog with Travis <http://www.andrewaitken.com/2014/04/deploying-my-blog-with-travis/>`_

.. [10] DuckDuckGo Search: `pelican travis <https://duckduckgo.com/?q=pelican+travis>`__

.. [11] `Customizing the Build - Travis CI <https://docs.travis-ci.com/user/customizing-the-build/>`_

.. [12] `Creating an access token for command-line use - GitHub Help <https://help.github.com/articles/creating-an-access-token-for-command-line-use/>`_


.. _Pelican: http://blog.getpelican.com/
.. _Travis CI: https://travis-ci.org/
.. _GitHub Pages: https://pages.github.com/
.. _personal access token: https://help.github.com/articles/creating-an-access-token-for-command-line-use/
.. _Travis CI command line client: https://github.com/travis-ci/travis.rb
