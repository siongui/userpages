Go基礎學習記錄之Session和Cookie
###############################

:date: 2018-09-11T20:33:25+08:00
:tags: Go語言, 小絨毛的足跡, Web開發, 轉錄
:category: Go語言
:author: durban
:summary: session和cookie是兩個非常常見的Web概念，也很容易被誤解。但是，它們對於頁面授權以及收集頁面統計信息非常重要。我們來看看這兩個用例。
:og_image: https://res.cloudinary.com/dy5dvcuc1/image/upload/v1536669103/golang_web_25_1.png


.. raw:: html

  <h2>Session和Cookie</h2><p>session和cookie是兩個非常常見的Web概念，也很容易被誤解。<br/>
  但是，它們對於頁面授權以及收集頁面統計信息非常重要。<br/>
  我們來看看這兩個用例。</p><p>假設我們要抓取限制公共訪問的頁面，例如Twitter用戶的主頁。<br/>
  當然，您可以打開瀏覽器並輸入用戶名和密碼來登錄和訪問該信息，但所謂的&ldquo;網絡爬行&rdquo;意味著我們使用程序自動執行此過程而無需任何人為干預。<br/>
  因此，當我們使用瀏覽器登錄時，我們必須找出幕後的真實情況。</p><p>當我們第一次收到登錄頁面並輸入用戶名和密碼時，按下&ldquo;登錄&rdquo;按鈕後，瀏覽器會向遠程服務器發送POST請求。<br/>
  服務器驗證登錄信息並返回HTTP響應後，瀏覽器重定向到用戶主頁。<br/>
  這裏的問題是，服務器如何知道我們擁有所需網頁的訪問權限？<br/>
  由於HTTP是無狀態的，因此服務器無法知道我們是否在最後一步中通過了驗證。<br/>
  最簡單也許最天真的解決方案是將用戶名和密碼附加到URL。<br/>
  這有效，但對服務器施加了太大的壓力（服務器必須驗證針對數據庫的每個請求），並且可能對用戶體驗有害。<br/>
  實現此目標的另一種方法是使用cookie和session在服務器端或客戶端保存用戶的身份。<br/>
  簡而言之，Cookie會在客戶端的計算機上存儲歷史信息（包括用戶登錄信息）。<br/>
  每次用戶訪問同一網站時，客戶端的瀏覽器都會發送這些cookie，自動完成用戶的登錄步驟。</p><p>另一方面，session在服務器端存儲歷史信息。<br/>
  服務器使用session ID來標識不同的session，並且服務器生成的session ID應始終是隨機且唯一的。<br/>
  您可以使用cookie或URL參數來獲取客戶端的身份。</p><p>Cookie由瀏覽器維護。<br/>
  可以在Web服務器和瀏覽器之間進行通信時修改它們。<br/>
  當用戶訪問相應的網站時，Web應用程序可以訪問cookie信息。<br/>
  在大多數瀏覽器設置中，有一個與cookie隱私相關的設置。<br/>
  打開它時，您應該能夠看到類似於以下內容的內容。</p><p><img alt="Go基礎學習記錄之Session和Cookie" src="https://res.cloudinary.com/dy5dvcuc1/image/upload/v1536669103/golang_web_25_1.png" style="height:310px;width:300px"/></p><p>Cookie具有到期時間，並且有兩種類型的Cookie以其生命周期區分：會話cookie和持久性cookie。<br/>
  如果您的應用程序未設置cookie到期時間，瀏覽器關閉後瀏覽器將不會將其保存到本地文件系統中。<br/>
  這些cookie稱為會話cookie，這種類型的cookie通常保存在內存中而不是本地文件系統中。<br/>
  如果您的應用程序確實設置了到期時間(例如，setMaxAge(606024))，瀏覽器會將此cookie保存到本地文件系統，並且在達到分配的到期時間之前不會刪除它。<br/>
  保存到本地文件系統的Cookie可以由不同的瀏覽器進程共享 - 例如，通過兩個IE窗口;<br/>
  不同的瀏覽器使用不同的進程來處理保存在內存中的cookie。</p><h2>Cookie</h2><h3>Set Cookie</h3><p>Go使用net/http包中的SetCookie函數來設置cookie：</p>

.. code-block:: go

  http.SetCookie(w ResponseWriter, cookie *Cookie)

w是請求的響應，cookie是結構。讓我們看看Cookie的結構：

.. code-block:: go

  // A Cookie represents an HTTP cookie as sent in the Set-Cookie header of an
  // HTTP response or the Cookie header of an HTTP request.
  //
  // See https://tools.ietf.org/html/rfc6265 for details.
  type Cookie struct {
      Name  string
      Value string

      Path       string    // optional
      Domain     string    // optional
      Expires    time.Time // optional
      RawExpires string    // for reading cookies only    // MaxAge=0 means no 'Max-Age' attribute specified.
      // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
      // MaxAge>0 means Max-Age attribute present and given in seconds
      MaxAge   int
      Secure   bool
      HttpOnly bool
      Raw      string
      Unparsed []string // Raw text of unparsed attribute-value pairs
  }

