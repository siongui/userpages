JavaScript Wait Page to be Loaded
#################################

:date: 2018-02-23T10:34+08:00
:tags: JavaScript, DOM
:category: JavaScript
:summary: When web pages are rendered with JavaScript, how to wait until the DOM
          nodes of original contents are available for appending our nodes?
:og_image: https://i.ytimg.com/vi/EA27xM71m0g/maxresdefault.jpg
:adsu: yes


Recently I am writing Chrome extension. I need to create DOM nodes and append
the nodes to the original contents of the wepage. But when I tried to append my
created DOM nodes to the nodes in the original content, nothing happened! I
cannot see any of my nodes appended.

I investigate the problem and found that the original page use a lot of
JavaScript to render contents of the page, and the DOM nodes of original
contents are not ready for appending my created nodes. I tried to do some Google
searches [1]_, and some answers said we can listen to *onload* event of
*DOMContentLoaded* event. But the answer does not work in my case.

I wrote post of DOM ready [2]_ long time ago, and the trick in the post is
periodically check if the node to which my nodes are appended is available. If
the node is not ready, sleep for a while and check again until the node is
available. I applied this trick to my case, and it worked!

The following is key point of the trick:

.. code-block:: javascript

  // wait the node to be ready
  // check interval: 500ms
  var timerId = setInterval(function() {
    var n = document.querySelector("css-selector-of-node");
    if (n != null) {
      // now the node is available, do something here.
      // and also clear the timer.
      clearInterval(timerId);
    }
  }, 500);

Hope this would be helpful for those who have the same problem as me.

.. adsu:: 2

----

References:

.. [1] | `javascript wait page load - Google search <https://www.google.com/search?q=javascript+wait+page+load>`_
       | `javascript wait page load - DuckDuckGo search <https://duckduckgo.com/?q=javascript+wait+page+load>`_
       | `javascript wait page load - Ecosia search <https://www.ecosia.org/search?q=javascript+wait+page+load>`_
       | `javascript wait page load - Qwant search <https://www.qwant.com/?q=javascript+wait+page+load>`_
       | `javascript wait page load - Bing search <https://www.bing.com/search?q=javascript+wait+page+load>`_
       | `javascript wait page load - Yahoo search <https://search.yahoo.com/search?p=javascript+wait+page+load>`_
       | `javascript wait page load - Baidu search <https://www.baidu.com/s?wd=javascript+wait+page+load>`_
       | `javascript wait page load - Yandex search <https://www.yandex.com/search/?text=javascript+wait+page+load>`_

.. [2] `DOM Ready without JavaScript Frameworks <{filename}/articles/2012/10/04/dom-ready-without-javascript-frameworks%en.rst>`_
