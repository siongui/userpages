向 Linus Torvalds 學習讓編出的代碼具有 “good taste”
###################################################

:date: 2017-05-11T08:17+08:00
:tags: Linux, 轉錄, Linux 中國
:category: Linux
:author: Linux 中國翻譯組
:summary: 你的測試使用的條件語句越少，你的代碼就會有更好的 「taste」 。
:og_image: https://dn-linuxcn.qbox.me/data/attachment/album/201705/13/092618z2vyzoqzriqzavcf.jpg

.. image:: https://dn-linuxcn.qbox.me/data/attachment/album/201705/13/092618z2vyzoqzriqzavcf.jpg
   :alt: 向 Linus Torvalds 學習讓編出的代碼具有 “good taste”
   :align: center

在 `最近關於 Linus Torvalds 的一個採訪中`_ ，這位 Linux 的創始人，在採訪過程中大約 14:20 的時候，提及了關於代碼的 「good taste」。good taste？採訪者請他展示更多的細節，於是，Linus Torvalds 展示了一張提前準備好的插圖。

.. raw:: html

  <video width="640" height="360" controls="controls" poster="https://dn-linuxcn.qbox.me/data/attachment/album/201705/13/092618z2vyzoqzriqzavcf.jpg">
  <source src="https://dn-linuxcn.qbox.me/static/video/LinusTorvalds_2016-480p-zh-cn.mp4" type="video/mp4" />
  </video>

他展示的是一個代碼片段。但這段代碼並沒有 「good taste」。這是一個具有 「poor taste」 的代碼片段，把它作為例子，以提供一些初步的比較。

.. image:: https://cdn-images-1.medium.com/max/800/1*X2VgEA_IkLvsCS-X4iPY7g.png
   :alt: poor taste code
   :align: center

這是一個用 C 寫的函數，作用是刪除鏈表中的一個對象，它包含有 10 行代碼。

他把注意力集中在底部的 ``if`` 語句。正是這個 ``if`` 語句受到他的批判。

我暫停了這段視頻，開始研究幻燈片。我發現我最近有寫過和這很像的代碼。Linus 不就是在說我的代碼品味很差嗎？我放下自傲，繼續觀看視頻。

隨後， Linus 向觀眾解釋，正如我們所知道的，當從鏈表中刪除一個對象時，需要考慮兩種可能的情況。當所需刪除的對象位於鏈表的表頭時，刪除過程和位於鏈表中間的情況不同。這就是這個 ``if`` 語句具有 「poor taste」 的原因。

但既然他承認考慮這兩種不同的情況是必要的，那為什麼像上面那樣寫如此糟糕呢？

接下來，他又向觀眾展示了第二張幻燈片。這個幻燈片展示的是實現同樣功能的一個函數，但這段代碼具有 「goog taste」　。

.. image:: https://cdn-images-1.medium.com/max/800/1*GHFLYFB3vDQeakMyUGPglw.png
   :alt: good taste code
   :align: center

原先的 10 行代碼現在減少為 4 行。

但代碼的行數並不重要，關鍵是 ``if`` 語句，它不見了，因為不再需要了。代碼已經被重構，所以，不用管對象在列表中的位置，都可以運用同樣的操作把它刪除。

Linus 解釋了一下新的代碼，它消除了邊緣情況，就是這樣。然後採訪轉入了下一個話題。

我琢磨了一會這段代碼。 Linus 是對的，的確，第二個函數更好。如果這是一個確定代碼具有 「good taste」 還是 「bad taste」 的測試，那麼很遺憾，我失敗了。我從未想到過有可能能夠去除條件語句。我寫過不止一次這樣的 ``if`` 語句，因為我經常使用鏈表。

這個例子的意義，不僅僅是教給了我們一個從鏈表中刪除對象的更好方法，而是啟發了我們去思考自己寫的代碼。你通過程序實現的一個簡單算法，可能還有改進的空間，只是你從來沒有考慮過。

以這種方式，我回去審查最近正在做的項目的代碼。也許是一個巧合，剛好也是用 C 寫的。

我盡最大的能力去審查代碼，「good taste」 的一個基本要求是關於邊緣情況的消除方法，通常我們會使用條件語句來消除邊緣情況。
**你的測試使用的條件語句越少，你的代碼就會有更好的 「taste」 。**

下面，我將分享一個通過審查代碼進行了改進的一個特殊例子。

這是一個關於初始化網格邊緣的算法。

下面所寫的是一個用來初始化網格邊緣的算法，網格 grid 以一個二維數組表示：grid[行][列] 。

再次說明，這段代碼的目的只是用來初始化位於 grid 邊緣的點的值，所以，只需要給最上方一行、最下方一行、最左邊一列以及最右邊一列賦值即可。

為了完成這件事，我通過循環遍歷 grid 中的每一個點，然後使用條件語句來測試該點是否位於邊緣。代碼看起來就是下面這樣：

