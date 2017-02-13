MathJax Tex/LaTex測試網頁
#########################

:date: 2016-04-11T21:27+08:00
:tags: Web開發
:category: Web開發
:summary: MathJax_ Tex_/LaTex_ 測試網頁
:og_image: https://upload.wikimedia.org/wikipedia/en/5/5f/MathJax.svg
:adsu: yes

.. raw:: html

  <script type="text/x-mathjax-config">
    MathJax.Hub.Config({
      tex2jax: {inlineMath: [["$","$"],["\\(","\\)"]]}
    });
  </script>
  <script type="text/javascript"
    src="//cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS_HTML-full">
  </script>

MathJax_ Tex_/LaTex_ 測試網頁(詳細請看 `Getting Started`_)：

.. adsu:: 2

.. raw:: html

  When $a \ne 0$, there are two solutions to \(ax^2 + bx + c = 0\) and they are
  $$x = {-b \pm \sqrt{b^2-4ac} \over 2a}.$$


.. _MathJax: https://www.mathjax.org/
.. _Tex: https://www.google.com/search?q=Tex
.. _LaTex: https://www.google.com/search?q=LaTex
.. _Getting Started: http://docs.mathjax.org/en/latest/start.html
