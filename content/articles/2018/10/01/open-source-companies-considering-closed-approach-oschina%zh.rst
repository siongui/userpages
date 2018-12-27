開源面臨生死存亡之際
####################

:date: 2018-12-24T10:03+08:00
:tags: 轉錄, 開源中國
:category: 轉錄
:author: CSDN资讯
:summary: 開源軟件在它的頂峰遇到了一場存亡危機。
:og_image: https://img-blog.csdnimg.cn/20181221141813595.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2NzZG5uZXdz,size_16,color_FFFFFF,t_70


.. highlights::

  開源軟件在它的頂峰遇到了一場存亡危機。


.. image:: https://img-blog.csdnimg.cn/20181221141813595.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2NzZG5uZXdz,size_16,color_FFFFFF,t_70
   :alt: 開源面臨生死存亡之際
   :align: center

毫無疑問，開源軟件的概唸給企業軟件的世界帶來了革命，後者花費了幾十億與開源的概念鬥爭了多年後，才不得不接受了新的未來。但越來越多的人開始擔心，開源軟件允許任何人幹任何事的本質可能會在分佈式雲計算的時代給開發者帶來巨大問題。

上次我們也討論了這個問題並發現，業界有兩家重要的開源軟件公司（Redis 和 MongoDB）決定改變它們的部分軟件發佈時採用的授權。此舉表明它們試圖讓雲計算提供商在它們的軟件基礎上提供服務變得更困難，甚至不可能。

兩家公司當然不能撼動整個世界。但隨著當前許多雲計算公司已經安排好了 2019 年的計畫，開源項目和雲計算服務之間的交集開始成為許多人擔心的問題。

.. image:: https://img-blog.csdnimg.cn/20181221141824720.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2NzZG5uZXdz,size_16,color_FFFFFF,t_70
   :alt: 開源面臨生死存亡之際
   :align: center

.. container:: align-center

  Cloud Foundry Foundation 的執行主管 Abby Kearns 在 2017 GeekWire Cloud Tech Summit 上演講

Cloud Foundry Foundation 的開源執行主管 Abby Kearns 說：“我認為，開源軟件在建立商業機會的過程中扮演的角色已經變了。這種擔憂只會越來越多。”


一、改變是為了保護
++++++++++++++++++

“直白地說，這些年來我們都很傻，讓他們白白使用我們開發的東西並賺了很多錢。”

Redis Lab 的 CEO Ofer Bengal 並沒有拐彎抹角。作為著名的開源內存數據庫的開發商，他的公司已經存在了八年之久，這在現代企業軟件的飛速變革的世界中算是非常長壽了。

雲計算在 2011 年剛剛起步，但直到現在，它仍然是許多無法負擔幾百萬服務器費用的早期創業者在嘗試新想法時的首選。絕大多數成熟的公司依然使用傳統方式構建自己的基礎設施，但越來越多的公司開始意識到，開源軟件能讓他們用比傳統企業軟件公司的私有軟件更靈活、成本更低的方式構建基礎設施。

.. image:: https://img-blog.csdnimg.cn/2018122114184031.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2NzZG5uZXdz,size_16,color_FFFFFF,t_70
   :alt: 開源面臨生死存亡之際
   :align: center

.. container:: align-center

  Ofer Bengal，Redis Lab 的創始人和 CEO

Redis 在這一時期變得非常流行，主要的大公司如美國運通、家得寶、夢工廠等都在 Redis 數據庫的基礎上建立了他們的技術基礎設施。Redis Lab 提供的產品 Redis Enterprise 能將該數據庫在公有雲上作為服務提供，或者為在自有基礎設施上運行該數據庫的公司提供技術支持。而且他們依然在不斷為 Redis 開源項目做出貢獻。

AWS 經常聲稱它的主要目標是客戶，但從這個目標中也能看出，為了給客戶提供相似的服務，哪些技術和服務開始受到關注。它於 2013 年啟動了一項雲服務，該服務基於開源的 Redis，由 AWS 管理。

Bengal 說，從那時起，AWS 通過為客戶提供 Redis 賺取了“幾億美元”，但並沒有像開源社區做出相當數量的貢獻幫助構建並維護該項目。很難得知 AWS 究竟賺了多少錢，但顯然 AWS 和其他雲服務提供商都從開源開發者的勞動成果中獲得了好處。

