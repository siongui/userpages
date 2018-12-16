Facebook 開源 NLP 建模框架 PyText，從研究到生產變得更容易
#########################################################

:date: 2018-12-16T15:29+08:00
:tags: 轉錄, 開源中國, 機器學習
:category: 轉錄
:author: h4cd(開源中國)
:summary: Facebook AI Research（FAIR）開源了 NLP 建模框架 PyText。
:og_image: https://oscimg.oschina.net/oscnet/8dcc8a28d530e211f3b47b4ff4afe3f738a.jpg

Facebook AI Research（FAIR） `開源了 NLP 建模框架 PyText`_ 。

.. image:: https://oscimg.oschina.net/oscnet/8dcc8a28d530e211f3b47b4ff4afe3f738a.jpg
   :alt: Facebook 開源 NLP 建模框架 PyText，從研究到生產變得更容易
   :align: center

PyText 是一個基於 PyTorch 構建的深度學習 NLP 建模框架。PyText 通過為模型組件提供簡單且可擴展的接口和抽象，以及使用 PyTorch 的 Caffe2 執行引擎導出模型進行推理的功能，模糊了實驗與大規模部署之間的界限。其預訓練模型包括文本分類、序列標註等。

PyTorch 是一個統一的框架，縮短了從研究到生產的路徑，而基於 PyTorch 的 PyText 則著眼於滿足 NLP 建模的特定需求。

核心特性：
++++++++++

適用於各種 NLP/NLU 任務的生產就緒模型

* 文本分類

  - `Yoon Kim (2014): Convolutional Neural Networks for Sentence Classification`_
  - `Lin et al. (2017): A Structured Self-attentive Sentence Embedding`_

* 序列標註

  - `Lample et al. (2016): Neural Architectures for Named Entity Recognition`_

* 聯合意圖時隙模型（Joint intent-slot model）

  - `Zhang et al. (2016): A Joint Model of Intent Determination and Slot Filling for Spoken Language Understanding`_

* 上下文意圖-時隙模型（Contextual intent-slot models）

支持在 PyTorch 1.0 中基於新 C10d 後端構建的分佈式訓練

可擴展組件，可輕鬆創建新模型和任務

參考實現和預訓練模型論文： `Gupta et al. (2018): Semantic Parsing for Task Oriented Dialog using Hierarchical Representations`_

支持聯合訓練



項目地址：https://github.com/facebookresearch/pytext

瞭解更多：https://code.fb.com/ai-research/pytext-open-source-nlp-framework/

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：Facebook 開源 NLP 建模框架 PyText，從研究到生產變得更容易
  | 本文地址：https://www.oschina.net/news/102736/facebook-opensource-pytext

.. _開源了 NLP 建模框架 PyText: https://code.fb.com/ai-research/pytext-open-source-nlp-framework/
.. _Yoon Kim (2014)\: Convolutional Neural Networks for Sentence Classification: https://arxiv.org/abs/1408.5882
.. _Lin et al. (2017)\: A Structured Self-attentive Sentence Embedding: https://arxiv.org/abs/1703.03130
.. _Lample et al. (2016)\: Neural Architectures for Named Entity Recognition: https://www.aclweb.org/anthology/N16-1030
.. _Zhang et al. (2016)\: A Joint Model of Intent Determination and Slot Filling for Spoken Language Understanding: https://www.ijcai.org/Proceedings/16/Papers/425.pdf
.. _Gupta et al. (2018)\: Semantic Parsing for Task Oriented Dialog using Hierarchical Representations: http://aclweb.org/anthology/D18-1300
