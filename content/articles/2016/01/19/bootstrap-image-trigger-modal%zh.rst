Bootstrap利用圖片來觸發modal視窗
################################

:date: 2016-01-19T09:13+08:00
:tags: 響應式網頁設計(Responsive Web Design), Web開發, CSS
:category: 響應式網頁設計(Responsive Web Design)
:summary: Bootstrap_ 利用圖片(image_)來觸發 modal_ 視窗

在 Bootstrap_ `官網裡的範例`_ 是用按鈕(Button_)來觸發 modal_ 視窗

.. code-block:: html

  <!-- Button trigger modal -->
  <button type="button" class="btn btn-primary btn-lg" data-toggle="modal" data-target="#myModal">
    Launch demo modal
  </button>

若不想用按鈕(Button_)，想用 image_ 來觸發，該怎樣做？
經過Google [3]_ 後，發現 [1]_ 有解答：

.. code-block:: html

  <a href="#myModal" data-toggle="modal">
    <img src="/img/myImage.png">
  </a>


----

參考：

.. [1] `Can I use an image to trigger a modal window in Bootstrap? <http://stackoverflow.com/questions/15423532/can-i-use-an-image-to-trigger-a-modal-window-in-bootstrap>`_

.. [2] `Bootstrap Image trigger modal example code <http://www.bootply.com/7wOLkC9AVX>`_

.. [3] Google搜尋 `Bootstrap image trigger modal <https://www.google.com/search?q=Bootstrap+image+trigger+modal>`__

.. _Bootstrap: http://getbootstrap.com/
.. _Button: http://www.w3schools.com/tags/tag_button.asp
.. _image: http://www.w3schools.com/tags/tag_img.asp
.. _modal: http://www.w3schools.com/bootstrap/bootstrap_modal.asp
.. _官網裡的範例: http://getbootstrap.com/javascript/#modals
