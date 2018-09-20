[Golang] Read Plain Texts in PDF Files
######################################

:date: 2018-09-21T03:08+08:00
:tags: Go, Golang
:category: Go
:summary: Extract plain text from PDF via *github.com/ledongthuc/pdf* package.
:og_image: https://unidoc.io/static/assets/unidoc2x.png
:adsu: yes


I read the thread about extracting text from PDF on Reddit. [1]_ It's
interesting and I made some searches [2]_ and decided to try
`github.com/ledongthuc/pdf`_ package. The following code comes from modification
of sample code of the package.

.. show_github_file:: siongui userpages content/code/go/read-pdf-text/readpdftext.go

**Usage**

.. show_github_file:: siongui userpages content/code/go/read-pdf-text/readpdftext_test.go

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 18.04``, ``Go 1.11``

**References**

.. [1] | `Best text extractor from PDFs without OCR : golang <https://old.reddit.com/r/golang/comments/9fwgjy/best_text_extractor_from_pdfs_without_ocr/>`_

.. [2] | `go - Extract words from PDF with golang? - Stack Overflow <https://stackoverflow.com/questions/39813890/extract-words-from-pdf-with-golang>`_
       | `go - How to extract plain text from PDF in golang - Stack Overflow <https://stackoverflow.com/questions/44560265/how-to-extract-plain-text-from-pdf-in-golang>`_
.. [3] `GitHub - ledongthuc/pdf: PDF reader <https://github.com/ledongthuc/pdf>`_

.. _github.com/ledongthuc/pdf: https://github.com/ledongthuc/pdf
