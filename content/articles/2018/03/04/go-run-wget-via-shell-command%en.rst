[Golang] Run wget via Shell Command
###################################

:date: 2018-03-04T08:50+08:00
:tags: Go, Golang, Commandline
:category: Go
:summary: Use Go standard *os/exec* package to run wget via shell command.
:og_image: https://upload.wikimedia.org/wikipedia/commons/6/60/Wget_1.13.4.png
:adsu: yes

In your Go program, you can download things via hget_ or pget_, which is pure Go
implementation and portable in all platforms. But if you do not care about
portability and just want to use wget_ on your Linux or Unix-like platform, we
can use Go standard `os/exec`_ package to run wget_ via shell command. The
following is howto:

.. code-block:: go

  import (
  	"os"
  	"os/exec"
  )

  // Call shell command wget to download. The reason to use wget is that wget
  // supports automatically resume download. So this package only runs on Linux
  // systems.
  func wget(url, filepath string) error {
  	// run shell `wget URL -O filepath`
  	cmd := exec.Command("wget", url, "-O", filepath)
  	cmd.Stdout = os.Stdout
  	cmd.Stderr = os.Stderr
  	return cmd.Run()
  }

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10``.

----

References:

.. [1] | `golang wget - Google search <https://www.google.com/search?q=golang+wget>`_
       | `golang wget - DuckDuckGo search <https://duckduckgo.com/?q=golang+wget>`_
       | `golang wget - Ecosia search <https://www.ecosia.org/search?q=golang+wget>`_
       | `golang wget - Qwant search <https://www.qwant.com/?q=golang+wget>`_
       | `golang wget - Bing search <https://www.bing.com/search?q=golang+wget>`_
       | `golang wget - Yahoo search <https://search.yahoo.com/search?p=golang+wget>`_
       | `golang wget - Baidu search <https://www.baidu.com/s?wd=golang+wget>`_
       | `golang wget - Yandex search <https://www.yandex.com/search/?text=golang+wget>`_
.. [2] | `GitHub - huydx/hget: interruptable, resumable download accelerator written in golang <https://github.com/huydx/hget>`_
       | `GitHub - Code-Hex/pget: The fastest file download client <https://github.com/Code-Hex/pget>`_
       | `GitHub - laher/wget-go: wget, partially re-implemented in go <https://github.com/laher/wget-go>`_
       | `wget - Go libraries and apps <https://golanglibs.com/top?q=wget>`_

.. _os/exec: https://golang.org/pkg/os/exec/
.. _hget: https://github.com/huydx/hget
.. _pget: https://github.com/Code-Hex/pget
.. _wget: https://www.gnu.org/software/wget/
