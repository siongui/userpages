如何透過GitHub在本站寫文章
##########################

:date: 2015-04-21 19:17
:modified: 2015-04-24T11:50+08:00
:tags: Web開發
:category: Web開發
:summary: 透過GitHub平台協同撰寫網站


背景
++++

先看Git背景：
`Git 10 周年访谈：Linus 讲述背后故事 <http://blog.jobbole.com/85772/>`_

後來 GitHub_ 推出 `GitHub Pages`_ ，結合近年興起的 `static site generator`_ ，
讓懂技術的部落客有非常powerful的工具來撰寫以前很難做到的部落格功能。

本站利用 `Pelican`_ 來製作部落格，有以下的優點：

  - 可用 reStructuredText_, Markdown_, 或 AsciiDoc_ 等等 `簡易標記語言`_ 撰寫

  - 因為用標記語言寫，可以自由搬移內容到各處/各服務。

  - 因為用標記語言寫，容易轉成其他格式，比方說HTML, PDF, 或是電子書等等。
    (可利用 Pandoc_ 之類的程式進行格式互轉)

  - 支援用LaTeX打數學式

  - 支援插入片斷的程式碼（包含程式碼上色）

  - 支援插入整個檔案的程式碼（包含程式碼上色）

  - 很容易就可以 HTML/JavaScript/CSS 程式demo

  - 容易在網頁嵌入影片以及聲音檔

  - 減少JavaScript，增進瀏覽網頁的速度

  - 可自由控制廣告位置，以及如何呈現，或是不放廣告

  - 可自由插入網頁評論系統： `Google+ comments`_, `Facebook comments`_,
    Disqus_, ...

  - 網站會根據螢幕大小做不同呈現(responsive web design)

  - 容易做SEO,以及可以指定Facebook分享時顯示的圖片及參數。

  - 支援多語系，同篇文章不同語言版本互相連結。

  - 幾乎可以這樣說：只有想不到，沒有做不到。

本網頁主要的repo是 userpages_, 含網站template以及網站內容都在此。然後透過
Pelican_ 編譯成整個網站，再上傳到 `siongui.github.io`_ 。

撰寫流程
++++++++

有兩種方式，一種是pull request方式，一種是collaborator方式。

(1) pull request方式
````````````````````

1. 首先註冊 GitHub_ 帳號

2. fork_ 本站 repo_, 利用 rst_ 格式寫文章，然後 `pull request`_ 提交。

(2) collaborator方式
````````````````````

1. 首先註冊 GitHub_ 帳號

2. 聯絡站主，站主將該帳號加為 userpages_ 的 collaborator，直接clone該repo，
   直接對 userpages_ 新增/修改/commit/push。

3. 此方式在commit前最好先git pull更新一下本機端clone的repo，若沒更新，
   若有兩人以上同時修改並push，則需要做一個merge的動作。

如何寫文章
++++++++++

透過範例學習最快，本站每篇文章都有一個 ``在GitHub上編輯`` 的連結，
可連回到原始檔案，然後在點選 ``raw`` 看原始 rst_ 格式。


文章的放置位置以及語言
``````````````````````

舉例來說，2015年4月21日的文章，就在
**content/articles/2015/04/21/**
目錄(若該目錄不存在則自己創建)下新增一個該文章的檔案，檔名取為

**SLUG%LANG.rst** ，其中 *SLUG* 是網址上的顯示名稱， *LANG* 是語言

舉例來說，若新增一個檔案名為 **content/articles/2015/04/21/post-on-this-site%zh.rst**
則該文章的網址是 **/zh/2015/04/21/post-on-this-site/** ，該文章的語言是傳統中文

若新增一個檔案名為 **content/articles/2015/04/21/post-on-this-site%en.rst**
則該文章的網址是 **/2015/04/21/post-on-this-site/** ，該文章的語言是英文


文章內容大致格式
````````````````

.. code-block:: txt

  如何透過GitHub在本站寫文章
  ##########################

  :date: 2015-04-21 19:17
  :modified: 2015-04-23 20:01
  :author: 某某某
  :tags: Web開發, CSS, HTML
  :category: Web開發
  :summary: 透過GitHub平台協同撰寫網站
  :og_image: http://....
  :license: CC|MIT|BSD|Apache2.0 ...

  (your main content here)

* **date** ：文章創建日期，可加可不加，不加的話則是用目錄裡的日期當成此文章日期

* **modified** ：文章修改日期

* **author** ：作者名

* **tags** ：可以有好幾個。

* **category** ：只可以有一個。

* **summary** ：文章摘要，此摘要即為 `HTML meta description`_ 以及 `og:description`_

* **og_image** ：是此文被分享或貼到Facebook或Google+之類的社交網站上時，顯示的圖片網址。
  (參考 [7]_)

* **license** ：此文章 and/or 程式碼的授權，可以是 CC_, MIT, BSD, Apache 2.0, ...

**tags**, **category**, **author**, **summary** 為建議必填，其他欄位可視情形填或不填。

rst_ 格式怎樣寫可參考 [1]_ ，至於用LaTeX寫數學，可看 [2]_ 。

如何撰寫文章的更多細節，請閱讀 `Pelican官方文件`_


注意事項
````````

- 每行建議不超過80個字母，一個中文算兩個字母。（非硬性規定）

- 若不確定LaTex顯示出來如何，可先在 [3]_ 輸入看結果。
  (參考 [4]_ 輸入數學符號)

- 可參考 [5]_ 來寫 rst_

- 可利用 `線上reStructuredText編輯器`_
  （可線上預覽，但因為CSS不同，預覽與實際網站呈現會有些差異）

- 亦可利用 `Sublime Text`_ + `OmniMarkupPreviewer`_ plugin
  來撰寫文章並預覽，但同樣因為CSS不同的關係，預覽與實際網站呈現會有些差異

- 標題下的 ``#`` 長度至少要比標題長度一樣長，或更長，例如以下是錯誤寫法：

  .. code-block:: rst

    [Math] The infamous Grasshopper problem
    ################################

  正確寫法是：

  .. code-block:: rst

    [Math] The infamous Grasshopper problem
    #######################################

- userpages_ 更改後，網站並不會更改，必須要 Pelican_ 把 userpages
  編譯成網站 HTML，然後更新對應的 `siongui.github.io`_ 內容。


SEO以及Facebook分享
```````````````````

網站SEO(意指容易在搜尋引擎被找到)有三個重點：

  - HTML title： 該網頁的title，必須配合搜尋關鍵字

  - URL：舉例來說，若網頁是談有關random number的文章，網址裡最好是：
    */2015/04/21/random-number/* ，將random number這兩個關鍵字包含在網址裡。
    若是用 */2015/04/21/blog-post_21.html* 之類的網址，將不利於SEO。

  - 日期：文章日期越新越好。

