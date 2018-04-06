[Golang] Get Last System Boot Time
##################################

:date: 2018-04-05T22:45+08:00
:tags: Go, Golang, Go time Package
:category: Go
:summary: Get last Ubuntu Linux system boot/restart time in Go.
:og_image: https://i.ytimg.com/vi/jTBtVrn-Cq8/hqdefault.jpg
:adsu: yes


This post shows you how to get last Ubuntu Linux system boot/restrart time in Go
via executing shell command.

To get last system boot time, we use *who -b* command [1]_ [2]_:

To get timezone of the system, we use *date +%Z* command [3]_ [4]_:

Then we combine the boot time and timezone string. Call `time.Parse`_ to parse
the string and return `time.Time`_ struct for further processing.

.. code-block:: go

  package main

  import (
  	"fmt"
  	"os/exec"
  	"strings"
  	"time"
  )

  func getLastBootTime() string {
  	out, err := exec.Command("who", "-b").Output()
  	if err != nil {
  		panic(err)
  	}
  	t := strings.TrimSpace(string(out))
  	t = strings.TrimPrefix(t, "system boot")
  	t = strings.TrimSpace(t)
  	return t
  }

  func getTimezone() string {
  	out, err := exec.Command("date", "+%Z").Output()
  	if err != nil {
  		panic(err)
  	}
  	return strings.TrimSpace(string(out))
  }

  func getLastSystemBootTime() (time.Time, error) {
  	return time.Parse(`2006-01-02 15:04MST`, getLastBootTime()+getTimezone())
  }

  func main() {
  	t, err := getLastSystemBootTime()
  	if err != nil {
  		panic(err)
  	}
  	fmt.Println(t.Format(time.RFC3339))
  }

Output from my Ubuntu Linux system:

.. code-block:: txt

  2018-04-05T17:56:00+08:00

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10.1``

**References**

.. [1] | `linux last startup time - Google search <https://www.google.com/search?q=linux+last+startup+time>`_
       | `linux last startup time - DuckDuckGo search <https://duckduckgo.com/?q=linux+last+startup+time>`_
       | `linux last startup time - Ecosia search <https://www.ecosia.org/search?q=linux+last+startup+time>`_
       | `linux last startup time - Qwant search <https://www.qwant.com/?q=linux+last+startup+time>`_
       | `linux last startup time - Bing search <https://www.bing.com/search?q=linux+last+startup+time>`_
       | `linux last startup time - Yahoo search <https://search.yahoo.com/search?p=linux+last+startup+time>`_
       | `linux last startup time - Baidu search <https://www.baidu.com/s?wd=linux+last+startup+time>`_
       | `linux last startup time - Yandex search <https://www.yandex.com/search/?text=linux+last+startup+time>`_
.. [2] `Linux Find Out Last System Reboot Time and Date Command - nixCraft <https://www.cyberciti.biz/tips/linux-last-reboot-time-and-date-find-out.html>`_
.. [3] | `linux date print timezone - Google search <https://www.google.com/search?q=linux+date+print+timezone>`_
       | `linux date print timezone - DuckDuckGo search <https://duckduckgo.com/?q=linux+date+print+timezone>`_
       | `linux date print timezone - Ecosia search <https://www.ecosia.org/search?q=linux+date+print+timezone>`_
       | `linux date print timezone - Qwant search <https://www.qwant.com/?q=linux+date+print+timezone>`_
       | `linux date print timezone - Bing search <https://www.bing.com/search?q=linux+date+print+timezone>`_
       | `linux date print timezone - Yahoo search <https://search.yahoo.com/search?p=linux+date+print+timezone>`_
       | `linux date print timezone - Baidu search <https://www.baidu.com/s?wd=linux+date+print+timezone>`_
       | `linux date print timezone - Yandex search <https://www.yandex.com/search/?text=linux+date+print+timezone>`_
.. [4] `How to Check Timezone in Linux <https://www.tecmint.com/check-linux-timezone/>`_

.. _time.Time: https://golang.org/pkg/time/#Time
.. _time.Parse: https://golang.org/pkg/time/#Parse
