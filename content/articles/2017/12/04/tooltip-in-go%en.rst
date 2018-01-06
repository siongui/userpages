Tooltip in Go
#############

:date: 2018-01-07T04:14+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go, CSS,
       Tooltip, mouseenter, mouseleave, element offset, DOM, element position
:category: Frontend Programming in Go
:summary: Simple tooltip implementation via Go/GopherJS.
:og_image: http://freefrontend.com/assets/img/css-tooltips/automation-tooltips-with-simple-data-attributes.png
:adsu: yes


Simple tooltip implementation in Go/GopherJS. This is actually translation of
my previous post `[JavaScript] Tooltip`_ [1]_. Please read my previous post and
see the demo first. Here I only show the Go code of the implementation. The
HTML/CSS code are the same as that in my previous post.
The full code example of this post is `on my GitHub`_.

.. code-block:: go

  package main

  import (
  	. "github.com/siongui/godom"
  	"strconv"
  )

  var tooltip *Object

  func GetPosition(elm *Object) (x, y float64) {
  	x = elm.GetBoundingClientRect().Left() + Window.PageXOffset()
  	y = elm.GetBoundingClientRect().Top() + Window.PageYOffset()
  	return
  }

  // Create and append invisible tooltip to DOM tree
  func SetupTooltip() {
  	tooltip = Document.CreateElement("div")
  	tooltip.ClassList().Add("tooltip")
  	tooltip.ClassList().Add("invisible")
  	Document.QuerySelector("body").AppendChild(tooltip)
  }

  // event handler for mouseenter event of elements with data-descr attribute
  func ShowTooltip(e Event) {
  	elm := e.Target()
  	tooltip.SetInnerHTML(elm.Dataset().Get("descr").String())

  	x, y := GetPosition(elm)
  	xpx := strconv.FormatFloat(x, 'E', -1, 64) + "px"
  	ypx := strconv.FormatFloat(y+elm.Get("OffsetHeight").Float()+5, 'E', -1, 64) + "px"
  	tooltip.Style().SetLeft(xpx)
  	tooltip.Style().SetTop(ypx)

  	tooltip.ClassList().Remove("invisible")
  }

  // event handler for mouseleave event of elements with data-descr attribute
  func HideTooltip(e Event) {
  	tooltip.ClassList().Add("invisible")
  }

  func main() {
  	SetupTooltip()

  	// select all elements with data-descr attribute and
  	// attach mouseenter and mouseleave event handler
  	els := Document.QuerySelectorAll("*[data-descr]")
  	for _, el := range els {
  		el.AddEventListener("mouseenter", ShowTooltip)
  		el.AddEventListener("mouseleave", HideTooltip)
  	}
  }

The above code use godom_ package to make the code more readable.

.. adsu:: 2

**Known Issue**:

The following line in above code:

.. code-block:: go

  ypx := strconv.FormatFloat(y+elm.Get("OffsetHeight").Float()+5, 'E', -1, 64) + "px"

The **+5** has no effect. If you remove it or change it to **+10**, the result
is all the same.

----

References:

.. [1] `[JavaScript] Tooltip <{filename}../../../2018/01/06/javascript-tooltip%en.rst>`_
.. [2] `Element Position (Scroll Included) in Go <{filename}element-position-scroll-included-in-go%en.rst>`_

.. _godom: https://github.com/siongui/godom
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/018-tooltip
.. _[JavaScript] Tooltip: {filename}../../../2018/01/06/javascript-tooltip%en.rst
