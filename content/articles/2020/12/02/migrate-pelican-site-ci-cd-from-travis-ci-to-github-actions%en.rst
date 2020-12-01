Migrate Pelican Site CI/CD From Travis CI to GitHub Actions
###########################################################

:date: 2020-12-02T03:41+08:00
:tags: Continuous Integration, Pelican, GitHub Pages, GitHub Actions, Travis CI
:category: Continuous Integration/Continuous Delivery (CI/CD)
:summary: Use GitHub Actions for Pelican Site CI/CD instead of Travis CI.
:og_image: https://nolanbconaway.github.io/pelican-deploy-gh-actions/images/workflow.png
:adsu: yes


Build and deploy Pelican site using GitHub Actions instead of Travis CI.

My original CI/CD file on Travis CI is:

.. code-block:: yaml

  sudo: required
  dist: bionic
  language: python
  python:
  - '3.8'
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
  deploy:
    provider: pages
    repo: siongui/siongui.github.io
    target_branch: master
    skip_cleanup: true
    github_token: $GITHUB_TOKEN
    local_dir: output
    on:
      branch: master

Summaries of above YAML config:

- Ubuntu machine/Python 3.8
- Install Ubuntu en/zh-hant/th language packs
- Install Pelican and Python packages using pip
- Build Pelican site using Makefile
- Deploy to another repo of GitHub Pages using Personal Access Token.

The following is the corresponding CI/CD file on GitHub Actions:

.. code-block:: yaml

  name: Pelican site CI

  on:
    push:
      branches:
        - master

  jobs:
    build:
      runs-on: ubuntu-latest
      steps:
      - name: Install language packs
        run: sudo apt-get install language-pack-en language-pack-zh-hant language-pack-th
      - uses: actions/checkout@v2.3.1
        with:
          persist-credentials: false
      - uses: actions/setup-python@v2
        with:
          python-version: '3.8'
      - name: Install dependencies
        run: |
          python -m pip install --upgrade pip
          pip install -r requirements.txt
      - name: Build site
        run: |
          make download
          make publish
      - name: Add nojekyll
        run: |
          touch ./output/.nojekyll
      - name: Deploy
        uses: JamesIves/github-pages-deploy-action@3.7.1
        with:
          ACCESS_TOKEN: ${{ secrets.ACCESS_TOKEN }}
          REPOSITORY_NAME: siongui/siongui.github.io
          BRANCH: master
          FOLDER: output
          CLEAN: true


Tested on: ``Ubuntu Linux 20.04``, ``Python 3.8.5``, ``GitHub Actions``.

----

References:

.. [1] | `pelican github actions - Google search <https://www.google.com/search?q=pelican+github+actions>`_
       | `pelican github actions - DuckDuckGo search <https://duckduckgo.com/?q=pelican+github+actions>`_
       | `pelican github actions - Ecosia search <https://www.ecosia.org/search?q=pelican+github+actions>`_
       | `pelican github actions - Qwant search <https://www.qwant.com/?q=pelican+github+actions>`_
       | `pelican github actions - Bing search <https://www.bing.com/search?q=pelican+github+actions>`_
       | `pelican github actions - Yahoo search <https://search.yahoo.com/search?p=pelican+github+actions>`_
       | `pelican github actions - Baidu search <https://www.baidu.com/s?wd=pelican+github+actions>`_
       | `pelican github actions - Yandex search <https://www.yandex.com/search/?text=pelican+github+actions>`_

.. [2] `Pelican and GitHub Pages <https://nimbinatus.com/2019/09/28/pelican-and-ghpages/>`_

.. _Pelican: https://github.com/getpelican/pelican
