不想雲廠商坐收漁翁之利，Kafka 團隊修改 KSQL 開源許可
####################################################

:date: 2018-12-16T09:12+08:00
:tags: 轉錄, 開源中國
:category: 轉錄
:author: h4cd(開源中國)
:summary: 15 日，Confluent 宣佈修改其平台部分組件的開源許可，從 Apache 2.0 切換到 Confluent Community License。
:og_image: https://oscimg.oschina.net/oscnet/02246128604c0125c076efcb3ae85bb1432.jpg

15 日，Confluent 宣佈 `修改其平台部分組件的開源許可`_ ，從 Apache 2.0 切換到 `Confluent Community License`_ 。

.. image:: https://oscimg.oschina.net/oscnet/02246128604c0125c076efcb3ae85bb1432.jpg
   :alt: 不想雲廠商坐收漁翁之利，Kafka 團隊修改 KSQL 開源許可
   :align: center

新的 Confluent 社區許可允許用戶免費下載、修改和重新發行代碼，這點類似於 Apache 2.0，但是不允許將其作為 SaaS 產品提供給用戶。流數據 SQL 引擎 KSQL 將受到新許可的影響，而由於 Kafka 是 Apache 軟件基金會的一部分，它將繼續使用 Apache 2.0 許可，此次修改方案只會影響到由 Confluent 維護的開源組件。

.. image:: https://oscimg.oschina.net/oscnet/f6060fec68a2b76ff563c99eee49a75d400.jpg
   :alt: 不想雲廠商坐收漁翁之利，Kafka 團隊修改 KSQL 開源許可
   :align: center

Confluent 解釋，新的授權方式下，用戶仍然可以將 KSQL 等應用作為產品或服務的一部分，無論這些產品是作為軟件發行還是作為 SaaS 服務提供給用戶都可以，但不能用它創建類似“KSQL 即服務”這樣的東西。

至於為什麼這麼做，Confluent 表示，雲廠商將產品作為 SaaS 提供給用戶，他們具有顯著的優勢，可以控制資源的定價，並且可以在他們的所有產品中集成自己的服務。而在這其中，有一些雲廠商會與開源公司合作，由開源公司提供雲系統的託管版本，並作為服務提供給用戶。但是其它公司則直接將開源代碼放到他們的雲產品中，並投入資金開發差異化的專有產品。Confluent 認為這無可非議，作為一家公司，可以考慮構建更多的專有軟件，並減少開源方面的投入。但是從其自己的角度來看，構建基礎設施層的正確方式應該是使用開源代碼，而隨著工作負載不斷遷移到雲端，需要有一種機制在保持自由的同時實現投資週期。

Confluent 表示，組件的開發仍然是開放的，並且繼續接受 PR 和功能建議，同時其還會繼續在開發上大量投入。而對於那些非商業雲提供商用戶來說，新許可並沒有實質上的限制。

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：不想雲廠商坐收漁翁之利，Kafka 團隊修改 KSQL 開源許可
  | 本文地址：https://www.oschina.net/news/102737/license-changes-confluent-platform

.. _修改其平台部分組件的開源許可: https://www.confluent.io/blog/license-changes-confluent-platform
.. _Confluent Community License: https://www.confluent.io/confluent-community-license