Bengal 說：“Redis 的 99% 的貢獻都來自 Redis Labs。”Puppet 的創始人 Luke Kanies 在今年早些時候的一篇報導中說過，開源軟件界的傳說一直是“項目由社區中的貢獻者驅動”，但實際上，絕大多數現代開源項目中的絕大部分代碼都是由領工資的開發者貢獻的。

而這些人的工資得有人提供。很長一段時間，Redis 就是成功的開源商業模型的典範，在維持最基礎的項目的同時，在其上開發自己的軟件並提供訪問。但隨著越來越多的公司開始使用雲計算並且將自己現有的應用程序和基礎設施轉移到 AWS 等雲服務提供商上，更現實的做法是與 EC2、S3 等其他 AWS 服務一起，使用 AWS 的 Redis 服務，而不是使用 Redis 通過 AWS 市場提供的 Redis 服務。

Bengal 說：“不僅是我們，幾乎任何成功的開源項目都會遇到這個問題。”儘管由於巨大的市場影響力，人們在談到這些問題時總會提到 AWS，但它絕不是唯一一家將這種開源項目作為服務提供的。

所以在八月，Redis 決定（ https://redislabs.com/community/licenses/ ）將他們在 Redis 的基礎上建立的數據庫擴展（但不包括 Redis 本身）的授權改成 Commons Clause 授權，這個授權表明其他公司不能將這些擴展作為雲服務提供。

“我們保留決定每個軟件採用寬容的開源授權或 Commons Clause 授權的自由。”Bengal 說。“基本上，這是個商業決定。”

然後在十月，另一個註明的開源軟件數據庫做出了類似的決定。MongoDB 宣佈（ https://www.mongodb.com/press/mongodb-issues-new-server-side-public-license-for-mongodb-community-server ），以後會對MongoDB Community Server 軟件採取不同的授權，名為SSPL（ https://www.mongodb.com/licensing/server-side-public-license ），該授權允許雲服務商將 MongoDB 作為服務提供，但要求他們開源所有創建該服務時編寫的源代碼，否則就必須與 MongoDB 達成商業協議。

“每當新的開源項目變得流行後，雲服務商就會竊取技術，將自由軟件放在他們的平台上，攫取絕大部分軟件的利益，但很少回饋社區。”MongoDB 的 CEO Dev Ittycheria 說。MongoDB 目前在納斯達克上的市值為 43 億美元。“我們認為，由我們這樣的公司來引領並幫助下一代開源公司和開源項目的生存和成長是非常重要的。”

兩家數據庫公司引導這次變革絕不是偶然。數據庫是極其複雜困難的項目，而且是任何大規模的企業級公司最重要的組成部分。

Ittycheria 估計，多年來 MongoDB 在研發上花費了超過 1 億 5 千萬美元，才得以建立並維護 MongoDB 的開源版本。在上一個財政年度， MongoDB 從它的商業軟件和支持的服務中獲取了 1 億 5450 萬美元的收入。

“我們的觀點是，開源軟件從來不應該讓雲服務公司拿去賣。”在 Redis 宣佈了它的決定後，Bain Capital Ventures 的管理總監 Salil Deshpande 在一篇 Techcrunch 上的文章中說。Deshpande 是 Redis Labs 的投資人之一，他幫助開源軟件公司撰寫了我們三月份看到的 Commons Clause 授權。

隨著開源軟件成為企業軟件世界中的重要組成部分，它的方向和許多方面不可避免地會受到商業意向的影響。現在的問題是，成為開源軟件的意義是什麼？那些利用別人開發的開源項目來提供服務並從中獲利的公司，他們欠開源項目的建立者和維護者什麼？


二、開放和封閉
++++++++++++++

在考慮第一個問題時，重要的是要考慮到 Redis 的 Commons Clause 授權顯然不是一份開源授權，這一點所有人都同意。Redis 在廣為人知的 BSD 授權下依然是個開源軟件，但 Redis 公司為它在 Redis 項目基礎上開發的擴展應用了 Commons Clause 授權。

MongoDB 的情況有點不一樣。因為 SSPL 授權要求雲服務商在將開源項目做成服務時，提供更多的開源軟件作為匯報，因此 MongoDB 公司認為，這一點不違背開源軟件的精神。

“每個人都希望有更多的開源軟件，但總要有人為之提供資助。而為了保證資金，就必須保證在商業上有存在的價值。”Ittycheria 說。

位於西雅圖的 Chef（ https://www.chef.io/ ）的聯合創始人兼 CTO Adam Jacob 維護著幾個開源項目，這些項目的目的是讓基礎設施和應用程序的管理更容易。Adam 很懷疑這是否是開源軟件項目、開源開發者和開源公司的正確發展方向。

