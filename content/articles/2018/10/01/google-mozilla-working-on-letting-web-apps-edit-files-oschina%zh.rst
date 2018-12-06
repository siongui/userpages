Google 與 Mozilla 開發 API 讓 Web 應用輕鬆編輯本地文件
######################################################

:date: 2018-11-25T09:07+08:00
:tags: 轉錄, 開源中國, Web開發, 網路安全
:category: Web開發
:author: h4cd(開源中國)
:summary: 由 Google 和 Mozilla 領導的一個小組正在開發能夠使用基於瀏覽器的 Web 應用輕鬆編輯文件的 API，這可以讓用戶不用每次編輯文件後下載新文件。但是開發團隊認為這一功能會造成濫用，帶來安全問題，所以其希望獲得有關如何防範主要安全和隱私風險的建議。
:og_image: https://avatars1.githubusercontent.com/u/1540855


由 Google 和 Mozilla 領導的一個小組正在開發能夠使用基於瀏覽器的 `Web 應用輕鬆編輯文件的 API`_ ，這可以讓用戶不用每次編輯文件後下載新文件。但是開發團隊認為這一功能會造成濫用，帶來安全問題，所以其希望獲得有關如何防範主要安全和隱私風險的建議。

開發團隊認為這樣一個功能放在今天是很有必要的，用戶只需要在 Web 應用上做好保存，而不需要每次編輯本地文件後都重新下載它們。

谷歌開發者布道師 Pete LePage 指出：“今天，如果用戶想要在 Web 應用中編輯本地文件，那麼 Web 應用需要讓用戶打開該文件，在編輯文件後，保存更改的唯一方法是將文件下載下來。這種用戶體驗很糟糕，並且很去難構建訪問用戶文件的 Web 應用。”

為此，由 Chrome 開發人員和 Firefox 開發人員代表擔任主席的 W3C Web 孵化社區組織（WICG）正致力於開發新的 Writable Files API，該 API 允許在瀏覽器中運行的 Web 應用打開文件、編輯，並將更改保存到同一文件。

然而，該組織表示，最大的挑戰將是防止惡意網站濫用持久訪問用戶系統上的文件。“到目前為止，這個 API 最難的部分當然是要使用的安全模型”，WICG 的 API 解釋器頁面警告到：“API 可能被濫用並帶來許多可怕的攻擊。”

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：Google 與 Mozilla 開發 API 讓 Web 應用輕鬆編輯本地文件
  | 本文地址：https://www.oschina.net/news/102097/google-mozilla-working-on-letting-web-apps-edit-files

.. _Web 應用輕鬆編輯文件的 API: https://www.techrepublic.com/article/google-mozilla-working-on-letting-web-apps-edit-files-despite-warning-it-could-be-abused-in-terrible/
