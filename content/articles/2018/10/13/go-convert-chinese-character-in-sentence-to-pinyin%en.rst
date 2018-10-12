[Golang] Convert Chinese Characters in String to Pinyin
#######################################################

:date: 2018-10-13T04:44+08:00
:tags: Go, Golang, String Manipulation, Locale, i18n, Chinese Pinyin
:category: Go
:summary: Given a string with Chinese characters, convert the Chinese characters
          in the string to Pinyin (romanization system for Chinese) in Go.
:og_image: https://sfschinese.weebly.com/uploads/5/0/6/2/5062697/pinyin_chart__948734.jpg
:adsu: yes


Given a string with Chinese characters, convert the Chinese characters in the
string to Pinyin_ (romanization system for Chinese) in Go.

For example,

.. code-block:: txt

  Hello 世界

will become the following text after conversion.

.. code-block:: txt

  Hello shijie

To do this, we need `go-pinyin`_ package to convert Chinese characters to
Pinyin. Install the package first:

.. code-block:: bash

  $ go get -u github.com/mozillazg/go-pinyin

Now we can convert Chinese characters to Pinyin in the string as follows:

.. code-block:: go

  import (
  	"unicode"

  	"github.com/mozillazg/go-pinyin"
  )

  var a = pinyin.NewArgs()

  func zhCharToPinyin(p string) (s string) {
  	for _, r := range p {
  		if unicode.Is(unicode.Han, r) {
  			s += string(pinyin.Pinyin(string(r), a)[0][0])
  		} else {
  			s += string(r)
  		}
  	}
  	return
  }

.. adsu:: 2

Example:

.. code-block:: go

  import (
  	"fmt"
  )

  func ExampleZhCharToPinyin() {
  	fmt.Println(zhCharToPinyin("Hello 世界"))
  	// Output: Hello shijie
  }

Tones can be included in Pinyin by configuring Args. See README of `go-pinyin`_
package for more details.

----

Tested on: ``Ubuntu 18.04``, ``Go 1.10.1``

----

References:

.. [1] `[Golang] First Letter of Chinese Character Pinyin <{filename}/articles/2017/05/05/go-chinese-character-pinyin-first-letter%en.rst>`_
.. [2] `[Golang] Check If The Rune is Chinese Character <{filename}/articles/2018/10/11/go-check-if-rune-value-is-chinese-character%en.rst>`_

.. _Pinyin: https://en.wikipedia.org/wiki/Pinyin
.. _go-pinyin: https://github.com/mozillazg/go-pinyin