下面讓我們實例演示一下

.. code-block:: go

  expired := time.Now().Add(365 * 24 * time.Hour)
  cookie := http.Cookie{
      Name:    "site_name",
      Value:   "gowhich",
      Expires: expired,
  }

  http.SetCookie(w, &cookie)

上面的示例顯示了如何設置cookie。現在讓我們看看如何獲取已設置的cookie：

.. raw:: html

  <h3>Get Cookie</h3>

.. code-block:: go

  var siteName string
  cookie, err := r.Cookie("site_name")
  if err != nil {
      siteName = cookie.Value
  }
  fmt.Println(siteName)

.. raw:: html

  <p>從請求中獲取cookie非常方便。</p><h2>Session</h2><p>session是一系列操作或消息。例如，您可以考慮在接聽電話和掛起電話之間採取的動作。<br/>
  在網絡協議方面，session更多地與瀏覽器和服務器之間的連接有關。<br/>
  session有助於存儲服務器和客戶端之間的連接狀態，這有時可以採取數據存儲結構的形式。<br/>
  session是服務器端機制，通常使用哈希表（或類似的東西）來保存傳入的信息。<br/>
  當應用程序需要為客戶端分配新會話時，服務器應檢查是否存在具有唯一session ID的同一客戶端的任何現有session。<br/>
  如果session ID已存在，則服務器將只返回同一session到客戶端。<br/>
  另一方面，如果客戶端不存在session ID，則服務器會創建一個全新的session（這通常發生在服務器刪除了相應的sessionID，但用戶已手動附加舊session時）。<br/>
  session本身並不複雜，但它的實現和部署是比較複雜的，所以你不能使用&ldquo;一種方式來統治它們&rdquo;。</p><p><br/>
  總之，session和cookie的目的是相同的。<br/>
  它們都是為了克服HTTP的無狀態，但它們使用不同的方法。<br/>
  session使用cookie在客戶端保存session ID，並在服務器端保存所有其他信息。<br/>
  Cookie會在客戶端保存所有客戶端信息。<br/>
  您可能已經注意到cookie存在一些安全問題。<br/>
  例如，惡意第三方網站可能會破解和收集用戶名和密碼。</p><p>以下是兩個常見的漏洞：</p><ul><li>appA為appB設置了一個意外的cookie。</li><li>XSS攻擊：appA使用JavaScript document.cookie訪問appB的cookie。</li></ul><p>看完本次分享後，您應該了解一些cookie和session的基本概念。你應該能夠理解它們之間的差異，這樣當不可避免地出現bug時你就不會自殺。<br/>
  &nbsp;</p>

.. raw:: html

  <div class="entry-copyright"><h5>版權聲明</h5><p>由<span><a href="https://www.xiaorongmao.com/author/durban" title="durban"> durban</a></span>創作並維護的<span><a href="https://www.xiaorongmao.com" title="小絨毛的足跡"> 小絨毛的足跡</a></span>博客採用<span><a href='http://creativecommons.org/licenses/by-nc-nd/4.0/' title='創作共用保留署名-非商業-禁止演繹4.0國際許可證。'><span>創作共用保留署名-非商業-禁止演繹4.0國際許可證。</span></a></span></p><p>本文首發於
        <a href="https://www.xiaorongmao.com" title="小絨毛的足跡"><span itemprop="publisher" itemscope="itemscope" itemtype="http://schema.org/Organization"><span itemprop="name">小絨毛的足跡</span><span itemprop="logo" itemscope="itemscope" itemtype="http://schema.org/ImageObject" style="display:none"><img itemprop="url" src="https://www.xiaorongmao.com/images/logo.png" alt="小絨毛的足跡"/></span></span></a>博客（
        <a href="https://www.xiaorongmao.com" title="小絨毛的足跡"><span>https://www.xiaorongmao.com</span></a> ），版權所有，侵權必究。</p><p><span>本文永久鏈接：</span><span itemprop="mainEntityOfPage" itemscope="itemscope" itemtype="http://schema.org/WebPage" id="80"><a href="https://www.xiaorongmao.com/blog/80"> https://www.xiaorongmao.com/blog/80</a></span><span style="display:none"/></p></div>

----

- `Go基础学习记录之Session和Cookie - 博文 - 小绒毛的足迹 <https://www.xiaorongmao.com/blog/80>`_
- `Go基础学习记录之Session和Cookie - Gowhich - SegmentFault 思否 <https://segmentfault.com/a/1190000016788585>`_
- `Go基础学习记录之Session和Cookie  - Go语言中文网 - Golang中文社区 <https://studygolang.com/articles/15753>`_
- `Go基礎學習記錄之Session和Cookie - 掃文資訊 <https://tw.saowen.com/a/cf342068a4963b5cc942bc1db2440c3e78ba94b1b8975e29055df035a5104724>`_
