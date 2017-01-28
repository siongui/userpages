[Golang] Append Line/String to File
###################################

:date: 2017-01-23T06:26+08:00
:tags: Go, Golang, String Manipulation, File Input/Output
:category: Go
:summary: Append text line or string to the end of file
          via Go_ programming language.
:adsu: yes


I summarize the answer [1]_ [2]_ on `Stack Overflow`_ and write a func for the
ease of use.

.. code-block:: go

  import "os"

  /**
   * Append string to the end of file
   *
   * path: the path of the file
   * text: the string to be appended. If you want to append text line to file,
   *       put a newline '\n' at the of the string.
   */
  func AppendStringToFile(path, text string) error {
  	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
  	if err != nil {
  		return err
  	}
  	defer f.Close()

  	_, err = f.WriteString(text)
  	if err != nil {
  		return err
  	}
  	return nil
  }

.. adsu:: 2

----

Tested on:

- ``Ubuntu Linux 16.10``
- ``Go 1.7.4``

----

References:

.. [1] `Append to a file in Go - Stack Overflow <http://stackoverflow.com/a/7165157>`_

.. [2] `go - Golang bad file descriptor - Stack Overflow <http://stackoverflow.com/a/33852107>`_

.. adsu:: 3

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Stack Overflow: http://stackoverflow.com/
