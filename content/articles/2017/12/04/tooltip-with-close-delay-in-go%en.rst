Tooltip with Close Delay in Go
##############################

:date: 2018-01-09T22:51+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go, CSS,
       Tooltip, mouseenter, mouseleave, element offset, DOM, element position
:category: Frontend Programming in Go
:summary: Show tooltip when the cursor hovers over the text, and close the
          tooltip with delay if the cursor is not in the tooltip
          via Go/GopherJS.
:og_image: http://freefrontend.com/assets/img/css-tooltips/automation-tooltips-with-simple-data-attributes.png
:adsu: yes


Show tooltip when the cursor hovers over the text, and close the tooltip with
delay if the cursor is not in the tooltip via Go/GopherJS.
This is actually the Go/GopherJS translation of my previous post
`[JavaScript] Tooltip with Close Delay`_ [1]_.
Please read my previous post and see the demo first.
Here I only show the Go code of the implementation. The HTML/CSS code are the
same as that in my previous post.
The full code example of this post is `on my GitHub`_.

.. code-block:: go

  package main

  import (
  	. "github.com/siongui/godom"
  	"strconv"
  	"time"
  )

  var tooltip *Object

  // indicate if the mouse cursor is in the tooltip
  var isCursorInTooltip = false

  // when cursor leaves the text, the delay time to close the tooltip if the
  // cursor is not in the tooltip. (milisecond)
  var delay = 250 * time.Millisecond

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

  	tooltip.AddEventListener("mouseenter", func(e Event) {
  		isCursorInTooltip = true
  	})
  	tooltip.AddEventListener("mouseleave", func(e Event) {
  		isCursorInTooltip = false
  		tooltip.ClassList().Add("invisible")
  	})

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
  	time.AfterFunc(delay, func() {
  		if !isCursorInTooltip {
  			tooltip.ClassList().Add("invisible")
  		}
  	})
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

.. [1] `[JavaScript] Tooltip with Close Delay <{filename}../../../2018/01/08/javascript-tooltip-with-close-delay%en.rst>`_

.. _godom: https://github.com/siongui/godom
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/019-tooltip-with-close-delay
.. _[JavaScript] Tooltip with Close Delay: {filename}../../../2018/01/08/javascript-tooltip-with-close-delay%en.rst