“我不覺得這是在社區的基礎上做出的決定，當然這次事件跟歷史上的開源軟件和免費軟件的誕生也不一樣。”Jacob 說。“我不覺得應該出現像‘我們的業務需要商業的保駕護航’這種三條腿的決定。”

.. image:: https://img-blog.csdnimg.cn/20181221141916690.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2NzZG5uZXdz,size_16,color_FFFFFF,t_70
   :alt: 開源面臨生死存亡之際
   :align: center

.. container:: align-center

  Chef 的聯合創始人兼 CTO Adam Jacob

Jacob 說，畢竟，當開源項目開始賺錢時，Redis 和 MondoDB 等公司很樂意看到開源社區的友好、開放的氛圍，從而能將他們的作品廣泛傳播到全世界。

“更有意思的是，從開源商業模型的角度來看，他們所作所為中開源的部分一直都是他們商業模型中的一部分。更重要的是開發者能觸及的深度。”Jacob 說，他們的主要目的是建立一個“漏斗策略”，利用免費版本吸引用戶，然後銷售商業版本。

但對於許多小公司，參與開源社區並維護項目是與成熟的公司競爭的唯一辦法。對於年輕的企業軟件公司來說，最難的事情之一就是說服其他公司為你的產品付費。

“開源軟件提供了許多機會，特別是為創業公司。”Kearns 說。但在某一點上，早期的戰略決策可能會成為沉重的負擔；另一個很難的事情就是說服習慣了免費的人為之付費。


三、地平線上的雲
++++++++++++++++

這一點可能是最重要的：開發者作為個人愛好而開發開源軟件的時代早已終結，而雲服務商能通過開源軟件攫取收益而不用付出任何貢獻的今天，通過社區的方式來開發任何人皆可使用的軟件，是否還有存在的價值？

Jacob 認為有。

“我相信我們的錯誤在於，我們實際上不再信任公眾，不再信任自由軟件的價值是構建更好的社區必須的，因為我們在一開始就認為，商業和社區是截然分開的。”他說。

Redis 和 MongoDB 認為，他們並沒有將開源社區棄之不顧。他們認為，他們別無選擇，必須找到新的途徑從財務方面支持他們的開發者，這些開發者對於社區的健康發展十分重要，儘管部分工作依然屬於私有財產，或者是有條件的開放。

最大的三家雲服務商對於該問題依然保持沉默，拒絕讓高管們討論這次開源授權的變動。Google 從早期就在布道開源軟件的價值，而微軟和 AWS 在新世界中採取了不同的途徑。

當了多年的開源頭號公敵之後，微軟開始擁抱開源軟件，開始僱傭擁有雄厚的開源經驗的開發者，並對一些社區做出了至關重要的回饋。AWS 在與開源社區合作的方面比較緩慢，但在過去幾年裡也在逐漸改變其基調，招募了許多開源開發者（如 James Gosling 和 Adrian Cockcroft）來改變他們對於開源貢獻的看法。

如果更多的公司切換到這種更激進的授權，就會強迫雲服務商改變產品開發戰略，因為他們必須評價哪些服務更值得進入商用，哪些不值得。但是現在，還有許多創業公司和項目依然在使用傳統的開源授權，意味著雲服務商們還在觀望市場對於 Redis 和 MongoDB 的反應。

但現代企業在競爭激勵的二十一世紀中構建並管理技術的一切前提都已經被雲計算改變了。因此不難想像，雲計算也可能會改變開源軟件開發背後的前提。


..
  .. image:: 
   :alt: 
   :align: center

.. highlights::

  | 原文：https://www.geekwire.com/2018/open-source-companies-considering-closed-approach/
  | 作者：Tom Krazit，GeekWire 的雲和企業話題編輯，負責 IDG、CNET 等新聞組織的技術話題。
  | 本文為 CSDN 翻譯，如需轉載，請註明來源出處。
  |
  | 作者：CSDN资讯 
  | 来源：CSDN 
  | 原文：https://blog.csdn.net/csdnnews/article/details/85161366 
  | 版权声明：本文为博主原创文章，转载请附上博文链接！

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：開源面臨生死存亡之際
  | 本文地址：https://www.oschina.net/news/102968/open-source-companies-considering-closed-approach

----

`开源面临生死存亡之际！ - CSDN资讯 - CSDN博客 <https://blog.csdn.net/csdnnews/article/details/85161366>`_

