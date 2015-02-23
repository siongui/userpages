Load External JavaScript or CSS file Dynamically
################################################

:date: 2012-06-18 20:29
:modified: 2015-02-23 10:27
:tags: JavaScript, CORS, CSS
:category: JavaScript
:summary: Load JavaScript or CSS on demand.


Sometimes it is useful to load JavaScript or CSS file dynamically. In [1]_ and
[2]_, there are several different method to do this. Here I do not do anything
original. I just summarize the posts and give a single function to load external
JavaScript or CSS files on demand in two different ways. The following is my
summarized function to load JavaScript or CSS:

Source Code
+++++++++++

.. code-block:: javascript

  function AJAXLoad(type, loadByXMLHttp, url) {
    var ext;
    if (type == "js") {
      ext = document.createElement('script');
      ext.setAttribute("type","text/javascript");
    }
    if (type == "css") {
      ext = document.createElement('link');
      ext.setAttribute("type", "text/css");
      ext.setAttribute("rel", "stylesheet");
    }
    if (!loadByXMLHttp) {
      if (type == "js") {
        ext.setAttribute("src", url);
        if (typeof ext != "undefined") {document.getElementsByTagName("head")[0].appendChild(ext);}
      }
      if (type == "css") {
        ext.setAttribute("href", url);
        if (typeof ext != "undefined") {document.getElementsByTagName("head")[0].appendChild(ext);}
      }
      return;
    }
    var xmlhttp;
    if (window.XMLHttpRequest) {xmlhttp=new XMLHttpRequest();}
    else {xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");}
    xmlhttp.onreadystatechange=function(){
      if (xmlhttp.readyState==4) {
        if (xmlhttp.status == 200) {
          ext.innerHTML=xmlhttp.responseText;
          document.getElementsByTagName("head")[0].appendChild(ext);
        } else {
          console.log('cannot load external file :'+url);
        }
      }
    }
    xmlhttp.open("GET",url,true);
    xmlhttp.send();
  }

Usage
+++++

.. code-block:: javascript

  // Usage 1: Load CSS in element tag
  AJAXLoad("css", false, "http://fonts.googleapis.com/css?family=Gentium+Basic|Special+Elite&subset=latin,latin-ext");

  // Usage 2: Load JavaScript in element tag
  AJAXLoad("js", false, "http://www.example.com/static/index.js");

  // Usage 3: Load CSS by XMLHttp call (AJAX way)
  AJAXLoad("css", true, "http://fonts.googleapis.com/css?family=Gentium+Basic|Special+Elite&subset=latin,latin-ext");

  // Usage 4: Load JavaScript in by XMLHttp call (AJAX way)
  AJAXLoad("js", true, "http://www.example.com/static/index.js");

For most cases, usage 1 and 2 will work as expected. But usage 3 or 4 will not
work because of `cross domain scripting`_. For security reasons, XMLHttp call to
servers on other domains is not allowed. There are several workaround to achieve
`cross domain scripting`_, but this is beyond the topic of this post.

If you load JavaScript after the page is ready, there are chances that you code
will not run automatically. Please refer to [3]_, [4]_, and [5]_ for more
details. And to know how XMLHttp works, please refer to [6]_.

----

References:

.. [1] `4 ways to dynamically load external JavaScript(with source) <http://ntt.cc/2008/02/10/4-ways-to-dynamically-load-external-javascriptwith-source.html>`_

.. [2] `Dynamically loading an external JavaScript or CSS file <http://www.javascriptkit.com/javatutors/loadjavascriptcss.shtml>`_

.. [3] `Are dynamically inserted <script> tags meant to work? <http://stackoverflow.com/questions/1891947/are-dynamically-inserted-script-tags-meant-to-work>`_

.. [4] `Can scripts be inserted with innerHTML? <http://stackoverflow.com/questions/1197575/can-scripts-be-inserted-with-innerhtml>`_

.. [5] `Executing <script> inside <div> retrieved by AJAX <http://stackoverflow.com/questions/4619668/executing-script-inside-div-retrieved-by-ajax>`_

.. [6] `AJAX Tutorial <http://www.w3schools.com/ajax/default.asp>`_


.. _cross domain scripting: http://en.wikipedia.org/wiki/Cross-site_scripting
