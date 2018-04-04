[Golang] Terminal Current Directory Item List via termui
########################################################

:date: 2018-04-04T23:39+08:00
:tags: Go, Golang, Terminal UI
:category: Go
:summary: Terminal UI - List items in directory via Go termui package.
:og_image: https://blog.nodejitsu.com/content/images/2014/Feb/blessed_list.png
:adsu: yes


.. image:: {filename}dirview.png
   :align: center
   :alt: view current directory items


Use Go terminal UI package - termui_ to list items in current directory. You can
press <up> or <down> key to select the item in the list, and press q to quit the
UI program. The text color of file is white and text color of directory is blue.

`ioutil.ReadDir`_ in Go standard package is used to read items in current
directory and termui list_ widget is used to provide the terminal UI.

The following is complete source code.

.. code-block:: go

  package main

  import (
  	"github.com/gizak/termui"
  	"io/ioutil"
  	"strings"
  )

  type ViewState struct {
  	ItemSelect int
  }

  func addColor(str, color string) string {
  	return "[" + str + "](" + color + ")"
  }

  func removeColor(str string) string {
  	if !strings.HasPrefix(str, "[") {
  		return str
  	}
  	if !strings.HasSuffix(str, ")") {
  		return str
  	}
  	if !strings.Contains(str, "](") {
  		return str
  	}
  	return strings.Split(str, "](")[0][1:]
  }

  func isDir(str string) bool {
  	if strings.Contains(str, "fg-blue") {
  		return true
  	}
  	return false
  }

  func setDirColor(str string) string {
  	return addColor(str, "fg-blue")
  }

  func setSelectedColor(str string) string {
  	colorBg := "bg-yellow"
  	if isDir(str) {
  		text := removeColor(str)
  		return addColor(text, "fg-blue,"+colorBg)
  	}
  	text := removeColor(str)
  	return addColor(text, colorBg)
  }

  func removeSelectedColor(str string) string {
  	if isDir(str) {
  		return setDirColor(removeColor(str))
  	}
  	return removeColor(str)
  }

  func getDirEntries(dirpath string) (entries []string) {
  	files, err := ioutil.ReadDir(dirpath)
  	if err != nil {
  		return
  	}

  	entries = append(entries, setDirColor(".."))
  	for _, file := range files {
  		if file.IsDir() {
  			entries = append(entries, setDirColor(file.Name()))
  		} else {
  			entries = append(entries, file.Name())
  		}
  	}
  	return
  }

  func main() {
  	err := termui.Init()
  	if err != nil {
  		panic(err)
  	}
  	defer termui.Close()

  	strs := getDirEntries(".")

  	ls := termui.NewList()
  	ls.Items = strs
  	ls.ItemFgColor = termui.ColorWhite
  	ls.BorderLabel = "Directory View"
  	ls.Height = 12
  	ls.Width = 25
  	ls.Y = 0

  	vs := ViewState{
  		ItemSelect: 0,
  	}
  	ls.Items[vs.ItemSelect] = setSelectedColor(ls.Items[vs.ItemSelect])

  	termui.Render(ls)
  	termui.Handle("/sys/kbd/q", func(termui.Event) {
  		termui.StopLoop()
  	})
  	termui.Handle("/sys/kbd/<up>", func(termui.Event) {
  		ls.Items[vs.ItemSelect] = removeSelectedColor(ls.Items[vs.ItemSelect])
  		vs.ItemSelect--
  		if vs.ItemSelect < 0 {
  			vs.ItemSelect = 0
  		}
  		ls.Items[vs.ItemSelect] = setSelectedColor(ls.Items[vs.ItemSelect])
  		termui.Render(ls)
  	})
  	termui.Handle("/sys/kbd/<down>", func(termui.Event) {
  		ls.Items[vs.ItemSelect] = removeSelectedColor(ls.Items[vs.ItemSelect])
  		vs.ItemSelect++
  		if vs.ItemSelect == len(ls.Items) {
  			vs.ItemSelect = len(ls.Items) - 1
  		}
  		ls.Items[vs.ItemSelect] = setSelectedColor(ls.Items[vs.ItemSelect])
  		termui.Render(ls)
  	})

  	termui.Loop()
  }


.. adsu:: 2

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10.1``

----

References:

.. [1] `GitHub - gizak/termui: Golang terminal dashboard <https://github.com/gizak/termui>`_

.. _ioutil.ReadDir: https://golang.org/pkg/io/ioutil/#ReadDir
.. _termui: https://github.com/gizak/termui
.. _list: https://github.com/gizak/termui/blob/master/_example/list.go