文章被分享或貼到Facebook或Google+之類的社交網站上時，
文章的顯示圖片網址是metadata裡的 *og_image*,
文章描述則是 *summary* 裡填寫的描述。
詳情請參考 [7]_ 或是 `Facebook官方指南`_ 。

預覽整個網站
++++++++++++

本站目前只能在 Ubuntu Linux 上將整個網站編譯出來並預覽，詳情請看：
`README <https://github.com/siongui/userpages/blob/master/README.rst>`_ 。
Windows平台理論上應該也可以將整個網站編譯出來並預覽，但從沒試過。

----

參考：

.. [1] `reStructuredText <http://docutils.sourceforge.net/rst.html>`_

.. [2] `reStructuredText Directives - math <http://docutils.sourceforge.net/docs/ref/rst/directives.html#math>`_

.. [3] `Online LaTeX Equation Editor - create, integrate and download <http://www.codecogs.com/latex/eqneditor.php>`_

.. [4] `LaTeX/Mathematics - Wikibooks, open books for an open world <http://en.wikibooks.org/wiki/LaTeX/Mathematics>`_

.. [5] `7. 附录：轻量级标记语言 — GotGitHub <http://www.worldhello.net/gotgithub/appendix/markups.html>`_

.. [6] `Online reStructuredText editor <http://rst.ninjs.org/>`_

.. [7] `Facebook Open Graph META Tags <http://davidwalsh.name/facebook-meta-tags>`_

----

附錄：

`travis ci + pelican <https://www.google.com/search?q=travis+ci+%2B+pelican>`_

.. _GitHub: https://github.com/
.. _fork: https://help.github.com/articles/fork-a-repo/
.. _repo: https://github.com/siongui/userpages
.. _rst: http://docutils.sourceforge.net/rst.html
.. _pull request: https://help.github.com/articles/using-pull-requests/
.. _GitHub Pages: https://pages.github.com/
.. _static site generator: https://www.google.com/search?q=static+site+generator
.. _Pelican: http://blog.getpelican.com/
.. _Google+ comments: http://browsingthenet.blogspot.com/2013/04/google-plus-comments-on-any-website.html
.. _Facebook comments: https://developers.facebook.com/docs/plugins/comments
.. _Disqus: https://disqus.com/
.. _reStructuredText: http://docutils.sourceforge.net/rst.html
.. _Markdown: http://daringfireball.net/projects/markdown/
.. _AsciiDoc: http://www.methods.co.nz/asciidoc/
.. _簡易標記語言: http://www.worldhello.net/gotgithub/appendix/markups.html
.. _Pandoc: http://pandoc.org/
.. _userpages: https://github.com/siongui/userpages
.. _siongui.github.io: https://github.com/siongui/siongui.github.io
.. _Facebook官方指南: https://developers.facebook.com/docs/sharing/best-practices
.. _CC: http://creativecommons.org.tw/
.. _Pelican官方文件: http://docs.getpelican.com/en/3.5.0/content.html
.. _線上reStructuredText編輯器: http://rst.ninjs.org/
.. _og\:description: http://davidwalsh.name/facebook-meta-tags
.. _HTML meta description: http://www.w3schools.com/tags/tag_meta.asp
.. _Sublime Text: http://www.sublimetext.com/
.. _OmniMarkupPreviewer: https://github.com/timonwong/OmniMarkupPreviewer
