[Golang] Save File in QR Code
#############################

:date: 2018-02-24T11:45+08:00
:tags: Go, Golang, File Input/Output, base64
:category: Go
:summary: Embed file in QR code via Go programming language.
:og_image: https://foomandoonian.files.wordpress.com/2012/04/qart1.png
:adsu: yes


Recently I got an idea to embed file in the `QR code`_, so I tried to find Go
QR code encoder/decoder and started to do experiment. For encoder, I found
`github.com/skip2/go-qrcode`_. For decoder, I found `github.com/tuotoo/qrcode`_.

My idea is simple, first convert the file content/name into base64_ encoding,
and then save the content in the QR code. Because the limitation of content size
in single QR code image, you can only save small files. For large files, we need
to save the content in several QR code images. Since this is only experiment, I
use only small files and do not make any segmentation.

My code is as follows:

.. show_github_file:: siongui userpages content/code/go/file-qrcode/fileqrcode.go
.. adsu:: 2

It looks fine, but when I run the following test:

**Example**:

.. show_github_file:: siongui userpages content/code/go/file-qrcode/fileqrcode_test.go

The decoded string from the QR code image is not legal string. I tried to decode
the QR code image from `online QR code decoder`_, the decoded string is correct.
So the implementation of `github.com/tuotoo/qrcode`_ is buggy. Anyway, this is
an interesting experiment to embed files in the QR code.

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10``.

.. adsu:: 3

References:

.. [1] | `base1024 emoji encoding implemented in Go : golang <https://www.reddit.com/r/golang/comments/83kta3/base1024_emoji_encoding_implemented_in_go/>`_
       | `GitHub - keith-turner/ecoji: Encodes (and decodes) data as emojis <https://github.com/keith-turner/ecoji>`_
.. [2] `Animated QR data transfer with Gomobile and Gopherjs : golang <https://old.reddit.com/r/golang/comments/9yrk1g/animated_qr_data_transfer_with_gomobile_and/>`_

.. _QR code: https://www.google.com/search?q=QR+code
.. _github.com/skip2/go-qrcode: https://github.com/skip2/go-qrcode
.. _github.com/tuotoo/qrcode: https://github.com/tuotoo/qrcode
.. _base64: https://www.google.com/search?q=base64
.. _online QR code decoder: https://zxing.org/w/decode.jspx
