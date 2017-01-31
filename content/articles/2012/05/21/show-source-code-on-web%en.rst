Show Source Code on Web
#######################

:date: 2012-05-21 01:07
:modified: 2015-04-02 21:11
:tags: web, html
:category: Web
:summary: Show source code on your website
:adsu: yes


Sometime it is helpful to show source code on the website or blog. One way to do
it is to write your code in rst_ or markdown_ format, and then convert the rst_
or markdown_ document to html. Another way is to use online source code to html
coverter. Google search "`source code to html`_". I found [1]_ is the best
online converter:

Demo (with line numbers):

.. raw:: html

  <!-- HTML generated using hilite.me --><div style="background: #ffffff; overflow:auto;width:auto;border:solid gray;border-width:.1em .1em .1em .8em;padding:.2em .6em;"><table><tr><td><pre style="margin: 0; line-height: 125%">1
  2
  3
  4
  5
  6
  7</pre></td><td><pre style="margin: 0; line-height: 125%"><span style="color: #557799">#include &lt;stdio.h&gt;</span>

  <span style="color: #333399; font-weight: bold">int</span> <span style="color: #0066BB; font-weight: bold">main</span>() {
    printf(<span style="background-color: #fff0f0">&quot;hello world</span><span style="color: #666666; font-weight: bold; background-color: #fff0f0">\n</span><span style="background-color: #fff0f0">&quot;</span>);

    <span style="color: #008800; font-weight: bold">return</span> <span style="color: #0000DD; font-weight: bold">0</span>;
  }
  </pre></td></tr></table></div>

.. adsu:: 2

Demo (without line numbers):

.. raw:: html

  <!-- HTML generated using hilite.me --><div style="background: #ffffff; overflow:auto;width:auto;border:solid gray;border-width:.1em .1em .1em .8em;padding:.2em .6em;"><pre style="margin: 0; line-height: 125%"><span style="color: #557799">#include &lt;stdio.h&gt;</span>

  <span style="color: #333399; font-weight: bold">int</span> <span style="color: #0066BB; font-weight: bold">main</span>() {
    printf(<span style="background-color: #fff0f0">&quot;hello world</span><span style="color: #666666; font-weight: bold; background-color: #fff0f0">\n</span><span style="background-color: #fff0f0">&quot;</span>);

    <span style="color: #008800; font-weight: bold">return</span> <span style="color: #0000DD; font-weight: bold">0</span>;
  }
  </pre></div>

----

.. adsu:: 3

References:

.. [1] `Source code beautifier / syntax highlighter – convert code snippets to HTML « hilite.me <http://hilite.me/>`_

More silimar online converters:

.. [2] `Online syntax highlighting <http://tohtml.com/>`_

.. [3] `Code2HTML <https://www.palfrader.org/code/code2html/cgi/>`_

.. [4] `Convert Code Online <http://puzzleware.net/CodeHtmler/default.aspx>`_

.. [5] `Online syntax highlighter like TextMate <http://markup.su/highlighter/>`_

.. [6] `Quick Highlighter: A simple to use code syntax highlighter <http://quickhighlighter.com/>`_


.. _rst: http://docutils.sourceforge.net/rst.html
.. _markdown: http://daringfireball.net/projects/markdown/
.. _source code to html: https://www.google.com/search?q=source+code+to+html
