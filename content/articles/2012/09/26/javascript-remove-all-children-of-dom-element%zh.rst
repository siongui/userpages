JavaScript刪除DOM元素的所有子節點
#################################

:date: 2018-09-29T01:36+08:00
:tags: JavaScript, Web開發
:category: JavaScript
:summary: JavaScript無臭蟲地移除一個DOM元素(element)的所有子節點(node)。
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


這裡有一個有趣的問題：JavaScript如何移除一個DOM元素(element)的所有子節點(node)？

假設我們有一個DOM元素（例如， *div* 或 *span* ），元素名是 *elm* ，
移除所有子節點，最簡單而且直覺的方式如下：

.. code-block:: javascript

  elm.innerHTML = '';

這看起來很完美，而且應該不會有任何問題。事實上，
我一開始使用這個方法來刪除所有子節點，而且也運作得沒問題。但有一次，
我又在程式碼裡使用這個方法，在最新版的Chrome，Firefix和Opera沒問題，
但在IE8卻沒辦法運作。IE8在下面這個程式碼片段出了出來：

.. code-block:: javascript

  elm.innerHTML = '';
  elm.appendChild(childs);

我不知道為何上面的代碼在IE8不能運作，
所以我搜尋了一下並找到了更好的方式來移除所有子節點 [1]_ [2]_。
正統而且無臭蟲的方式如下：

.. code-block:: javascript

  while (elm.hasChildNodes()) {
    elm.removeChild(elm.lastChild);
  }

也許對某些人而言正統的方式是顯而易見的，
但這花了我一些時間去找臭蟲並領悟到有時簡單的方法不一定能用。
所以我寫了這篇博文來幫那些跟我有同樣問題的人。

.. adsu:: 2

----

**附錄**

JavaScript把陣列(array)清空的其中一個方法是： [5]_

.. code-block:: javascript

  while(A.length > 0) {
    A.pop();
  }

就如所見，這裡 *清空陣列* 的作法跟 *移除所有子節點* 是一樣的。
這蠻有趣的所以我把此放在附錄。

.. adsu:: 3

----

References:

.. [1] `javascript - Remove all the children DOM elements in div - Stack Overflow <http://stackoverflow.com/questions/683366/remove-all-the-children-dom-elements-in-div>`_
.. [2] `Remove all child elements of a DOM node in JavaScript - Stack Overflow <http://stackoverflow.com/questions/3955229/remove-all-child-elements-of-a-dom-node-in-javascript>`_
.. [3] `[Dart] DOM Element Remove All Children <{filename}../../../2014/01/31/dart-element-remove-all-children%en.rst>`_
.. [4] `[Golang] Remove All Child Nodes of a DOM Element by GopherJS <{filename}../../../2016/01/31/go-remove-all-children-of-dom-element-by-gopherjs%en.rst>`_
.. [5] `How do I empty an array in JavaScript? - Stack Overflow <http://stackoverflow.com/questions/1232040/how-do-i-empty-an-array-in-javascript>`_
