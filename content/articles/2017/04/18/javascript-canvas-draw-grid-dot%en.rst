[JavaScript] Draw Grid of Dots using Canvas
###########################################

:date: 2017-04-18T22:53+08:00
:tags: html, JavaScript, DOM, Canvas
:category: JavaScript
:summary: Exercise to draw 3x3 grid of dots via *Canvas*.
:og_image: https://i.stack.imgur.com/H1Qsz.png
:adsu: yes

Exercise to draw 3x3 grid of dots via HTML Canvas_.

.. raw:: html

  <div style="padding: 2em; background-color: Azure; padding: 5px;">

  <canvas id="gridDots" width="300" height="300"></canvas>

  </div>
  <script>

  var c = document.getElementById("gridDots");
  var ctx = c.getContext("2d");

  var n = 3; // 3x3 grid of dots

  for (var i=0; i<n; i++) {
    for (var j=0; j<n; j++) {
      var gridXSideLength = c.width/n;
      var gridYSideLength = c.height/n;
      var x = (gridXSideLength/2) * (i+1);
      var y = (gridYSideLength/2) * (j+1);
      var r = Math.min(gridXSideLength, gridYSideLength) / (n*2);
      var sAngle = 0;
      var eAngle = 2*Math.PI;

      ctx.beginPath();
      ctx.arc(x, y, r, sAngle, eAngle);
      ctx.closePath();
      ctx.stroke();
    }
  }

  </script>


**HTML**:

.. code-block:: html

  <canvas id="gridDots" width="300" height="300"></canvas>


**JavaScript**:

.. code-block:: javascript

  var c = document.getElementById("gridDots");
  var ctx = c.getContext("2d");

  var n = 3; // 3x3 grid of dots

  for (var i=0; i<n; i++) {
    for (var j=0; j<n; j++) {
      var gridXSideLength = c.width/n;
      var gridYSideLength = c.height/n;
      var x = (gridXSideLength/2) * (i+1);
      var y = (gridYSideLength/2) * (j+1);
      var r = Math.min(gridXSideLength, gridYSideLength) / (n*2);
      var sAngle = 0;
      var eAngle = 2*Math.PI;

      ctx.beginPath();
      ctx.arc(x, y, r, sAngle, eAngle);
      ctx.closePath();
      ctx.stroke();
    }
  }


.. adsu:: 2

----

References:

.. [1] | `用 canvas 实现 Web 手势解锁 - 渔人 - SegmentFault <https://segmentfault.com/a/1190000008923731>`_
       | `用 canvas 实现 Web 手势解锁 - WEB前端 - 伯乐在线 <http://web.jobbole.com/90970/>`_
       | `GitHub - songjinzhong/H5HandLock: H5 手势密码，炫酷！！https://songjinzhong.github.io/H5HandLock/ <https://github.com/songjinzhong/H5HandLock>`_
       | `GitHub - lvming6816077/H5lock: H5手势解锁 <https://github.com/lvming6816077/H5lock>`_

.. [2] | `gesture unlock screen grid dots - Google search <https://www.google.com/search?q=gesture+unlock+screen+grid+dots>`_
       | `gesture unlock screen grid dots - DuckDuckGo search <https://duckduckgo.com/?q=gesture+unlock+screen+grid+dots>`_
       | `gesture unlock screen grid dots - Ecosia search <https://www.ecosia.org/search?q=gesture+unlock+screen+grid+dots>`_
       | `gesture unlock screen grid dots - Qwant search <https://www.qwant.com/?q=gesture+unlock+screen+grid+dots>`_
       | `gesture unlock screen grid dots - Bing search <https://www.bing.com/search?q=gesture+unlock+screen+grid+dots>`_
       | `gesture unlock screen grid dots - Yahoo search <https://search.yahoo.com/search?p=gesture+unlock+screen+grid+dots>`_
       | `gesture unlock screen grid dots - Baidu search <https://www.baidu.com/s?wd=gesture+unlock+screen+grid+dots>`_
       | `gesture unlock screen grid dots - Yandex search <https://www.yandex.com/search/?text=gesture+unlock+screen+grid+dots>`_

.. [3] | `HTML5 Canvas <https://www.w3schools.com/html/html5_canvas.asp>`_
       | `HTML canvas arc() Method <https://www.w3schools.com/tags/canvas_arc.asp>`_
       | `Math.min() - JavaScript | MDN <https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/min>`_

.. _Canvas: https://www.google.com/search?q=html+canvas
