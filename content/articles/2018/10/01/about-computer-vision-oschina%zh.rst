AlphaGo 早已擊敗圍棋冠軍，計算機視覺還是3 歲的“智力”
####################################################

:date: 2020-08-25
:tags: 轉錄, 開源中國, AI, 人工智能
:category: 轉錄
:author: 虎皮卷很乖(開源中國)
:summary: 儘管人工智能領域取得了令人難以置信的進步，但計算機視覺的應用仍有很長的路要走，因為距離計算機可以像人類一樣地去解釋圖像還需要很長時間。就像文章開頭提到的那樣，AlphaGo 早已擊敗了人類圍棋冠軍，計算機視覺的識圖能力僅相當於一個3歲的小孩。
:og_image: https://oscimg.oschina.net/oscnet/up-c83b1345841bc055e9e0ef5423009892d83.png

.. raw:: html

  <div>
  <p>20 世紀50 年代和60 年代，計算機視覺並沒有被看成重頭戲，人們認為視覺系統很容易複製，而教計算機下棋更加困難。但是現在，AlphaGo 已經擊敗圍棋冠軍，IBM Watson 也在Jeopardy 中擊敗人類競賽者，而大多數計算機視覺軟件最多只能完成3 歲兒童的任務&hellip;&hellip;</p>
  <p>理論與實踐不斷證明，人類視覺神經非常複雜，計算機視覺實現並非易事。計算機視覺研究從上世紀50 年代興起之後，也歷經了狂歡、冷靜，又重新燃起希望的階段。</p>
  <p>本篇編譯整理自計算機視覺相關文章，介紹計算機視覺各階段的理論支撐與外部輿論變化。</p>
  <h1>最早的人工智能想像</h1>
  <p>大多數人都認為，是現代人創造了人工智能的概念，實際上遠古祖先也提出了思考型機器人的理論。</p>
  <p>大約3000年前，荷馬描述了火神赫菲斯托斯（Hephaestus）的故事。赫菲斯托斯用黃金塑造了機械侍女，並賦予她們理性和學習能力。在無人駕駛汽車問世之前的幾個世紀，古希臘作家阿波羅紐斯（Apollonius）用想像力創造了塔洛斯（Talos），這是一個青銅自動機，負責保衛克里特島。</p>
  <p>但這些歷史性的敘述並不能準確描述當今正在開發的人工智能的種類。因為儘管現在大多數人工智能程序無法將其目標任務之外的知識概括化，但是對於本身給定的預期，他們已經達到或超過人類水平。</p>
  <h1>人的視覺與計算機視覺</h1>
  <p>人類視覺系統非常特別，超過50%的神經組織直接或間接地與視覺有關，其中超過66%的神經活動僅為處理視覺。今天人們對視覺、知覺的了解大部分來自1950年代和1960年代對貓進行的神經生理學研究。</p>
  <p>&nbsp;<img src="https://oscimg.oschina.net/oscnet/up-8a0af1883110b717053377ec1795037d2f0.JPEG" alt="" width="600" height="401" /></p>
  <p>上世紀50年代初期，哈佛醫學院的兩位神經科學家David Hubel 和Torsten Wiesel 在貓的視覺皮層上進行了一項實驗。通過研究神經元對各種刺激的反應，兩位科學家觀察到，人類的視覺是分級的。神經元檢測簡單的特徵，例如邊緣，然後輸入更複雜的特徵，如形狀，最後再輸入更複雜的視覺表示。基於這些知識，計算機科學家就可以專注於以數字形式重建人類神經結構。</p>
  <h1>早期樂觀</h1>
  <p>人工智能領域成立於1956年在達特茅斯學院舉行的夏季研討會上，來自不同領域的科學家們聚集在一起，以闡明並發展關於&ldquo;思維機器&rdquo;的觀點。</p>
  <p>1960年代，大學開始認真進行計算機視覺研究，並將該項目視為人工智能的奠基石。</p>
  <p>麻省理工學院的神經科學家戴維&middot;馬爾（David Marr）在70年代出版了一本《VISION》，匯集了一些方法，作出了可檢測的預測，提供解決神經科學問題的框架，並激發一代年輕科學家研究大腦及計算。該書提出了一種用於研究生物視覺系統的計算範例，並介紹信息處理系統的三個不同分析層次概念，即計算理論層次、表示形式和算法層次、以及實現層次。他們分別指向：計算的目標是什麼；解決問題、實現目標的陳述與流程；這些表示和過程的物理實例化，例如如何在神經元中完成特定任務。這三個層次劃分的意義是，研究者從視覺系統的宏觀表示出發進行思考，而不是查看如單個神經元式的微觀實體。</p>
  <p>Marr 建立了視覺表示框架，任何視覺系統感知到的強度，都是四個主要因素的函數：幾何形狀，意指形狀和相對位置；可見表面的反射率和絕對吸附特性，即物理特性；照明，即光源；相機，包含光學。</p>
  <p><img src="https://oscimg.oschina.net/oscnet/up-da41dc521897d807cb11ea9a93cbb71f647.png" alt="" width="600" height="252" /></p>
  <p>早期研究人員對相關領域的未來非常樂觀，並提倡把人工智能作為一種可以改變世界的技術。一些人預測，一代人的時間內就會創造出像人類一樣聰明的機器，這種炒作為研究人員贏得了數百萬美元的公共和私人資金，研究中心在全球湧現。但是由於接下里的發展未能達到人們的高預期，國際上的人工智能開發工作受到了限制。</p>
  <h1>AI 視覺的冬天</h1>
  <p>研究人員以極大的樂觀度，將公眾的期望提高到了令人難以置信的高度，卻未能體會到他們為自己設定的挑戰的難度。當研究人員承諾的預期未能達成時，這個研究領域遭到了猛烈的批評，和嚴重的財務挫折。</p>
  <p>早期的計算資源在技術上無法跟上科學家提出的複雜問題，即使是最令人印象深刻的項目也只能解決微不足道的問題。此外，大多數研究人員都是在孤立的小組內工作，缺乏比較有意義的，可以推進該領域科學進展的方式。</p>
  <p>有一則故事可以反映當時研究人員的美好預期、以及預期落空後他們自身及和外界的失落與嘲諷。</p>
  <p>1966年，美國計算機科學家、麻省理工學院AI 實驗室聯合創始人馬文&middot;明斯基（Marvin Minsky）獲得了暑期津貼，聘請了一年級的本科生杰拉爾德&middot;蘇斯曼（Gerald Sussman） ，讓他花費整個夏天的時間把一台攝像機與計算機連接起來，並讓計算機描述它看到了什麼。&ldquo;不用說，蘇斯曼沒有在截止日期前完成，&rdquo;Motion Metrics 的機器學習開發人員Hooman Shariati 曾說，&ldquo;在接下來的四十年中，視覺成為人工智能領域最困難、最令人沮喪的挑戰之一。正如機器視覺專家貝特霍爾德&middot;霍恩（Berthold Horn）曾經指出的那樣，蘇斯曼選擇不再在視覺領域工作。&rdquo;</p>
  <p>到70年代中期，政府和公司對人工智能失去了信心，行業資金枯竭。數學家詹姆斯&middot;萊特希爾（James Lighthill）1973年發表了一篇論文，批評早期人工智能研究，這為後來英國政府撤回對該領域的支持奠定了研究基礎。</p>
  <p>隨後的這段時間被稱為&ldquo;人工智能的冬天&rdquo;。雖然20世紀80年代和90年代研究還在繼續，也有過一些小規模的複興，但人工智能基本上被被歸入了科幻小說的範疇，嚴肅的計算機科學家都避免使用這個詞。</p>
  <h1>卷積神經網絡出現與多倫多大學的突破</h1>
  <p>隨著互聯網成為主流，計算機科學家有了可以訪問更多數據的權限。計算機硬件在繼續改進，成本則在下降。80年代到90年代，基本神經網絡和算法得到改進。</p>
  <p>1998年，Bengio、Le Cun、Bottou 和Haffner 在一篇論文中首次介紹了第一個卷積神經網絡LeNet-5，能夠分類手寫數字。</p>
  <p>卷積神經網絡可以做到平移不變形，即使對象的外觀發生某種方式的變化，也可以識別出對象。卷積神經網絡通過監督學習和反向傳播對輸入到卷積網絡中的數據做訓練，並反复、自我校正。和同樣可以做反向傳播的深度神經網絡相比，卷積神經網絡的特殊之處在於神經元之間的鏈接結構和獨特的隱藏架構的方式，這是由人類視覺皮層內部的視覺數據處理機制啟發得來的。此外，CNN 中的圖層按照寬度、高度和深度三個維度進行組織。</p>
  <p>卷積網絡最重要的屬性之一就是，不管有多少層，整個CNN 系統僅由兩個部分組成：特徵提取和分類。通過對特定特徵的選擇，以及通過前饋鏈接增加空間不變性，這也是人工視覺系統如CNN 非常獨特的原因。</p>
  <p><img src="https://oscimg.oschina.net/oscnet/up-c83b1345841bc055e9e0ef5423009892d83.png" alt="" width="600" height="286" /></p>
  <p>（視覺皮層和卷積神經網絡有許多相似）</p>
  <p>深度神經網絡研究應用也有進步，並且使人們信心大增。2012年，人工智能在ImageNet 大規模視覺識別挑戰（ILSVRC）上取得突破。</p>
  <p>ILSVRC 是一個年度圖像分類比賽，研究團隊在給定的數據集上做視覺識別任務，評估算法準確性。2010年和2011年，ILSVRC 獲獎者的錯誤率一直在26%左右。2012年，來自多倫多大學的團隊帶來一個名為AlexNet 的深度神經網絡，實現了16.4%的錯誤率。在接下來的幾年中，ILSRVC 的錯誤率下降到了幾個百分點。</p>
  <h1>基於AI 的計算機視覺的未來</h1>
  <p>當下，人們已經知道，視覺能力是人類承擔的生物學生最複雜的任務之一，對計算機視覺的研究和預期也更加貼近世界。同時，基於對人類視覺能力了解的深入，計算機視覺研究人員也在不斷更新算法和理論。</p>
  <p>CNN 已經廣泛用於需要處理視覺和空間信息的系統中。但隨著人工智能需要解決更高級的問題，對計算和電力資源的增長需求成為CNN 最突出的問題之一。研究人員的注意力也在逐漸轉向尖峰神經網絡SNN，這是一種新型的ANN ，受大腦神經動力學的啟發，具有事件驅動，快速推理和省電的特性，也被認為是第三代神經網絡。</p>
  <p>接下來SNN 要優化解決的一個問題是視覺注意VA 與智力。人類可分散的注意力使得人能同時執行多個任務，注意力轉移可以使人快速訪問新信息。視覺注意力研究的核心目標是要使處理的視覺信息量最少，以解決複雜的高級任務，例如對象識別。</p>
  <p>計算機視覺任務主要涉及處理靜態圖像，人類眼睛在檢測到場景變化向大腦傳遞信息&mdash;&mdash;這是一個事件，生物視覺系統的這一關鍵特性允許將注意力選擇性地集中在場景的顯著部分上，從而大量減少需要處理的信息量。</p>
  <p>假設針對一張人在草地上打高爾夫的圖像。傳統的傳感器中，數據以幀的形式傳輸，圖像上的所有內容都要經過處理，而重要的信息是人的運動，以及帶動的球桿和球的運動。這時，基於事件的傳感器並不會讀取每個像素並且以恆定速率發送幀，而是在檢測到像素局部亮度變化是，從每個像素異步發送數據包或事件，從而減少計算、傳輸的數據和功耗。</p>
  <p>研究人員認為，CNN 非常適合靜態圖像中的對象識別，但它缺乏動態特性來處理基於事件的傳感器的實時數據集。因此，SNN 被寄予厚望。</p>
  <p>現在人工智能已經無縫集成到日常生活的多方面。研究人員表示，近年來，人工智能在許多研究領域都取得了巨大的成功。像AlphaGo 這樣的遊戲系統已經使用強化學習來自學，助聽器使用深度學習算法過濾掉環境噪音，這些技術甚至為自然語言處理與翻譯、對象識別以及模式匹配系統提供了動力，我們已經對谷歌、亞馬遜、 iTunes 等提供的類似服務習以為常。這種趨勢也絲毫沒有放慢的跡象，人們可以用計算機自動化執行許多小的重複性任務以節省時間。</p>
  <p>儘管人工智能領域取得了令人難以置信的進步，但計算機視覺的應用仍有很長的路要走，因為距離計算機可以像人類一樣地去解釋圖像還需要很長時間。就像文章開頭提到的那樣，AlphaGo 早已擊敗了人類圍棋冠軍，計算機視覺的識圖能力僅相當於一個3歲的小孩。</p>
  <p>參考鏈接：</p>
  <ul>
  <li><a href="https://www.motionmetrics.com/how-artificial-intelligence-revolutionized-computer-vision-a-brief-history/" target="_blank" rel="noopener">https://www.motionmetrics.com/how-artificial-intelligence-revolutionized-computer-vision-a-brief-history/</a></li>
  <li><a href="https://becominghuman.ai/from-human-vision-to-computer-vision-how-far-off-are-we-part1-3-b35d37a196a4" target="_blank" rel="noopener">https://becominghuman.ai/from-human-vision-to-computer-vision-how-far-off-are-we-part1-3-b35d37a196a4</a></li>
  <li><a href="https://becominghuman.ai/from-human-vision-to-computer-vision-a-brief-history-part2-4-fcb1565d5492" target="_blank" rel="noopener">https://becominghuman.ai/from-human-vision-to-computer-vision-a-brief-history-part2-4-fcb1565d5492</a></li>
  <li><a href="https://becominghuman.ai/from-human-vision-to-computer-vision-convolutional-neural-network-part3-4-24b55ffa7045" target="_blank" rel="noopener">https://becominghuman.ai/from-human-vision-to-computer-vision-convolutional-neural-network-part3-4-24b55ffa7045</a></li>
  <li><a href="https://becominghuman.ai/from-human-vision-to-computer-vision-towards-spiked-based-visual-intelligence-and-neuromorphic-913e5de21bf9" target="_blank" rel="noopener">https://becominghuman.ai/from-human-vision-to-computer-vision-towards-spiked-based-visual-intelligence-and-neuromorphic-913e5de21bf9</a></li>
  </ul>
  </div>


.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自OSCHINA社區[http://www.oschina.net]
  | 本文標題： AlphaGo早已擊敗圍棋冠軍，計算機視覺還是3歲的“智力”
  | 本文地址：https://www.oschina.net/news/118149/about-computer-vision

