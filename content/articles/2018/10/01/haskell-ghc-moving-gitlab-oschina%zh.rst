Haskell 開發者再次建議將 Haskell GHC 遷移至 GitLab
###################################################

:date: 2018-12-03T13:00+08:00
:tags: 轉錄, 開源中國, git
:category: 轉錄
:author: 局长(開源中國)
:summary: 昨日，Haskell 開發者 Ben Gamari 再度發佈郵件建議社區將 Haskell 編譯器項目 Haskell GHC 遷移至 GitLab。
:og_image: https://upload.wikimedia.org/wikipedia/commons/thumb/1/1c/Haskell-Logo.svg/1200px-Haskell-Logo.svg.png

昨日，Haskell 開發者 Ben Gamari `再度發佈郵件`_ 建議社區將 Haskell 編譯器項目 Haskell GHC 遷移至 GitLab。

Ben Gamari 曾於幾週前提過這個建議 —— 考慮將 Haskell GHC 的基礎開發設施遷移至 GitLab。不過當初他只提供了一小時的測試用例，這對於要說服其他人去使用 GitLab 來說顯然不夠。

而這次 Ben Gamari 明顯是有備而來，帶著體驗 GitLab 使用方式的網站( https://gitlab.staging.haskell.org )再次推廣他的提議。Ben Gamari 還列舉了一些遷移至 GitLab 後的好處：

- 可完整導入 Trac tickets，包括附件
- 通過 CircleCI 進行持續集成
- 對所有 boot 庫進行鏡像
- 可通過 GitHub 憑據進行登錄

但 Ben Gamari 也坦言，所有這些相關的功能均存在問題，不過目前處於正在解決中的狀態。

Ben Gamari 表示他的最終目標仍然是在12月18日切換到 GitLab，但現在尚未看到有相關開發者對他的這個提案做出回覆。

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：Haskell 開發者再次建議將 Haskell GHC 遷移至 GitLab
  | 本文地址：https://www.oschina.net/news/102336/haskell-ghc-moving-gitlab

.. _再度發佈郵件: https://mail.haskell.org/pipermail/ghc-devs/2018-December/016613.html
