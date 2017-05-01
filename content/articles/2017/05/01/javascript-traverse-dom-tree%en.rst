[JavaScript] Traverse DOM Tree
##############################

:date: 2017-05-01T14:49+08:00
:tags: JavaScript, DOM, Traverse DOM Tree
:category: JavaScript
:summary: Traverse DOM tree to find out all text nodes via JavaScript.
:adsu: yes

`Traverse DOM tree`_ to find out all text nodes via JavaScript_.

.. code-block:: javascript

  "use strict";

  function traverse(elm) {
    if (elm.nodeType == Node.ELEMENT_NODE || elm.nodeType == Node.DOCUMENT_NODE) {
      for (var i=0; i < elm.childNodes.length; i++) {
        // recursively call to traverse
        traverse(elm.childNodes[i]);
      }
    }

    if (elm.nodeType == Node.TEXT_NODE) {
      console.log(elm.nodeValue);
    }
  }

  traverse(document);

Recursive call to *traverse* function to visit all nodes in the DOM tree. In
*traverse* function, we check the node type of the element. If it is
**ELEMENT_NODE** or **DOCUMENT_NODE**, recursively call self with child nodes of
current element as the argument. If it is **TEXT_NODE**, then we are done.

.. adsu:: 2

Sometimes we need only text nodes with visible texts. The invisible texts in
elements, such as *style*, or texts consisting of only spaces are not needed.
The following code can help traverse the DOM tree to find out only visible text
nodes.

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

      console.log(elm.nodeValue);
    }
  }

  traverse(document);

.. adsu:: 3

Tested on: ``Chromium Version 58.0.3029.81 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``

----

References:

.. [1] | `Traverse DOM Tree - Google search <https://www.google.com/search?q=Traverse+DOM+Tree>`_
       | `Traverse DOM Tree - DuckDuckGo search <https://duckduckgo.com/?q=Traverse+DOM+Tree>`_
       | `Traverse DOM Tree - Ecosia search <https://www.ecosia.org/search?q=Traverse+DOM+Tree>`_
       | `Traverse DOM Tree - Qwant search <https://www.qwant.com/?q=Traverse+DOM+Tree>`_
       | `Traverse DOM Tree - Bing search <https://www.bing.com/search?q=Traverse+DOM+Tree>`_
       | `Traverse DOM Tree - Yahoo search <https://search.yahoo.com/search?p=Traverse+DOM+Tree>`_
       | `Traverse DOM Tree - Baidu search <https://www.baidu.com/s?wd=Traverse+DOM+Tree>`_
       | `Traverse DOM Tree - Yandex search <https://www.yandex.com/search/?text=Traverse+DOM+Tree>`_

.. [2] | `javascript traverse childnodes - Google search <https://www.google.com/search?q=javascript+traverse+childnodes>`_
       | `javascript traverse childnodes - DuckDuckGo search <https://duckduckgo.com/?q=javascript+traverse+childnodes>`_
       | `javascript traverse childnodes - Ecosia search <https://www.ecosia.org/search?q=javascript+traverse+childnodes>`_
       | `javascript traverse childnodes - Qwant search <https://www.qwant.com/?q=javascript+traverse+childnodes>`_
       | `javascript traverse childnodes - Bing search <https://www.bing.com/search?q=javascript+traverse+childnodes>`_
       | `javascript traverse childnodes - Yahoo search <https://search.yahoo.com/search?p=javascript+traverse+childnodes>`_
       | `javascript traverse childnodes - Baidu search <https://www.baidu.com/s?wd=javascript+traverse+childnodes>`_
       | `javascript traverse childnodes - Yandex search <https://www.yandex.com/search/?text=javascript+traverse+childnodes>`_

.. [3] `javascript - TEXT_NODE: returns ONLY text? - Stack Overflow <http://stackoverflow.com/questions/6087399/text-node-returns-only-text>`_
.. [4] `HTML DOM tagName Property <https://www.w3schools.com/jsref/prop_element_tagname.asp>`_

.. _Traverse DOM tree: https://www.google.com/search?q=Traverse+DOM+Tree
.. _JavaScript: https://www.google.com/search?q=JavaScript