.. code-block:: c

  for (r = 0; r < GRID_SIZE; ++r) {
      for (c = 0; c < GRID_SIZE; ++c) {

          // Top Edge
          if (r == 0)
              grid[r][c] = 0;

          // Left Edge
          if (c == 0)
              grid[r][c] = 0;

          // Right Edge
          if (c == GRID_SIZE - 1)
              grid[r][c] = 0;

          // Bottom Edge
          if (r == GRID_SIZE - 1)
              grid[r][c] = 0;
      }
  }

雖然這樣做是對的，但回過頭來看，這個結構存在一些問題。

1. 複雜性 — 在雙層循環裡面使用 4 個條件語句似乎過於複雜。
2. 高效性 — 假設 ``GRID_SIZE`` 的值為 64，那麼這個循環需要執行 4096 次，但需要進行賦值的只有位於邊緣的 256 個點。

用 Linus 的眼光來看，將會認為這段代碼沒有 「good taste」 。

所以，我對上面的問題進行了一下思考。經過一番思考，我把複雜度減少為包含四個條件語句的單層 ``for`` 循環。雖然只是稍微改進了一下複雜性，但在性能上也有了極大的提高，因為它只是沿著邊緣的點進行了 256 次循環。

.. code-block:: c

  for (i = 0; i < GRID_SIZE * 4; ++i) {

      // Top Edge
      if (i < GRID_SIZE)
          grid[0][i] = 0;

      // Right Edge
      else if (i < GRID_SIZE * 2)
          grid[i - GRID_SIZE][GRID_SIZE - 1] = 0;

      // Left Edge
      else if (i < GRID_SIZE * 3)
          grid[i - (GRID_SIZE * 2)][0] = 0;

      // Bottom Edge
      else
          grid[GRID_SIZE - 1][i - (GRID_SIZE * 3)] = 0;
  }

的確是一個很大的提高。但是它看起來很醜，並不是易於閱讀理解的代碼。基於這一點，我並不滿意。

我繼續思考，是否可以進一步改進呢？事實上，答案是 YES！最後，我想出了一個非常簡單且優雅的算法，老實說，我不敢相信我會花了那麼長時間才發現這個算法。

下面是這段代碼的最後版本。它只有一層 ``for`` 循環並且沒有條件語句。另外。循環只執行了 64 次迭代，極大的改善了複雜性和高效性。

.. code-block:: c

  for (i = 0; i < GRID_SIZE; ++i) {

      // Top Edge
      grid[0][i] = 0;

      // Bottom Edge
      grid[GRID_SIZE - 1][i] = 0;

      // Left Edge
      grid[i][0] = 0;

      // Right Edge
      grid[i][GRID_SIZE - 1] = 0;
  }

這段代碼通過每次循環迭代來初始化四條邊緣上的點。它並不複雜，而且非常高效，易於閱讀。和原始的版本，甚至是第二個版本相比，都有天壤之別。

至此，我已經非常滿意了。

那麼，我是一個有 「good taste」 的開發者麼？

我覺得我是，但是這並不是因為我上面提供的這個例子，也不是因為我在這篇文章中沒有提到的其它代碼……而是因為具有 「good taste」 的編碼工作遠非一段代碼所能代表。Linus 自己也說他所提供的這段代碼不足以表達他的觀點。

我明白 Linus 的意思，也明白那些具有 「good taste」 的程序員雖各有不同，但是他們都是會將他們之前開發的代碼花費時間重構的人。他們明確界定了所開發的組件的邊界，以及是如何與其它組件之間的交互。他們試著確保每一樣工作都完美、優雅。

其結果就是類似於 Linus 的 「good taste」 的例子，或者像我的例子一樣，不過是千千萬萬個 「good taste」。

你會讓你的下個項目也具有這種 「good taste」 嗎？

----

| 編譯自：https://medium.com/@bartobri/applying-the-linus-tarvolds-good-taste-coding-requirement-99749f37684a
| 作者： Brian Barto
| 原創：LCTT_ https://linux.cn/article-8498-1.html
| 譯者： ucasFL_
| 本文地址：https://linux.cn/article-8498-1.html

- `Applying the Linus Torvalds “Good Taste” Coding Requirement <https://medium.com/@bartobri/applying-the-linus-tarvolds-good-taste-coding-requirement-99749f37684a>`_
- `向 Linus Torvalds 学习让编出的代码具有 “good taste”-观点 ◆ 热议|Linux.中国-开源社区 <https://linux.cn/article-8498-1.html>`_
- `向 Linus 学习，让代码具有 good taste - 文章 - 伯乐在线 <http://blog.jobbole.com/111159/>`_

.. _LCTT: https://linux.cn/lctt/
.. _ucasFL: https://linux.cn/lctt/ucasFL
.. _最近關於 Linus Torvalds 的一個採訪中: https://www.ted.com/talks/linus_torvalds_the_mind_behind_linux
