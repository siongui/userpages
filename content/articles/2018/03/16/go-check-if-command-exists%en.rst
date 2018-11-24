[Golang] Check If A Program (Command) Exists
############################################

:date: 2018-03-16T23:30+08:00
:tags: Go, Golang, Bash, Commandline
:category: Go
:summary: Check if a program (or command), such as *wget* or *ffmpeg*, is
          available in Go.
:og_image: https://i.stack.imgur.com/qN5s4.png
:adsu: yes


Check if a program (or command), such as *wget* or *ffmpeg*, is available in Go.

This looks like a easy task. I made some Google searches and found [2]_ you can
check if a command exists in shell as follows:

.. code-block:: bash

  $ command -v <the_command>

Now we use standard `os.exec`_ package to run the command for us:

.. code-block:: go

  import (
  	"os/exec"
  )

  func isCommandAvailable(name string) bool {
  	cmd := exec.Command("command", "-v", name)
  	if err := cmd.Run(); err != nil {
  		return false
  	}
  	return true
  }

Looks very easy. To be more confident, we write a test to see if it works:

.. code-block:: go

  import (
  	"testing"
  )

  func TestIsCommandAvailable(t *testing.T) {
  	if isCommandAvailable("ls") == false {
  		t.Error("ls command does not exist!")
  	}
  	if isCommandAvailable("ls111") == true {
  		t.Error("ls111 command should not exist!")
  	}
  }

And the test failed!!!

.. code-block:: bash

  === RUN   TestIsCommandAvailable
  --- FAIL: TestIsCommandAvailable (0.00s)
  	util_test.go:9: ls command does not exist!

I print the error message and got:

.. code-block:: bash

  exec: "command": executable file not found in $PATH

This is really strange, and I googled again and found [4]_ that you cannot run
built-in command directly via `exec.Command`_, so the following way of running
built-in command will not work:

.. code-block:: go

  	cmd := exec.Command("command", "-v", name)

Instead, you should run the built-in command as follows:

.. code-block:: go

  	cmd := exec.Command("/bin/sh", "-c", "command -v "+name)

And finally it works as expected and the test passed! The correct way to check
if a program exists in Go is:

.. code-block:: go

  import (
  	"os/exec"
  )

  func isCommandAvailable(name string) bool {
  	cmd := exec.Command("/bin/sh", "-c", "command -v "+name)
  	if err := cmd.Run(); err != nil {
  		return false
  	}
  	return true
  }


.. adsu:: 2

Tested on: `Ubuntu 17.10`, `Go 1.10`

----

References:

.. [1] | `bash check command exists - Google search <https://www.google.com/search?q=bash+check+command+exists>`_
       | `bash check command exists - DuckDuckGo search <https://duckduckgo.com/?q=bash+check+command+exists>`_
       | `bash check command exists - Ecosia search <https://www.ecosia.org/search?q=bash+check+command+exists>`_
       | `bash check command exists - Qwant search <https://www.qwant.com/?q=bash+check+command+exists>`_
       | `bash check command exists - Bing search <https://www.bing.com/search?q=bash+check+command+exists>`_
       | `bash check command exists - Yahoo search <https://search.yahoo.com/search?p=bash+check+command+exists>`_
       | `bash check command exists - Baidu search <https://www.baidu.com/s?wd=bash+check+command+exists>`_
       | `bash check command exists - Yandex search <https://www.yandex.com/search/?text=bash+check+command+exists>`_

.. [2] `shell - Check if a program exists from a Bash script - Stack Overflow <https://stackoverflow.com/a/677212>`_

.. [3] | `golang exec command path - Google search <https://www.google.com/search?q=golang+exec+command+path>`_
       | `golang exec command path - DuckDuckGo search <https://duckduckgo.com/?q=golang+exec+command+path>`_
       | `golang exec command path - Ecosia search <https://www.ecosia.org/search?q=golang+exec+command+path>`_
       | `golang exec command path - Qwant search <https://www.qwant.com/?q=golang+exec+command+path>`_
       | `golang exec command path - Bing search <https://www.bing.com/search?q=golang+exec+command+path>`_
       | `golang exec command path - Yahoo search <https://search.yahoo.com/search?p=golang+exec+command+path>`_
       | `golang exec command path - Baidu search <https://www.baidu.com/s?wd=golang+exec+command+path>`_
       | `golang exec command path - Yandex search <https://www.yandex.com/search/?text=golang+exec+command+path>`_

.. [4] `go - How to execute a linux built in command in golang - Stack Overflow <https://stackoverflow.com/a/34229542>`_

.. [5] | `golang exec builtin command - Google search <https://www.google.com/search?q=golang+exec+builtin+command>`_
       | `golang exec builtin command - DuckDuckGo search <https://duckduckgo.com/?q=golang+exec+builtin+command>`_
       | `golang exec builtin command - Ecosia search <https://www.ecosia.org/search?q=golang+exec+builtin+command>`_
       | `golang exec builtin command - Qwant search <https://www.qwant.com/?q=golang+exec+builtin+command>`_
       | `golang exec builtin command - Bing search <https://www.bing.com/search?q=golang+exec+builtin+command>`_
       | `golang exec builtin command - Yahoo search <https://search.yahoo.com/search?p=golang+exec+builtin+command>`_
       | `golang exec builtin command - Baidu search <https://www.baidu.com/s?wd=golang+exec+builtin+command>`_
       | `golang exec builtin command - Yandex search <https://www.yandex.com/search/?text=golang+exec+builtin+command>`_

.. [6] `Help! os/exec Output() on Non-English Windows cmd! : golang <https://old.reddit.com/r/golang/comments/9zsipj/help_osexec_output_on_nonenglish_windows_cmd/>`_

.. _os.exec: https://golang.org/pkg/os/exec/
.. _exec.Command: https://golang.org/pkg/os/exec/#Command
