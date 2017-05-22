Deploy Website by Pelican, Travis CI, and GitHub Pages
######################################################

:date: 2016-01-05T03:53+08:00
:modified: 2017-05-23T04:25+08:00
:tags: Continuous Integration, Pelican, GitHub Pages, apt command
:category: Web Development
:summary: Automatic deployment of websites, by Pelican_, `Travis CI`_, and
          `Github Pages`_
:adsu: yes


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

.. adsu:: 2

Deploy Website to `Github Pages`_
+++++++++++++++++++++++++++++++++

Now Travis CI officially support deploy your static files to GitHub Pages after
a successful build. [13]_

First save your `personal access token`_ in `repository settings`_.

For User Pages, the following is sample config:

.. code-block:: yaml

  deploy:
    provider: pages
    repo: siongui/siongui.github.io
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

.. [13] `GitHub Pages Deployment - Travis CI <https://docs.travis-ci.com/user/deployment/pages/>`_

.. _Pelican: http://blog.getpelican.com/
.. _Travis CI: https://travis-ci.org/
.. _GitHub Pages: https://pages.github.com/
.. _personal access token: https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
.. _Travis CI command line client: https://github.com/travis-ci/travis.rb
.. _repository settings: https://docs.travis-ci.com/user/environment-variables#Defining-Variables-in-Repository-Settings
