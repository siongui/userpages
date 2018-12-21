Facebook 開源 DeepFocus，實現更逼真的 VR 圖像
#############################################

:date: 2018-12-21
:tags: 轉錄, 開源中國, 機器學習
:category: 轉錄
:author: h4cd(開源中國)
:summary: Facebook 開源了一種基於 AI 可實現更逼真 VR 圖像的系統 DeepFocus。
:og_image: https://oscimg.oschina.net/oscnet/ce23586bebbd45ba4365d3b90378654c9c9.jpg

Facebook 開源了一種基於 AI 可實現更逼真 VR 圖像的系統 DeepFocus_ 。

.. image:: https://oscimg.oschina.net/oscnet/ce23586bebbd45ba4365d3b90378654c9c9.jpg
   :alt: Facebook 開源 DeepFocus，實現更逼真的 VR 圖像
   :align: center

DeepFocus 可與高級原型頭戴設備配合使用，實時渲染模糊和各種焦距。例如，當有頭戴支持 DeepFocus 的設備觀看附近的物體時，它會立即聚焦並變得清晰，而背景物體則會失去焦點，這與現實生活中的感觀一樣。這種散焦模糊（也稱為視網膜模糊）對於實現 VR 中的真實感和深度感知非常重要。DeepFocus 是目前第一個能夠為 VR 應用實時生成此效果的系統。

一些傳統方法，例如使用累積緩衝器，可以實現物理上精確的散焦模糊。但它們無法實時產生複雜而豐富的內容，因為即使是最先進的芯片，處理需求也太高了。DeepFocus 使用深度學習解決了這個問題，團隊開發了一種新的端到端卷積神經網絡，眼睛看到場景的不同部分，就會產生具有精確視網膜模糊的圖像。該網絡包括新的保持體積的交織層，以在完全保留圖像細節的同時減小輸入的空間維度。然後網絡的卷積層以相同的、降低的空間分辨率運行，運行時間顯著減少。

由於 DeepFocus 僅依賴於標準 RGB-D 顏色和深度輸入，因此它幾乎適用於所有現有的 VR 遊戲和應用。

項目地址：https://github.com/facebookresearch/DeepFocus

更詳細的分析查看 `發佈公告`_ 。

..
  .. image:: 
   :alt: 
   :align: center

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：Facebook 開源 DeepFocus，實現更逼真的 VR 圖像
  | 本文地址：https://www.oschina.net/news/102888/facebook-opensource-deepfocus

.. _DeepFocus: https://code.fb.com/virtual-reality/deepfocus/
.. _發佈公告: https://code.fb.com/virtual-reality/deepfocus/
