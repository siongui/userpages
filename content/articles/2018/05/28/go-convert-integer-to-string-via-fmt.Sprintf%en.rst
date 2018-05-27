[Golang] Convert Integer to String via fmt.Sprintf
##################################################

:date: 2018-05-28T00:47+08:00
:tags: Go, Golang, Type Casting, Type Conversion
:category: Go
:summary: Convert *int* or *int64* to *string* in Go via *fmt.Sprintf*.
:og_image: http://2.bp.blogspot.com/-O2uOFWMD_zs/Ve3Iq9rfX_I/AAAAAAAAPxE/ZKRsGxscEsY/s1600/go_lang_mascot_by_kirael_art-d7kunhu.gif
:adsu: yes


If you use strconv_ package, you can convert integer to string using following
methods:

- *int* to *string*: strconv.Itoa_
- *int64* to *string*: strconv.FormatInt_

There is another way to do this, that is, use fmt.Sprintf_

For example, assume variable **i** is either *int* or *int64* type, you can
convert it to *string* type as follows:

.. code-block:: go

  s := fmt.Sprintf("%d", i)

This is a small trick I found recently. You can try it yourself online:

`Run Example on Go Playground <https://play.golang.org/p/oQ4ilZRxW2B>`__

.. adsu:: 2


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _strconv: https://golang.org/pkg/strconv/
.. _strconv.Itoa: https://golang.org/pkg/strconv/#Itoa
.. _strconv.FormatInt: https://golang.org/pkg/strconv/#FormatInt
.. _fmt.Sprintf: https://golang.org/pkg/fmt/#Sprintf
