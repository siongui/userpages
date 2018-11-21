Lisp 是怎麼成為上帝的編程語言的
###############################

:date: 2018-11-20T17:15+08:00
:tags: Linux, 轉錄, Linux 中國, 開源中國
:category: Linux
:author: Linux 中國翻譯組
:summary: 但有一門語言似乎受到和用途無關的特殊尊敬：那就是 Lisp。即使是恨不得給每個說出形如“某某語言比其他所有語言都好”這類話的人都來一拳的鍵盤遠征軍們，也會承認 Lisp 處於另一個層次。
:og_image: https://dn-linuxcn.qbox.me/data/attachment/album/201811/20/171510jy9911uyu1syvat1.jpg

.. raw:: html

  <div class="d" style="width: 100%;margin: 0;">
  <div id="article_content" style="border-radius:0;">
  
  <p>當程序員們談論各類編程語言的相對優勢時，他們通常會采用相當平淡的措詞，就好像這些語言是一條工具帶上的各種工具似的 &mdash;&mdash; 有適合寫操作系統的，也有適合把其它程序黏在一起來完成特殊工作的。這種討論方式非常合理；不同語言的能力不同。不聲明特定用途就聲稱某門語言比其他語言更優秀只能導致侮辱性的無用爭論。</p>
  <p>但有一門語言似乎受到和用途無關的特殊尊敬：那就是 Lisp。即使是恨不得給每個說出形如&ldquo;某某語言比其他所有語言都好&rdquo;這類話的人都來一拳的鍵盤遠征軍們，也會承認 Lisp 處於另一個層次。 Lisp 超越了用於評判其他語言的實用主義標準，因為普通程序員並不使用 Lisp 編寫實用的程序 &mdash;&mdash; 而且，多半他們永遠也不會這麼做。然而，人們對 Lisp 的敬意是如此深厚，甚至於到了這門語言會時而被加上神話屬性的程度。</p>
  <p>大家都喜歡的網絡漫畫合集 xkcd 就至少在兩組漫畫中如此描繪過 Lisp：<a href="https://xkcd.com/224/">其中一組漫畫</a>中，某人得到了某種 Lisp 啟示，而這好像使他理解了宇宙的基本構架。</p>
  <p><img src="https://dn-linuxcn.qbox.me/data/attachment/album/201811/20/171504x6sgg5ghbxg12rey.jpg" alt="xkcd漫畫" /></p>
  <p>在<a href="https://xkcd.com/297/">另一組漫畫</a>中，一個穿著長袍的老程序員給他的徒弟遞了一沓圓括號，說這是&ldquo;文明時代的優雅武器&rdquo;，暗示著 Lisp 就像原力那樣擁有各式各樣的神秘力量。</p>
  <p><img src="https://dn-linuxcn.qbox.me/data/attachment/album/201811/20/171505edp3pz4eytpxywsc.png" alt="xkcd漫畫" /></p>
  <p>另一個絕佳例子是 Bob Kanefsky 的滑稽劇插曲，《上帝就在人間》。這部劇叫做《永恒之火》，撰寫於 1990 年代中期；劇中描述了上帝必然是使用 Lisp 創造世界的種種原因。完整的歌詞可以在 <a href="https://www.gnu.org/fun/jokes/eternal-flame.en.html">GNU 幽默合集</a>中找到，如下是一段摘抄：</p>
  <blockquote>
  <p>因為上帝用祂的 Lisp 代碼</p>
  <p>讓樹葉充滿綠意。</p>
  <p>分形的花兒和遞歸的根：</p>
  <p>我見過的奇技淫巧之中沒什麼比這更可愛。</p>
  <p>當我對著雪花深思時，</p>
  <p>從未見過兩片相同的，</p>
  <p>我知道，上帝偏愛那一門</p>
  <p>名字是四個字母的語言。</p>
  </blockquote>
  <p>以下這句話我實在不好在人前說；不過，我還是覺得，這樣一種 &ldquo;Lisp 是奧術魔法&rdquo;的文化模因實在是有史以來最奇異、最迷人的東西。Lisp 是象牙塔的產物，是人工智能研究的工具；因此，它對於編程界的俗人而言總是陌生的，甚至是帶有神秘色彩的。然而，當今的程序員們<a href="https://www.reddit.com/r/ProgrammerHumor/comments/5c14o6/xkcd_lisp/d9szjnc/">開始慫恿彼此，&ldquo;在你死掉之前至少試一試 Lisp&rdquo;</a>，就像這是一種令人恍惚入迷的致幻劑似的。盡管 Lisp 是廣泛使用的編程語言中第二古老的（只比 Fortran 年輕一歲）<sup id="fnref1"><a href="#fn1" rel="footnote">1</a></sup> ，程序員們也仍舊在互相慫恿。想象一下，如果你的工作是為某種組織或者團隊推廣一門新的編程語言的話，忽悠大家讓他們相信你的新語言擁有神力難道不是絕佳的策略嗎？&mdash;&mdash; 但你如何能夠做到這一點呢？或者，換句話說，一門編程語言究竟是如何變成人們口中&ldquo;隱晦知識的載體&rdquo;的呢？</p>
  <p>Lisp 究竟是怎麼成為這樣的？</p>
  <p><img src="https://dn-linuxcn.qbox.me/data/attachment/album/201811/20/171510jy9911uyu1syvat1.jpg" alt="Byte 雜誌封面,1979年八月。" /></p>
  <p><em>Byte 雜誌封面，1979年八月。</em></p>
  <h3 id="toc_1">理論 A ：公理般的語言</h3>
  <p>Lisp 的創造者<ruby>約翰&middot;麥卡錫<rt>John McCarthy</rt></ruby>最初並沒有想過把 Lisp 做成優雅、精煉的計算法則結晶。然而，在一兩次運氣使然的深謀遠慮和一系列優化之後，Lisp 的確變成了那樣的東西。 <ruby>保羅&middot;格雷厄姆<rt>Paul Graham</rt></ruby>（我們一會兒之後才會聊到他）曾經這麼寫道， 麥卡錫通過 Lisp &ldquo;為編程作出的貢獻就像是歐幾裏得對幾何學所做的貢獻一般&rdquo; <sup id="fnref2"><a href="#fn2" rel="footnote">2</a></sup>。人們可能會在 Lisp 中看出更加隱晦的含義 &mdash;&mdash; 因為麥卡錫創造 Lisp 時使用的要素實在是過於基礎，基礎到連弄明白他到底是創造了這門語言、還是發現了這門語言，都是一件難事。</p>
  <p>最初， 麥卡錫產生要造一門語言的想法，是在 1956 年的<ruby>達特茅斯人工智能夏季研究項目<rt>Darthmouth Summer Research Project on Artificial Intelligence</rt></ruby>上。夏季研究項目是個持續數周的學術會議，直到現在也仍舊在舉行；它是此類會議之中最早開始舉辦的會議之一。 麥卡錫當初還是個達特茅斯的數學助教，而&ldquo;<ruby>人工智能<rt>artificial intelligence</rt></ruby>（AI）&rdquo;這個詞事實上就是他建議舉辦該會議時發明的 <sup id="fnref3"><a href="#fn3" rel="footnote">3</a></sup>。在整個會議期間大概有十人參加 <sup id="fnref4"><a href="#fn4" rel="footnote">4</a></sup>。他們之中包括了<ruby>艾倫&middot;紐厄爾<rt>Allen Newell</rt></ruby>和<ruby>赫伯特&middot;西蒙<rt>Herbert Simon</rt></ruby>，兩名隸屬於<ruby>蘭德公司<rt>RAND Corporation</rt></ruby>和<ruby>卡內基梅隆大學<rt>Carnegie Mellon</rt></ruby>的學者。這兩人不久之前設計了一門語言，叫做 IPL。</p>
  <p>當時，紐厄爾和西蒙正試圖制作一套能夠在命題演算中生成證明的系統。兩人意識到，用電腦的原生指令集編寫這套系統會非常困難；於是他們決定創造一門語言&mdash;&mdash;他們的原話是&ldquo;<ruby>偽代碼<rt>pseudo-code</rt></ruby>&rdquo;，這樣，他們就能更加輕松自然地表達這臺&ldquo;<ruby>邏輯理論機器<rt>Logic Theory Machine</rt></ruby>&rdquo;的底層邏輯了 <sup id="fnref5"><a href="#fn5" rel="footnote">5</a></sup>。這門語言叫做 IPL，即&ldquo;<ruby>信息處理語言<rt>Information Processing Language</rt></ruby>&rdquo;；比起我們現在認知中的編程語言，它更像是一種高層次的匯編語言方言。 紐厄爾和西蒙提到，當時人們開發的其它&ldquo;偽代碼&rdquo;都抓著標準數學符號不放 &mdash;&mdash; 也許他們指的是 Fortran <sup id="fnref6"><a href="#fn6" rel="footnote">6</a></sup>；與此不同的是，他們的語言使用成組的符號方程來表示命題演算中的語句。通常，用 IPL 寫出來的程序會調用一系列的匯編語言宏，以此在這些符號方程列表中對表達式進行變換和求值。</p>
  <p>麥卡錫認為，一門實用的編程語言應該像 Fortran 那樣使用代數表達式；因此，他並不怎麼喜歡 IPL <sup id="fnref7"><a href="#fn7" rel="footnote">7</a></sup>。然而，他也認為，在給人工智能領域的一些問題建模時，符號列表會是非常好用的工具 &mdash;&mdash; 而且在那些涉及演繹的問題上尤其有用。麥卡錫的渴望最終被訴諸行動；他要創造一門代數的列表處理語言 &mdash;&mdash; 這門語言會像 Fortran 一樣使用代數表達式，但擁有和 IPL 一樣的符號列表處理能力。</p>
  <p>當然，今日的 Lisp 可不像 Fortran。在會議之後的幾年中，麥卡錫關於&ldquo;理想的列表處理語言&rdquo;的見解似乎在逐漸演化。到 1957 年，他的想法發生了改變。他那時候正在用 Fortran 編寫一個能下國際象棋的程序；越是長時間地使用 Fortran ，麥卡錫就越確信其設計中存在不當之處，而最大的問題就是尷尬的 <code>IF</code> 聲明 <sup id="fnref8"><a href="#fn8" rel="footnote">8</a></sup>。為此，他發明了一個替代品，即條件表達式 <code>true</code>；這個表達式會在給定的測試通過時返回子表達式 <code>A</code> ，而在測試未通過時返回子表達式 <code>B</code> ，<em>而且</em>，它只會對返回的子表達式進行求值。在 1958 年夏天，當麥卡錫設計一個能夠求導的程序時，他意識到，他發明的 <code>true</code> 條件表達式讓編寫遞歸函數這件事變得更加簡單自然了 <sup id="fnref9"><a href="#fn9" rel="footnote">9</a></sup>。也是這個求導問題讓麥卡錫創造了 <code>maplist</code> 函數；這個函數會將其它函數作為參數並將之作用於指定列表的所有元素 <sup id="fnref10"><a href="#fn10" rel="footnote">10</a></sup>。在給項數多得叫人抓狂的多項式求導時，它尤其有用。</p>
  <p>然而，以上的所有這些，在 Fortran 中都是沒有的；因此，在 1958 年的秋天，麥卡錫請來了一群學生來實現 Lisp。因為他那時已經成了一名麻省理工助教，所以，這些學生可都是麻省理工的學生。當麥卡錫和學生們最終將他的主意變為能運行的代碼時，這門語言得到了進一步的簡化。這之中最大的改變涉及了 Lisp 的語法本身。最初，麥卡錫在設計語言時，曾經試圖加入所謂的 &ldquo;M 表達式&rdquo;；這是一層語法糖，能讓 Lisp 的語法變得類似於 Fortran。雖然 M 表達式可以被翻譯為 S 表達式 &mdash;&mdash; 基礎的、&ldquo;用圓括號括起來的列表&rdquo;，也就是 Lisp 最著名的特征 &mdash;&mdash; 但 S 表達式事實上是一種給機器看的低階表達方法。唯一的問題是，麥卡錫用方括號標記 M 表達式，但他的團隊在麻省理工使用的 IBM 026 鍵盤打孔機的鍵盤上根本沒有方括號 <sup id="fnref11"><a href="#fn11" rel="footnote">11</a></sup>。於是 Lisp 團隊堅定不移地使用著 S 表達式，不僅用它們表示數據列表，也拿它們來表達函數的應用。麥卡錫和他的學生們還作了另外幾樣改進，包括將數學符號前置；他們也修改了內存模型，這樣 Lisp 實質上就只有一種數據類型了 <sup id="fnref12"><a href="#fn12" rel="footnote">12</a></sup>。</p>
  <p>到 1960 年，麥卡錫發表了他關於 Lisp 的著名論文，《用符號方程表示的遞歸函數及它們的機器計算》。那時候，Lisp 已經被極大地精簡，而這讓麥卡錫意識到，他的作品其實是&ldquo;一套優雅的數學系統&rdquo;，而非普通的編程語言 <sup id="fnref13"><a href="#fn13" rel="footnote">13</a></sup>。他後來這麼寫道，對 Lisp 的許多簡化使其&ldquo;成了一種描述可計算函數的方式，而且它比圖靈機或者一般情況下用於遞歸函數理論的遞歸定義更加簡潔&rdquo; <sup id="fnref14"><a href="#fn14" rel="footnote">14</a></sup>。在他的論文中，他不僅使用 Lisp 作為編程語言，也將它當作一套用於研究遞歸函數行為方式的表達方法。</p>
  <p>通過&ldquo;從一小撮規則中逐步實現出 Lisp&rdquo;的方式，麥卡錫將這門語言介紹給了他的讀者。後來，保羅&middot;格雷厄姆在短文《<ruby><a href="http://languagelog.ldc.upenn.edu/myl/llog/jmc.pdf">Lisp 之根</a><rt>The Roots of Lisp</rt></ruby>》中用更易讀的語言回顧了麥卡錫的步驟。格雷厄姆只用了七種原始運算符、兩種函數寫法，以及使用原始運算符定義的六個稍微高級一點的函數來解釋 Lisp。毫無疑問，Lisp 的這種只需使用極少量的基本規則就能完整說明的特點加深了其神秘色彩。格雷厄姆稱麥卡錫的論文為&ldquo;使計算公理化&rdquo;的一種嘗試 <sup id="fnref15"><a href="#fn15" rel="footnote">15</a></sup>。我認為，在思考 Lisp 的魅力從何而來時，這是一個極好的切入點。其它編程語言都有明顯的人工構造痕跡，表現為 <code>While</code>，<code>typedef</code>，<code>public static void</code> 這樣的關鍵詞；而 Lisp 的設計卻簡直像是純粹計算邏輯的鬼斧神工。Lisp 的這一性質，以及它和晦澀難懂的&ldquo;遞歸函數理論&rdquo;的密切關系，使它具備了獲得如今聲望的充分理由。</p>
  <h3 id="toc_2">理論 B：屬於未來的機器</h3>
  <p>Lisp 誕生二十年後，它成了著名的《<ruby><a href="https://en.wikipedia.org/wiki/Jargon_File">黑客詞典</a><rt>Hacker&rsquo;s Dictionary</rt></ruby>》中所說的，人工智能研究的&ldquo;母語&rdquo;。Lisp 在此之前傳播迅速，多半是托了語法規律的福 &mdash;&mdash; 不管在怎麼樣的電腦上，實現 Lisp 都是一件相對簡單直白的事。而學者們之後堅持使用它乃是因為 Lisp 在處理符號表達式這方面有巨大的優勢；在那個時代，人工智能很大程度上就意味著符號，於是這一點就顯得十分重要。在許多重要的人工智能項目中都能見到 Lisp 的身影。這些項目包括了 <a href="https://hci.stanford.edu/winograd/shrdlu/">SHRDLU 自然語言程序</a>、<a href="https://en.wikipedia.org/wiki/Macsyma">Macsyma 代數系統</a> 和 <a href="https://en.wikipedia.org/wiki/ACL2">ACL2 邏輯系統</a>。</p>
  <p>然而，在 1970 年代中期，人工智能研究者們的電腦算力開始不夠用了。PDP-10 就是一個典型。這個型號在人工智能學界曾經極受歡迎；但面對這些用 Lisp 寫的 AI 程序，它的 18 位地址空間一天比一天顯得吃緊 <sup id="fnref16"><a href="#fn16" rel="footnote">16</a></sup>。許多的 AI 程序在設計上可以與人互動。要讓這些既極度要求硬件性能、又有互動功能的程序在分時系統上優秀發揮，是很有挑戰性的。麻省理工的<ruby>彼得&middot;杜奇<rt>Peter Deutsch</rt></ruby>給出了解決方案：那就是針對 Lisp 程序來特別設計電腦。就像是我那<a href="https://twobithistory.org/2018/09/30/chaosnet.html">關於 Chaosnet 的上一篇文章</a>所說的那樣，這些<ruby>Lisp 計算機<rt>Lisp machines</rt></ruby>會給每個用戶都專門分配一個為 Lisp 特別優化的處理器。到後來，考慮到硬核 Lisp 程序員的需求，這些計算機甚至還配備上了完全由 Lisp 編寫的開發環境。在當時那樣一個小型機時代已至尾聲而微型機的繁盛尚未完全到來的尷尬時期，Lisp 計算機就是編程精英們的&ldquo;高性能個人電腦&rdquo;。</p>
  <p>有那麼一會兒，Lisp 計算機被當成是未來趨勢。好幾家公司雨後春筍般出現，追著趕著要把這項技術商業化。其中最成功的一家叫做 Symbolics，由麻省理工 AI 實驗室的前成員創立。上世紀八十年代，這家公司生產了所謂的 3600 系列計算機，它們當時在 AI 領域和需要高性能計算的產業中應用極廣。3600 系列配備了大屏幕、位圖顯示、鼠標接口，以及<a href="https://youtu.be/gV5obrYaogU?t=201">強大的圖形與動畫軟件</a>。它們都是驚人的機器，能讓驚人的程序運行起來。例如，之前在推特上跟我聊過的機器人研究者 Bob Culley，就能用一臺 1985 年生產的 Symbolics 3650 寫出帶有圖形演示的尋路算法。他向我解釋說，在 1980 年代，位圖顯示和面向對象編程（能夠通過 <a href="https://en.wikipedia.org/wiki/Flavors_(programming_language)">Flavors 擴展</a>在 Lisp 計算機上使用）都剛剛出現。Symbolics 站在時代的最前沿。</p>
  <p><img src="https://dn-linuxcn.qbox.me/data/attachment/album/201811/20/171510lw4x5e6sx9usgwj4.jpg" alt="Bob Culley 的尋路程序。" /></p>
  <p><em>Bob Culley 的尋路程序。</em></p>
  <p>而以上這一切導致 Symbolics 的計算機奇貴無比。在 1983 年，一臺 Symbolics 3600 能賣 111,000 美金 <sup id="fnref16:1"><a href="#fn16" rel="footnote">16</a></sup>。所以，絕大部分人只可能遠遠地贊嘆 Lisp 計算機的威力和操作員們用 Lisp 編寫程序的奇妙技術。不止他們贊嘆，從 1979 年到 1980 年代末，Byte 雜誌曾經多次提到過 Lisp 和 Lisp 計算機。在 1979 年八月發行的、關於 Lisp 的一期特別雜誌中，雜誌編輯激情洋溢地寫道，麻省理工正在開發的計算機配備了&ldquo;大坨大坨的內存&rdquo;和&ldquo;先進的操作系統&rdquo; <sup id="fnref17"><a href="#fn17" rel="footnote">17</a></sup>；他覺得，這些 Lisp 計算機的前途是如此光明，以至於它們的面世會讓 1978 和 1977 年 &mdash;&mdash; 誕生了 Apple II、Commodore PET 和 TRS-80 的兩年 &mdash;&mdash; 顯得黯淡無光。五年之後，在 1985 年，一名 Byte 雜誌撰稿人描述了為&ldquo;復雜精巧、性能強悍的 Symbolics 3670&rdquo;編寫 Lisp 程序的體驗，並力勸讀者學習 Lisp，稱其為&ldquo;絕大數人工智能工作者的語言選擇&rdquo;，和將來的通用編程語言 <sup id="fnref18"><a href="#fn18" rel="footnote">18</a></sup>。</p>
  <p>我問過<ruby>保羅&middot;麥克瓊斯<rt>Paul McJones</rt></ruby>（他在<ruby>山景城<rt>Mountain View</rt><rt></rt></ruby>的<ruby>計算機歷史博物館<rt>Computer History Museum</rt></ruby>做了許多 Lisp 的<a href="http://www.softwarepreservation.org/projects/LISP/">保護工作</a>），人們是什麼時候開始將 Lisp 當作高維生物的贈禮一樣談論的呢？他說，這門語言自有的性質毋庸置疑地促進了這種現象的產生；然而，他也說，Lisp 上世紀六七十年代在人工智能領域得到的廣泛應用，很有可能也起到了作用。當 1980 年代到來、Lisp 計算機進入市場時，象牙塔外的某些人由此接觸到了 Lisp 的能力，於是傳說開始滋生。時至今日，很少有人還記得 Lisp 計算機和 Symbolics 公司；但 Lisp 得以在八十年代一直保持神秘，很大程度上要歸功於它們。</p>
  <h3 id="toc_3">理論 C：學習編程</h3>
  <p>1985 年，兩位麻省理工的教授，<ruby>哈爾&middot;阿伯爾森<rt>Harold <q>Hal</q> Abelson</rt></ruby>和<ruby>傑拉爾德&middot;瑟斯曼<rt>Gerald Sussman</rt></ruby>，外加瑟斯曼的妻子<ruby>朱莉&middot;瑟斯曼<rt>Julie Sussman</rt></ruby>，出版了一本叫做《<ruby>計算機程序的構造和解釋<rt>Structure and Interpretation of Computer Programs</rt></ruby>》的教科書。這本書用 Scheme（一種 Lisp 方言）向讀者們示範了如何編程。它被用於教授麻省理工入門編程課程長達二十年之久。出於直覺，我認為 SICP（這本書的名字通常縮寫為 SICP）倍增了 Lisp 的&ldquo;神秘要素&rdquo;。SICP 使用 Lisp 描繪了深邃得幾乎可以稱之為哲學的編程理念。這些理念非常普適，可以用任意一種編程語言展現；但 SICP 的作者們選擇了 Lisp。結果，這本陰陽怪氣、卓越不凡、吸引了好幾代程序員（還成了一種<a href="https://knowyourmeme.com/forums/meme-research/topics/47038-structure-and-interpretation-of-computer-programs-hugeass-image-dump-for-evidence">奇特的模因</a>）的著作臭名遠揚之後，Lisp 的聲望也順帶被提升了。Lisp 已不僅僅是一如既往的&ldquo;麥卡錫的優雅表達方式&rdquo;；它現在還成了&ldquo;向你傳授編程的不傳之秘的語言&rdquo;。</p>
  <p>SICP 究竟有多奇怪這一點值得好好說；因為我認為，時至今日，這本書的古怪之處和 Lisp 的古怪之處是相輔相成的。書的封面就透著一股古怪。那上面畫著一位朝著桌子走去，準備要施法的巫師或者煉金術士。他的一只手裏抓著一副測徑儀 &mdash;&mdash; 或者圓規，另一只手上拿著個球，上書&ldquo;eval&rdquo;和&ldquo;apply&rdquo;。他對面的女人指著桌子；在背景中，希臘字母 &lambda; （lambda）漂浮在半空，釋放出光芒。</p>
  <p><img src="https://dn-linuxcn.qbox.me/data/attachment/album/201811/20/171512jueceanmca9whatm.jpg" alt="SICP 封面上的畫作" /></p>
  <p><em>SICP 封面上的畫作。</em></p>
  <p>說真的，這上面畫的究竟是怎麼一回事？為什麼桌子會長著動物的腿？為什麼這個女人指著桌子？墨水瓶又是幹什麼用的？我們是不是該說，這位巫師已經破譯了宇宙的隱藏奧秘，而所有這些奧秘就蘊含在 eval/apply 循環和 Lambda 微積分之中？看似就是如此。單單是這張圖片，就一定對人們如今談論 Lisp 的方式產生了難以計量的影響。</p>
  <p>然而，這本書的內容通常並不比封面正常多少。SICP 跟你讀過的所有計算機科學教科書都不同。在引言中，作者們表示，這本書不只教你怎麼用 Lisp 編程 &mdash;&mdash; 它是關於&ldquo;現象的三個焦點：人的心智、復數的計算機程序，和計算機&rdquo;的作品 <sup id="fnref19"><a href="#fn19" rel="footnote">19</a></sup>。在之後，他們對此進行了解釋，描述了他們對如下觀點的堅信：編程不該被當作是一種計算機科學的訓練，而應該是&ldquo;<ruby>程序性認識論<rt>procedural epistemology</rt></ruby>&rdquo;的一種新表達方式 <sup id="fnref20"><a href="#fn20" rel="footnote">20</a></sup>。程序是將那些偶然被送入計算機的思想組織起來的全新方法。這本書的第一章簡明地介紹了 Lisp，但是之後的絕大部分都在講述更加抽象的概念。其中包括了對不同編程範式的討論，對於面向對象系統中&ldquo;時間&rdquo;和&ldquo;一致性&rdquo;的討論；在書中的某一處，還有關於通信的基本限制可能會如何帶來同步問題的討論 &mdash;&mdash; 而這些基本限制在通信中就像是光速不變在相對論中一樣關鍵 <sup id="fnref21"><a href="#fn21" rel="footnote">21</a></sup>。都是些高深難懂的東西。</p>
  <p>以上這些並不是說這是本糟糕的書；這本書其實棒極了。在我讀過的所有作品中，這本書對於重要的編程理念的討論是最為深刻的；那些理念我琢磨了很久，卻一直無力用文字去表達。一本入門編程教科書能如此迅速地開始描述面向對象編程的根本缺陷，和函數式語言&ldquo;將可變狀態降到最少&rdquo;的優點，實在是一件讓人印象深刻的事。而這種描述之後變為了另一種震撼人心的討論：某種（可能類似於今日的 <a href="https://rxjs-dev.firebaseapp.com/">RxJS</a> 的）流範式能如何同時具備兩者的優秀特性。SICP 用和當初麥卡錫的 Lisp 論文相似的方式提純出了高級程序設計的精華。你讀完這本書之後，會立即想要將它推薦給你的程序員朋友們；如果他們找到這本書，看到了封面，但最終沒有閱讀的話，他們就只會記住長著動物腿的桌子上方那神秘的、根本的、給予魔法師特殊能力的、寫著 eval/apply 的東西。話說回來，書上這兩人的鞋子也讓我印象頗深。</p>
  <p>然而，SICP 最重要的影響恐怕是，它將 Lisp 由一門怪語言提升成了必要的教學工具。在 SICP 面世之前，人們互相推薦 Lisp，以學習這門語言為提升編程技巧的途徑。1979 年的 Byte 雜誌 Lisp 特刊印證了這一事實。之前提到的那位編輯不僅就麻省理工的新 Lisp 計算機大書特書，還說，Lisp 這門語言值得一學，因為它&ldquo;代表了分析問題的另一種視角&rdquo; <sup id="fnref22"><a href="#fn22" rel="footnote">22</a></sup>。但 SICP 並未只把 Lisp 作為其它語言的陪襯來使用；SICP 將其作為<em>入門</em>語言。這就暗含了一種論點，那就是，Lisp 是最能把握計算機編程基礎的語言。可以認為，如今的程序員們彼此慫恿&ldquo;在死掉之前至少試試 Lisp&rdquo;的時候，他們很大程度上是因為 SICP 才這麼說的。畢竟，編程語言 <a href="https://en.wikipedia.org/wiki/Brainfuck">Brainfuck</a> 想必同樣也提供了&ldquo;分析問題的另一種視角&rdquo;；但人們學習 Lisp 而非學習 Brainfuck，那是因為他們知道，前者的那種 Lisp 視角在二十年中都被看作是極其有用的，有用到麻省理工在給他們的本科生教其它語言之前，必然會先教 Lisp。</p>
  <h3 id="toc_4">Lisp 的回歸</h3>
  <p>在 SICP 出版的同一年，<ruby>本賈尼&middot;斯特勞斯特盧普<rt>Bjarne Stroustrup</rt></ruby>發布了 C++ 語言的首個版本，它將面向對象編程帶到了大眾面前。幾年之後，Lisp 計算機的市場崩盤，AI 寒冬開始了。在下一個十年的變革中， C++ 和後來的 Java 成了前途無量的語言，而 Lisp 被冷落，無人問津。</p>
  <p>理所當然地，確定人們對 Lisp 重新燃起熱情的具體時間並不可能；但這多半是保羅&middot;格雷厄姆發表他那幾篇聲稱 Lisp 是首選入門語言的短文之後的事了。保羅&middot;格雷厄姆是 Y-Combinator 的聯合創始人和《Hacker News》的創始者，他這幾篇短文有很大的影響力。例如，在短文《<ruby><a href="http://www.paulgraham.com/avg.html">勝於平庸</a><rt>Beating the Averages</rt></ruby>》中，他聲稱 Lisp 宏使 Lisp 比其它語言更強。他說，因為他在自己創辦的公司 Viaweb 中使用 Lisp，他得以比競爭對手更快地推出新功能。至少，<a href="https://web.archive.org/web/20061004035628/http://wiki.alu.org/Chris-Perkins">一部分程序員</a>被說服了。然而，龐大的主流程序員群體並未換用 Lisp。</p>
  <p>實際上出現的情況是，Lisp 並未流行，但越來越多 Lisp 式的特性被加入到廣受歡迎的語言中。Python 有了列表理解。C# 有了 Linq。Ruby&hellip;&hellip;嗯，<a href="http://www.randomhacks.net/2005/12/03/why-ruby-is-an-acceptable-lisp/">Ruby 是 Lisp 的一種</a>。就如格雷厄姆之前在 2001 年提到的那樣，&ldquo;在一系列常用語言中所體現出的&lsquo;默認語言&rsquo;正越發朝著 Lisp 的方向演化&rdquo; <sup id="fnref23"><a href="#fn23" rel="footnote">23</a></sup>。盡管其它語言變得越來越像 Lisp，Lisp 本身仍然保留了其作為&ldquo;很少人了解但是大家都該學的神秘語言&rdquo;的特殊聲望。在 1980 年，Lisp 的誕生二十周年紀念日上，麥卡錫寫道，Lisp 之所以能夠存活這麼久，是因為它具備&ldquo;編程語言領域中的某種近似局部最優&rdquo; <sup id="fnref24"><a href="#fn24" rel="footnote">24</a></sup>。這句話並未充分地表明 Lisp 的真正影響力。Lisp 能夠存活超過半個世紀之久，並非因為程序員們一年年地勉強承認它就是最好的編程工具；事實上，即使絕大多數程序員根本不用它，它還是存活了下來。多虧了它的起源和它的人工智能研究用途，說不定還要多虧 SICP 的遺產，Lisp 一直都那麼讓人著迷。在我們能夠想象上帝用其它新的編程語言創造世界之前，Lisp 都不會走下神壇。</p>
  <hr />
  <p>via: <a href="https://twobithistory.org/2018/10/14/lisp.html">https://twobithistory.org/2018/10/14/lisp.html</a></p>
  <p>作者：<a href="https://twobithistory.org">Two-Bit History</a> 選題：<a href="https://github.com/lujun9972">lujun9972</a> 譯者：<a href="https://github.com/Northurland">Northurland</a> 校對：<a href="https://github.com/wxy">wxy</a></p>
  <p>本文由 <a href="https://github.com/LCTT/TranslateProject">LCTT</a> 原創編譯，<a href="https://linux.cn/">Linux中國</a> 榮譽推出</p>
  <hr />
  <ol>
  <li id="fn1">
  <p>John McCarthy, &ldquo;History of Lisp&rdquo;, 14, Stanford University, February 12, 1979, accessed October 14, 2018, <a href="http://jmc.stanford.edu/articles/lisp/lisp.pdf">http://jmc.stanford.edu/articles/lisp/lisp.pdf</a>&nbsp;<a href="#fnref1" rev="footnote">↩</a></p>
  </li>
  <li id="fn2">
  <p>Paul Graham, &ldquo;The Roots of Lisp&rdquo;, 1, January 18, 2002, accessed October 14, 2018, <a href="http://languagelog.ldc.upenn.edu/myl/llog/jmc.pdf">http://languagelog.ldc.upenn.edu/myl/llog/jmc.pdf</a>. &nbsp;<a href="#fnref2" rev="footnote">↩</a></p>
  </li>
  <li id="fn3">
  <p>Martin Childs, &ldquo;John McCarthy: Computer scientist known as the father of AI&rdquo;, The Independent, November 1, 2011, accessed on October 14, 2018, <a href="https://www.independent.co.uk/news/obituaries/john-mccarthy-computer-scientist-known-as-the-father-of-ai-6255307.html">https://www.independent.co.uk/news/obituaries/john-mccarthy-computer-scientist-known-as-the-father-of-ai-6255307.html</a>. &nbsp;<a href="#fnref3" rev="footnote">↩</a></p>
  </li>
  <li id="fn4">
  <p>Lisp Bulletin History. <a href="http://www.artinfo-musinfo.org/scans/lb/lb3f.pdf">http://www.artinfo-musinfo.org/scans/lb/lb3f.pdf</a> &nbsp;<a href="#fnref4" rev="footnote">↩</a></p>
  </li>
  <li id="fn5">
  <p>Allen Newell and Herbert Simon, &ldquo;Current Developments in Complex Information Processing,&rdquo; 19, May 1, 1956, accessed on October 14, 2018, <a href="http://bitsavers.org/pdf/rand/ipl/P-850_Current_Developments_In_Complex_Information_Processing_May56.pdf">http://bitsavers.org/pdf/rand/ipl/P-850_Current_Developments_In_Complex_Information_Processing_May56.pdf</a>. &nbsp;<a href="#fnref5" rev="footnote">↩</a></p>
  </li>
  <li id="fn6">
  <p>ibid. &nbsp;<a href="#fnref6" rev="footnote">↩</a></p>
  </li>
  <li id="fn7">
  <p>Herbert Stoyan, &ldquo;Lisp History&rdquo;, 43, Lisp Bulletin #3, December 1979, accessed on October 14, 2018, <a href="http://www.artinfo-musinfo.org/scans/lb/lb3f.pdf">http://www.artinfo-musinfo.org/scans/lb/lb3f.pdf</a> &nbsp;<a href="#fnref7" rev="footnote">↩</a></p>
  </li>
  <li id="fn8">
  <p>McCarthy, &ldquo;History of Lisp&rdquo;, 5. &nbsp;<a href="#fnref8" rev="footnote">↩</a></p>
  </li>
  <li id="fn9">
  <p>ibid. &nbsp;<a href="#fnref9" rev="footnote">↩</a></p>
  </li>
  <li id="fn10">
  <p>McCarthy &ldquo;History of Lisp&rdquo;, 6. &nbsp;<a href="#fnref10" rev="footnote">↩</a></p>
  </li>
  <li id="fn11">
  <p>Stoyan, &ldquo;Lisp History&rdquo;, 45 &nbsp;<a href="#fnref11" rev="footnote">↩</a></p>
  </li>
  <li id="fn12">
  <p>McCarthy, &ldquo;History of Lisp&rdquo;, 8. &nbsp;<a href="#fnref12" rev="footnote">↩</a></p>
  </li>
  <li id="fn13">
  <p>McCarthy, &ldquo;History of Lisp&rdquo;, 2. &nbsp;<a href="#fnref13" rev="footnote">↩</a></p>
  </li>
  <li id="fn14">
  <p>McCarthy, &ldquo;History of Lisp&rdquo;, 8. &nbsp;<a href="#fnref14" rev="footnote">↩</a></p>
  </li>
  <li id="fn15">
  <p>Graham, &ldquo;The Roots of Lisp&rdquo;, 11. &nbsp;<a href="#fnref15" rev="footnote">↩</a></p>
  </li>
  <li id="fn16">
  <p>Guy Steele and Richard Gabriel, &ldquo;The Evolution of Lisp&rdquo;, 22, History of Programming Languages 2, 1993, accessed on October 14, 2018, <a href="http://www.dreamsongs.com/Files/HOPL2-Uncut.pdf">http://www.dreamsongs.com/Files/HOPL2-Uncut.pdf</a>. &nbsp;<a href="#fnref16" rev="footnote">↩</a>&nbsp;<a href="#fnref16:1" rev="footnote">↩<sup>1</sup></a></p>
  </li>
  <li id="fn17">
  <p>Carl Helmers, &ldquo;Editorial&rdquo;, Byte Magazine, 154, August 1979, accessed on October 14, 2018, <a href="https://archive.org/details/byte-magazine-1979-08/page/n153">https://archive.org/details/byte-magazine-1979-08/page/n153</a>. &nbsp;<a href="#fnref17" rev="footnote">↩</a></p>
  </li>
  <li id="fn18">
  <p>Patrick Winston, &ldquo;The Lisp Revolution&rdquo;, 209, April 1985, accessed on October 14, 2018, <a href="https://archive.org/details/byte-magazine-1985-04/page/n207">https://archive.org/details/byte-magazine-1985-04/page/n207</a>. &nbsp;<a href="#fnref18" rev="footnote">↩</a></p>
  </li>
  <li id="fn19">
  <p>Harold Abelson, Gerald Jay. Sussman, and Julie Sussman, Structure and Interpretation of Computer Programs (Cambridge, Mass: MIT Press, 2010), xiii. &nbsp;<a href="#fnref19" rev="footnote">↩</a></p>
  </li>
  <li id="fn20">
  <p>Abelson, xxiii. &nbsp;<a href="#fnref20" rev="footnote">↩</a></p>
  </li>
  <li id="fn21">
  <p>Abelson, 428. &nbsp;<a href="#fnref21" rev="footnote">↩</a></p>
  </li>
  <li id="fn22">
  <p>Helmers, 7. &nbsp;<a href="#fnref22" rev="footnote">↩</a></p>
  </li>
  <li id="fn23">
  <p>Paul Graham, &ldquo;What Made Lisp Different&rdquo;, December 2001, accessed on October 14, 2018, <a href="http://www.paulgraham.com/diff.html">http://www.paulgraham.com/diff.html</a>. &nbsp;<a href="#fnref23" rev="footnote">↩</a></p>
  </li>
  <li id="fn24">
  <p>John McCarthy, &ldquo;Lisp&mdash;Notes on its past and future&rdquo;, 3, Stanford University, 1980, accessed on October 14, 2018, <a href="http://jmc.stanford.edu/articles/lisp20th/lisp20th.pdf">http://jmc.stanford.edu/articles/lisp20th/lisp20th.pdf</a>. &nbsp;<a href="#fnref24" rev="footnote">↩</a></p>
  </li>
  </ol><div style="position: relative;bottom: 180px;right: 10px;float: right;height: 0;"><div id="translator_info" class="bm">
  <div class="bm_h cl">
  LCTT 譯者
  </div>
  <div class="bm_c">
  <div >
  <a href="/lctt/Northurland" target="_blank"><img class="avatar" src="https://avatars1.githubusercontent.com/u/40388212?v=4" style="max-width: 64px; float: left;margin-top: 10px;" /></a>
  <div style="float: left;margin-left: 10px;font-size: 12px;line-height: 1.5em; ">
  <div>
  <a href="https://github.com/Northurland" target="_blank"><img src="https://dn-linuxcn.qbox.me/static/image/common/github_icon.png" style="vertical-align:middle;" /></a> <a href="/lctt/Northurland" target="_blank" style="font-weight: bold;">Northurland</a> 🌟</div>
  <div class="addfiles">共計翻譯： <span class="num" style="color: #f00;font-weight: 700;">1.0</span> 篇
  | 共計貢獻： <span class="num" style="color: #f00;font-weight: 700;">16</span> 天</div>
  <div>貢獻時間：2018-10-26 -&gt; 2018-11-10</div><a href="/lctt/Northurland" target="_blank">訪問我的 LCTT 主頁</a> | <a href="https://github.com/Northurland" target="_blank">在 GitHub 上關注我</a>
  </div>
  <br class="clear" />
  </div>
  </div>
  </div>
  </div>
  </div>
  </div>

  <div class="copyright">
      	 
      		    			    			<span class="z textcut">編譯自：<a href="https://twobithistory.org/2018/10/14/lisp.html" target="_blank">https://twobithistory.org/2018/10/14/lisp.html</a></span>
      			    		    		<span class="y">作者： Two-bit History</span>    		<br class="clear" />
      	    	<span class="z">原創：<a href="https://linux.cn/lctt/" target=_blank>LCTT</a> <a href="https://linux.cn/article-10255-1.html" target="_blank">https://linux.cn/article-10255-1.html</a></span>
              			<span class="y">譯者： Northurland</span>    		<br class="clear" />
      	    	<br />
      	    		<span>本文由 LCTT 原創翻譯，<a href="https://linux.cn/article-10255-1.html">Linux中國首發</a>。也想加入譯者行列，為開源做一些自己的貢獻麼？歡迎加入 <a href="http://lctt.github.io/" target=_blank>LCTT</a>！<br />翻譯工作和譯文發表僅用於學習和交流目的，翻譯工作遵照 <a href="http://creativecommons.org/licenses/by-nc-sa/3.0/deed.zh" target=_blank>CC-BY-NC-SA 協議規定</a>，如果我們的工作有侵犯到您的權益，請及時聯繫我們。</span>
      		    		<br class="clear" />
      		<span style="color:red;">歡迎遵照 <a href="http://creativecommons.org/licenses/by-nc-sa/3.0/deed.zh" target=_blank>CC-BY-NC-SA 協議規定</a>轉載，敬請在正文中標註並保留原文/譯文鏈接和作者/譯者等信息。</span>
                  		<br class="clear" />
  </div>

- `Lisp 是怎么成为上帝的编程语言的 - 开源中国 <https://www.oschina.net/news/101976/lisp-became-gods-programming-language>`_
