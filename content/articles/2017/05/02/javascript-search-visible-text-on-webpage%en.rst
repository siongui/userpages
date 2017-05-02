[JavaScript] Find All Visible Texts on Web Page
###############################################

:date: 2017-05-02T22:23+08:00
:tags: JavaScript, DOM, Traverse DOM Tree
:category: JavaScript
:summary: Traverse DOM tree to grab visible texts on webpage via JavaScript.
:adsu: yes

Show how to grab all visible texts on webpage via JavaScript.

This post is actually part of **[JavaScript] Traverse DOM Tree** [1]_. I extract
the code and make it an independent post for search engine friendliness.

The motivation to search for visible texts on the webpage is to find out all
visible texts on Simplified Chinese website and convert them to Traditional
Chinese. I write a Chrome extension in Go to do this job [2]_.

The following is the JavaScript code for finding out visible texts on web page:

.. code-block:: javascript

  "use strict";

  function isExcluded(elm) {
    if (elm.tagName == "STYLE") {
      return true;
    }
    if (elm.tagName == "SCRIPT") {
      return true;
    }
    if (elm.tagName == "NOSCRIPT") {
      return true;
    }
    if (elm.tagName == "IFRAME") {
      return true;
    }
    if (elm.tagName == "OBJECT") {
      return true;
    }
    return false
  }

  function traverse(elm) {
    if (elm.nodeType == Node.ELEMENT_NODE || elm.nodeType == Node.DOCUMENT_NODE) {

      // exclude elements with invisible text nodes
      if (isExcluded(elm)) {
        return
      }

      for (var i=0; i < elm.childNodes.length; i++) {
        // recursively call to traverse
        traverse(elm.childNodes[i]);
      }

    }

    if (elm.nodeType == Node.TEXT_NODE) {

      // exclude text node consisting of only spaces
      if (elm.nodeValue.trim() == "") {
        return
      }

      // elm.nodeValue here is visible text we need.
      console.log(elm.nodeValue);
    }
  }

  traverse(document);

`Traverse DOM tree`_ to find out all text nodes. Recursive call to *traverse*
function to visit all nodes in the DOM tree. In *traverse* function, we check
the node type of the element. If it is **ELEMENT_NODE** or **DOCUMENT_NODE**,
recursively call self with child nodes of current element as the argument. If it
is **TEXT_NODE**, then we are done.

To search for only text nodes with visible texts, we need some checks in the
code. First check the tag name of the element node, If the tag name is *STYLE*,
*SCRIPT*, *NOSCRIPT*, *IFRAME*, or *OBJECT*, rule out such elements. Also texts
consisting of only spaces are not needed either. We can check this by *trim()*
method.

.. adsu:: 2

Tested on: ``Chromium Version 58.0.3029.81 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``

----

References:

.. [1] `[JavaScript] Traverse DOM Tree <{filename}../01/javascript-traverse-dom-tree%en.rst>`_
.. [2] `[Golang/GopherJS] Chrome Extension for Chinese Conversion <{filename}../../04/30/go-gopherjs-chrome-extension-for-chinese-translation%en.rst>`_
.. [3] | `visible texts in web page - Google search <https://www.google.com/search?q=visible+texts+in+web+page>`_
       | `visible texts in web page - DuckDuckGo search <https://duckduckgo.com/?q=visible+texts+in+web+page>`_
       | `visible texts in web page - Ecosia search <https://www.ecosia.org/search?q=visible+texts+in+web+page>`_
       | `visible texts in web page - Qwant search <https://www.qwant.com/?q=visible+texts+in+web+page>`_
       | `visible texts in web page - Bing search <https://www.bing.com/search?q=visible+texts+in+web+page>`_
       | `visible texts in web page - Yahoo search <https://search.yahoo.com/search?p=visible+texts+in+web+page>`_
       | `visible texts in web page - Baidu search <https://www.baidu.com/s?wd=visible+texts+in+web+page>`_
       | `visible texts in web page - Yandex search <https://www.yandex.com/search/?text=visible+texts+in+web+page>`_

.. _Traverse DOM tree: https://www.google.com/search?q=Traverse+DOM+Tree
.. _JavaScript: https://www.google.com/search?q=JavaScript
