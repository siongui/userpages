擁抱開源，這個城市的法典都通過 GitHub 發佈
##########################################

:date: 2018-11-27T08:23+08:00
:tags: 轉錄, 開源中國, git
:category: Web開發
:author: h4cd(開源中國)
:summary: 如果你訪問下邊這個 GitHub repo，會發現它記錄的是一部法典：
:og_image: https://oscimg.oschina.net/oscnet/a307d5bc3092b6a7cf30a28191d50184cb9.jpg


如果你訪問下邊這個 GitHub repo，會發現它記錄的是一部法典：

https://github.com/DCCouncil/dc-law-xml

.. image:: https://oscimg.oschina.net/oscnet/a307d5bc3092b6a7cf30a28191d50184cb9.jpg
   :alt: 擁抱開源，這個城市的法典都通過 GitHub 發佈
   :align: center

這是美國華盛頓哥倫比亞特區（Washington District of Columbia，WDC/DC）的法律條款，它不是法律的副本，而是一個權威的 DC 委員會存儲已頒布法律的數字版本的地方，內容直接提供給理事會的 DC 條款網站： https://code.dccouncil.us/dc/council/code 。

幾年前，DC 律師與軟件工程師發現，在跟蹤代碼變更時，立法和軟件工程有很多共同之處，於是他們決定將法律條款託管到 GitHub。他們成功實施了該計畫，於是有了上邊提到的這個 GitHub repo。

是的，這是一個真實而有意義 repo，它是可以通過 PR 進行更新的，最近黑客 `Joshua Tauberer`_ 還發文介紹了他為該法典提交了一個錯別字 PR 並被接受的案例。Joshua 發現條款中這句描述出現錯誤：

.. highlights::

  (d) The Office of Open Government may issue advisory opinions on the implementation of subchapter I of Chapter 5 of Title 2.

其中，“subchapter I” 應為 “subchapter II”，於是 Joshua 提了一個 PR 更正該錯誤。幾天後，委員會的編纂律師將其合併：

.. image:: https://oscimg.oschina.net/oscnet/26b1aa54d2a7be9d68f990e070570959b95.jpg
   :alt: 擁抱開源，這個城市的法典都通過 GitHub 發佈
   :align: center

自然地，DC 條款網站也自動更新了該處失誤：

.. image:: https://oscimg.oschina.net/oscnet/b6a2396d2e7e03ab1bd175bb72b2f041312.jpg
   :alt: 擁抱開源，這個城市的法典都通過 GitHub 發佈
   :align: center

當然了，因為這是技術性錯誤 request，所以它可以被接受。

Joshua 認為，在 GitHub 上發佈法律條款是一項真正的政府創新，它改變了歷史，具體意義體現在這幾個方面：

  - 理事會能夠比以往更好、更快地公佈法律。以往律法每年發佈三次更新，並且可能會有五到七個月的延遲才能看到最新版本，但是，開源縮短這個過程，新的法律頒布約一個星期後民眾就可以看到。

  - 這也意味著，相比以往民眾能夠接收到更多法律條款。

  - 法律配套發佈在一個易用、現代化、可搜索、移動友好且免費的網站上，也就是上邊提到的 https://code.dccouncil.us/dc/council/code 。

  - 這給條件有限而沒法擁有更好的工具去研究律法的民眾和律師提供了更多機會。

同時 Joshua 也希望其它地方向 DC 學習，通過開源庫推進律法的發展。

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：擁抱開源，這個城市的法典都通過 GitHub 發佈
  | 本文地址：https://www.oschina.net/news/102155/dc-law-on-github

.. _Joshua Tauberer: https://razor.occams.info/
