[Go語言] 檢查元素是否在切片或陣列裡
###################################

:date: 2018-10-05T01:28+08:00
:tags: Go語言, 資料結構
:category: Go語言
:summary: Go語言 - 檢查某個元素是否在切片(slice)或陣列(array)裡
:og_image: https://research.swtch.com/godata3.png
:adsu: yes


**問題**

  在Go語言中，如何檢查某個元素是否在切片(slice)或陣列(array)裡？

**解法**

  有許多方式可以檢查 [1]_ 。 我的方法是先建立一個 *map* [2]_ 類型，
  並且把切片/陣列的所有元素存在 *map* 裡，然後再檢查某個元素是否在該 *map* 裡。

  下面是上述方法的反例。

.. rubric:: `在Go Playground執行程式碼 <https://play.golang.org/p/gk-otCgvwC>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func main() {
  	// initialize a slice of int
  	slice := []int{2, 3, 5, 7, 11, 13}

  	// save the items of slice in map
  	m := make(map[int]bool)
  	for i := 0; i < len(slice); i++ {
  		m[slice[i]] = true
  	}

  	// check if 5 is in slice by checking map
  	if _, ok := m[5]; ok {
  		fmt.Println("5 is in slice")
  	} else {
  		fmt.Println("5 is not in slice")
  	}

  	// check if 6 is in slice by checking map
  	if _, ok := m[6]; ok {
  		fmt.Println("6 is in slice")
  	} else {
  		fmt.Println("6 is not in slice")
  	}
  }

參考 [3]_ 是真實世界的應用。質數被傳回在切片(slice)裡。接著這些質數被存在map裡，
我們藉由檢查一個數值是否在該map裡來確認該數值是否為質數。

.. adsu:: 2

Tested on: `Go Playground`_

----

參考：

.. [1] | `golang check if item in slice - Google search <https://www.google.com/search?q=golang+check+if+item+in+slice>`_
       | `golang check if item in slice - DuckDuckGo search <https://duckduckgo.com/?q=golang+check+if+item+in+slice>`_
       | `golang check if item in slice - Ecosia search <https://www.ecosia.org/search?q=golang+check+if+item+in+slice>`_
       | `golang check if item in slice - Qwant search <https://www.qwant.com/?q=golang+check+if+item+in+slice>`_
       | `golang check if item in slice - Bing search <https://www.bing.com/search?q=golang+check+if+item+in+slice>`_
       | `golang check if item in slice - Yahoo search <https://search.yahoo.com/search?p=golang+check+if+item+in+slice>`_
       | `golang check if item in slice - Baidu search <https://www.baidu.com/s?wd=golang+check+if+item+in+slice>`_
       | `golang check if item in slice - Yandex search <https://www.yandex.com/search/?text=golang+check+if+item+in+slice>`_
.. [2] `Go maps in action - The Go Blog <https://blog.golang.org/go-maps-in-action>`_
.. [3] `[Golang] Goldbach's conjecture <{filename}../06/go-goldbach-conjecture%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _Goldbach's conjecture: https://www.google.com/search?q=Goldbach's+conjecture
