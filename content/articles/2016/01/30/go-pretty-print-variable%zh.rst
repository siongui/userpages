[Go語言] 優美列印變數 (struct, map, 陣列, 切片)
###############################################

:date: 2018-10-06T21:25+08:00
:tags: 優美列印, Go語言, JSON
:category: Go語言
:summary: Go語言 優美列印變數 (struct, map, 陣列, 切片)。
:og_image: https://cdn-images-1.medium.com/max/600/1*i2skbfmDsHayHhqPfwt6pA.png
:adsu: yes

Go語言 優美列印變數 (struct, map, 陣列, 切片)。
最簡單的方法是透過 Go 語言標準函示庫裡 `encoding/json`_ 包的 MarshalIndent_
方法。

.. code-block:: go

  import (
  	"encoding/json"
  	"fmt"
  )

  func PrettyPrint(v interface{}) (err error) {
  	b, err := json.MarshalIndent(v, "", "  ")
  	if err == nil {
  		fmt.Println(string(b))
  	}
  	return
  }

.. rubric:: `在Go Playground執行程式碼 <https://play.golang.org/p/DfF_H_9uLJ_z>`__
   :class: align-center

測試環境： `Go Playground`_

.. adsu:: 2

.. _MarshalIndent: https://golang.org/pkg/encoding/json/#MarshalIndent
.. _encoding/json: https://golang.org/pkg/encoding/json/
.. _Go Playground: https://play.golang.org/
