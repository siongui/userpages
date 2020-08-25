歷史上第一個聊天機器人：200 行代碼就能欺騙人類感情
##################################################

:date: 2020-08-17
:tags: 轉錄, 開源中國, AI, 人工智能
:category: 轉錄
:author: 摆摊卖报纸(開源中國)
:summary: 在1964年至1966年間，麻省理工學院人工智能實驗室的德裔美國計算機科學家約瑟夫·維森鮑姆（Joseph Weizenbaum）開發了歷史上第一個聊天機器人  —— Eliza。
:og_image: https://oscimg.oschina.net/oscnet/up-890ada733b7ef37ad860d5e03b97cbade71.png

.. raw:: html

  <div>
  <p>在1964年至1966年間，麻省理工學院人工智能實驗室的德裔美國計算機科學家約瑟夫&middot;維森鮑姆（<a href="https://en.wikipedia.org/wiki/Joseph_Weizenbaum" target="_blank" rel="noopener">Joseph Weizenbaum</a>）開發了歷史上<a href="https://web.stanford.edu/class/linguist238/p36-weizenabaum.pdf" target="_blank" rel="noopener">第一個聊天機器人</a> &mdash;&mdash; Eliza。</p>
  <p>&nbsp;<img src="https://oscimg.oschina.net/oscnet/up-890ada733b7ef37ad860d5e03b97cbade71.png" alt="" width="500" height="375" />&nbsp;</p>
  <p>Eliza 的名字源於愛爾蘭劇作家蕭伯納的戲劇作品《賣花女》中的角色，劇中出身貧寒的賣花女Eliza 通過學習與上流社會溝通的方式，變成大使館舞會上人人艷羨的&ldquo;匈牙利王家公主&rdquo;。作為世界上第一個聊天機器人，Eliza 被其作者賦予了充滿戲劇性的內涵。</p>
  <p>儘管在當時已經存在一些基本的數字語言生成器（可以輸出一些連貫文本的程序），但Eliza 是第一個明確設計用於與人互動的程序。用戶可以使用打字機輸入人類的自然語言，獲得來自機器的響應。正如維森鮑姆解釋的那樣，Eliza 使&ldquo;人與計算機之間的對話成為可能&rdquo;。&nbsp;</p>
  <p>Eliza是維森鮑姆用一種名為MAD-SLIP的編程語言，在36位的 <a href="https://en.wikipedia.org/wiki/IBM_7090#IBM_7094" target="_blank" rel="noopener">IBM 7094</a>（一種早期的晶體管大型計算機）上開發的。在那個年代裡，計算機時間是非常寶貴的資源，所以Eliza只能通過分時系統運行，用戶只能通過電動打字機和打印機與程序進行遠程交互。當用戶鍵入一個句子並按Enter鍵時，一條消息會被發送到遠端的大型機。Eliza掃描消息中是否存在關鍵字，並將這些關鍵字組成新的句子，用以回复響應，打印出來並供用戶閱讀。&nbsp;</p>
  <p>為了鼓勵用戶與Eliza 對話，維森鮑姆使用&ldquo;狡猾策略&rdquo;為Eliza 設計了一種模擬羅傑式心理醫生（一種以人為本的心理諮詢）的對話風格。該程序將接受用戶所說的內容，並將以一個問題作為回應，如下面的示例所示。（可以注意該程序如何接收&ldquo;男朋友&rdquo;和&ldquo;沮喪&rdquo;之類的關鍵字並重新使用它們。）&nbsp;</p>
  <p>人：男人都是一樣的。</p>
  <p>Eliza：怎麼說？</p>
  <p>人：他們總是在困擾我們一些事情。</p>
  <p>Eliza：您可以舉一個具體示例嗎？</p>
  <p>人：好吧，我的男朋友讓我來了這裡。</p>
  <p>Eliza：您的男朋友讓您來了這裡？</p>
  <p>人：他說我很多時候都很沮喪。</p>
  <p>Eliza：很抱歉聽到您感到沮喪。</p>
  <p>人：是的，我不開心。&nbsp;</p>
  <p>Eliza 的這種對話方式，給人的印象就是計算機可以理解用戶所講的內容，而不必為對話提供任何新的補充。事實上，維森鮑姆僅用200 行代碼就實現了這一過程，使Eliza 能夠讓人產生被機器聆聽、理解的錯覺。&nbsp;</p>
  <p>為了測試Eliza 與對話者互動的能力，維森鮑姆邀請學生和同事進入他的辦公室，並讓他們在他的觀察下與機器聊天。他有些擔憂地註意到，在與Eliza 進行短暫互動期間，許多測試者竟開始對該程序產生情感上的依戀。他們會向機器敞開心，坦承他們在生活和人際關係中面臨的問題。&nbsp;</p>
  <p>更令人驚訝的是，甚至在維森鮑姆向他們介紹了Eliza 的工作原理，並解釋說它並不是真正理解了用戶所說的任何內容之後，測試者對Eliza 的這種親密感仍然存在。尤其是魏森鮑姆的助手，儘管她一直看著該程序從零開始構建的全過程，但在測試時，這位助手仍然堅持要魏森鮑姆離開房間，以便她可以與Eliza 私下交談。&nbsp;</p>
  <p>通過Eliza的實驗，魏森鮑姆開始質疑阿蘭&middot;圖靈（Alan Turing）在1950年提出的關於人工智能的想法。圖靈在他的題為&ldquo; <a href="https://academic.oup.com/mind/article/LIX/236/433/986238" target="_blank" rel="noopener">計算機械與智能</a> &rdquo;的論文中提出，如果一台計算機可以通過文本與人類進行令人信服的對話，則可以認為它是智能的。這一思想也就是著名的圖靈測試的基礎。&nbsp;</p>
  <p>但是Eliza 的測試證明，即使人機之間的理解只從人類這一側產生，也可以在人機之間進行令人信服的對話。也就是說，對人類智能的模擬（而不是智能本身）足以使人蒙昧。維森鮑姆稱這種現象為&ldquo;Eliza 效應&rdquo;，並認為這是數字時代人類共同遭受的一種&ldquo;妄想&rdquo;。這一見解對維森鮑姆來說是一次深刻的衝擊，並直接影響了他在未來十年裡所做研究的思想軌跡。&nbsp;</p>
  <p>1976年，魏森鮑姆發表了<a href="http://blogs.evergreen.edu/cpat/files/2013/05/Computer-Power-and-Human-Reason.pdf" target="_blank" rel="noopener">《</a><a href="http://blogs.evergreen.edu/cpat/files/2013/05/Computer-Power-and-Human-Reason.pdf" target="_blank" rel="noopener">計算能力與人為原因：從判斷到計算》</a> ，該書對人們為何願意相信&ldquo;一台簡單的機器也許能夠理解複雜的人類情感&rdquo;進行了深刻的剖析。</p>
  <p>他在這本書中認為，&ldquo;Eliza 效應&rdquo; 代表著一種困擾現代人的廣泛病理學。在一個被科學技術和資本主義所佔領的世界中，人們已經習慣於將自己看作是一台大型且冷漠的社會機器中一枚孤立的齒輪。魏岑鮑姆認為，正是由於當時的社會環境日漸冷漠，才使得人們變得如此絕望，以至於拋棄應有的理性和判斷力，轉而去相信一個機器程序可以聆聽他們的心聲。</p>
  <p>魏森鮑姆的餘生都在致力於這種對人工智能和計算機技術的人文主義批評。他的任務是提醒人們，他們的機器並不像通常所說的那樣聰明，&ldquo;即使有時好像他們會說話，但他們從未真正聆聽過你，它們只是機器。&rdquo;</p>
  <blockquote>
  <p>Eliza的Java實現：<a href="https://github.com/codeanticode/eliza" target="_blank" rel="noopener">詳情點擊</a></p>
  </blockquote>
  <p>譯自 <a href="https://spectrum.ieee.org/tech-talk/artificial-intelligence/machine-learning/why-people-demanded-privacy-to-confide-in-the-worlds-first-chatbot" target="_blank" rel="noopener">Oscar Schwartz</a></p>
  </div>

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自OSCHINA社區[http://www.oschina.net]
  | 本文標題：歷史上第一個聊天機器人：200行代碼就能欺騙人類感情
  | 本文地址：https://www.oschina.net/news/117968/eliza-for-communciation-between-man-and-machine

