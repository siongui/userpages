[Golang] First Letter of Chinese Character Pinyin
#################################################

:date: 2017-05-05T00:22+08:00
:tags: Go, Golang, String Manipulation, Chinese Pinyin
:category: Go
:summary: My answer to How can I get Chinese first letter(Pinyin) by using Go
          language? - Stack Overflow
:og_image: https://s-media-cache-ak0.pinimg.com/originals/20/09/0e/20090eb8ff9dee87372f499b3e4f61f8.jpg
:adsu: yes


My answer to *How can I get Chinese first letter(Pinyin) by using Go language?
- Stack Overflow* [4]_

.. code-block:: go

  package first

  import (
  	"github.com/mozillazg/go-pinyin"
  )

  var a = pinyin.NewArgs()

  func FirstLetterOfPinYin(r rune) string {
  	result := pinyin.Pinyin(string(r), a)
  	return string(result[0][0][0])
  }

Testing:

.. code-block:: go

  package first

  import (
  	"testing"
  )

  func TestFirstLetterOfPinYin(t *testing.T) {
  	if FirstLetterOfPinYin('世') != "s" {
  		t.Error("世")
  	}
  	if FirstLetterOfPinYin('界') != "j" {
  		t.Error("界")
  	}
  }

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.04``, ``Go 1.8.1``.

----

References:

.. [1] | `golang chinese pinyin - Google search <https://www.google.com/search?q=golang+chinese+pinyin>`_
       | `golang chinese pinyin - DuckDuckGo search <https://duckduckgo.com/?q=golang+chinese+pinyin>`_
       | `golang chinese pinyin - Ecosia search <https://www.ecosia.org/search?q=golang+chinese+pinyin>`_
       | `golang chinese pinyin - Qwant search <https://www.qwant.com/?q=golang+chinese+pinyin>`_
       | `golang chinese pinyin - Bing search <https://www.bing.com/search?q=golang+chinese+pinyin>`_
       | `golang chinese pinyin - Yahoo search <https://search.yahoo.com/search?p=golang+chinese+pinyin>`_
       | `golang chinese pinyin - Baidu search <https://www.baidu.com/s?wd=golang+chinese+pinyin>`_
       | `golang chinese pinyin - Yandex search <https://www.yandex.com/search/?text=golang+chinese+pinyin>`_
.. [2] `pinyin - Go libraries and apps <https://golanglibs.com/top?q=pinyin>`_
.. [3] `GitHub - mozillazg/go-pinyin: 汉字拼音转换工具 Go 版 <https://github.com/mozillazg/go-pinyin>`_
.. [4] `How can I get Chinese first letter(Pinyin) by using Go language? - Stack Overflow <http://stackoverflow.com/questions/32109918/how-can-i-get-chinese-first-letterpinyin-by-using-go-language>`_
.. [5] `[Golang] Iterate Over UTF-8 Strings (non-ASCII strings) <{filename}../../../2016/02/03/go-iterate-over-utf8-non-ascii-string%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _for: https://tour.golang.org/flowcontrol/1
.. _range: https://github.com/golang/go/wiki/Range
