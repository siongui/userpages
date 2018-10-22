[Golang] Names scores - Problem 22 - Project Euler
##################################################

:date: 2018-10-23T01:10+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler, Go sort, String Manipulation,
       File Input/Output, Go net/http, Go strings Package
:category: Go
:summary: Go solution to Names scores
          - Problem 22 - Project Euler.
:og_image: https://2.bp.blogspot.com/-26hOIWw2rhs/VxXPEOn_DdI/AAAAAAAABxY/z6f21C74KPM3ypPThrDFPS7SLW5NL4J3gCLcB/s1600/project%2Beuler%2Bproblem%2B22%2Bwith%2Bsolution.png
:adsu: yes

**Problem**: [1]_

.. raw:: html

  <p>Using <a href="https://projecteuler.net/project/resources/p022_names.txt">names.txt</a> (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.</p>
  <p>For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.</p>
  <p>What is the total of all the name scores in the file?</p>

**Solution**:

  871198282

.. show_github_file:: siongui userpages content/code/go/project-euler/22.go

.. adsu:: 2

----

Test on:

- ``Ubuntu 18.04``, ``Go 1.11.1``

References:

.. [1] `Names scores - Problem 22 - Project Euler <https://projecteuler.net/problem=22>`_
.. [2] `[Golang] Sort Words Alphabetically <{filename}/articles/2018/04/10/go-sort-words-alphabetically%en.rst>`_
.. [3] `[Golang] Read Lines From URL <{filename}/articles/2017/02/02/go-readlines-from-url%en.rst>`_

.. _Go Playground: https://play.golang.org/
