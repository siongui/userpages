JavaScript Back to Top Link
###########################

:date: 2012-06-27 07:30
:modified: 2015-04-06 22:30
:tags: JavaScript, html
:category: JavaScript
:summary: JavaScript Back to Top Link
:adsu: yes


To add a "back to top" link is very easy, just add one line of code in the HTML
document (from [1]_):

.. code-block:: javascript

  <a href="#" onclick="window.scrollTo(0,0); return false">Back to Top</a>

The above code is "staic", what if I want to dynamically create a "back to top"
link?

It seems pretty easy, so I wrote the following code:

.. code-block:: javascript

  var backToTop = document.createElement("a");
  backToTop.href = "#";
  backToTop.onclick = "window.scrollTo(0,0); return false;";
  backToTop.innerHTML = "Back to Top";

  // append backToTop to some element...

This looks exactly the same as above, but actually not. The correct way to write
"dynamic" version of "back to top" link should be:

.. code-block:: javascript

  var backToTop = document.createElement("a");
  backToTop.href = "#";
  backToTop.onclick = function(e){window.scrollTo(0,0);return false;};
  backToTop.innerHTML = "Back to Top";

  // append backToTop to some element...

Look at the difference closely again:

**Wrong**:

.. code-block:: javascript

  backToTop.onclick = "window.scrollTo(0,0); return false;";

**Right**:

.. code-block:: javascript

  backToTop.onclick = function(e){window.scrollTo(0,0);return false;};

If you do it the wrong way as I originally did, once the link is clicked, users
will be re-directed to "#" before the window is scrolled to top! That means
"*return default;*" does not prevent the default action of the *onclick* event.
This conclusion comes from my trial and error, and I don't know the syntax
differences of the two versions.

There is another way to create "back to top" link dynamically:

.. code-block:: javascript

  var backToTop = document.createElement("a");
  backToTop.href = "javascript:window.scrollTo(0,0);";
  backToTop.innerHTML = "Back to Top";

  // append backToTop to some element...

The above code assigns JavaScript command directly to the *href* property.
However, I don't know whether this is good practice for coding or not. I still
put the code here for references.

----

Reference:

.. [1] `CodeSnippets: Back to Top [javascript] <http://codesnippets.joyent.com/posts/show/214>`_
