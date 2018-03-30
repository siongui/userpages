Communication Between Chrome Extension and Local Application
############################################################

:date: 2018-03-30T23:09+08:00
:tags: JavaScript, Chrome Extension, Web application, Go, Golang, Go net/http,
       web
:category: Chrome Extension
:summary: Use local web server to achieve communication between Chrome extension
          and local program.
:og_image: http://smashinghub.com/wp-content/uploads/2011/09/google_chrome_extensions.jpg
:adsu: yes


The capability of Chrome extension running in browsers is restricted because of
security reasons. Sometimes we may want the extension to read the data we need
within browsers and pass the data to local applications for further processing.

A possible approach, with only one click from extension user (you), is that we
can create an clickable HTML link in the web page and the *href* property of the
link points to *localhost*, i.e., your local program running a web server. We
can embed the browser data, such as cookies or username, in the *href* property
of the link. When extension user (you) click the link, local application running
web servers receives HTTP request with browser data can perform further
processing.

For example, you can create an HTML link in your extension and append the link
to the web page as follows:

.. code-block:: javascript

  var link = document.createElement("a");
  link.setAttribute("href", "http://localhost:8999/download/username/profile_pic/");
  link.setAttribute("target", "_blank");
  document.querySelector("body").appendChild(link);

And in your local program (use Go as example), you can receive HTTP requests as
follows:

.. code-block:: go

  package main

  import (
  	"fmt"
  	"log"
  	"net/http"
  	"strings"
  )

  func handler(w http.ResponseWriter, r *http.Request) {
  	parts := strings.Split(r.URL.Path, "/")
  	if len(parts) != 5 {
  		fmt.Fprintf(w, "invalid request")
  		return
  	}
  	username := parts[2]
  	fmt.Fprintf(w, "username: %s\n", username)
  	action1 := parts[1]
  	action2 := parts[3]
  	fmt.Fprintf(w, "action: %s %s\n", action1, action2)

  	if action1 == "download" && action2 == "profile_pic" {
  		// download profile pic here
  	}
  }

  func main() {
  	http.HandleFunc("/", handler)
  	log.Fatal(http.ListenAndServe(":8999", nil))
  }

This approach requires only one click from extension user. If you only use it
for private, this is a possible and interesting approach to fit your needs.

----

.. adsu:: 2

References:

.. [1] | `Communication Between Chrome Extension and Local Program - Google search <https://www.google.com/search?q=Communication+Between+Chrome+Extension+and+Local+Program>`_
       | `Communication Between Chrome Extension and Local Program - DuckDuckGo search <https://duckduckgo.com/?q=Communication+Between+Chrome+Extension+and+Local+Program>`_
       | `Communication Between Chrome Extension and Local Program - Ecosia search <https://www.ecosia.org/search?q=Communication+Between+Chrome+Extension+and+Local+Program>`_
       | `Communication Between Chrome Extension and Local Program - Qwant search <https://www.qwant.com/?q=Communication+Between+Chrome+Extension+and+Local+Program>`_
       | `Communication Between Chrome Extension and Local Program - Bing search <https://www.bing.com/search?q=Communication+Between+Chrome+Extension+and+Local+Program>`_
       | `Communication Between Chrome Extension and Local Program - Yahoo search <https://search.yahoo.com/search?p=Communication+Between+Chrome+Extension+and+Local+Program>`_
       | `Communication Between Chrome Extension and Local Program - Baidu search <https://www.baidu.com/s?wd=Communication+Between+Chrome+Extension+and+Local+Program>`_
       | `Communication Between Chrome Extension and Local Program - Yandex search <https://www.yandex.com/search/?text=Communication+Between+Chrome+Extension+and+Local+Program>`_

.. _Chrome extension: https://developer.chrome.com/extensions/getstarted
