微軟 2018 開源大事記
####################

:date: 2018-12-26T09:09+08:00
:tags: 轉錄, 開源中國
:category: 轉錄
:author: 局长(開源中國)
:summary: 從微軟公開宣佈 "Microsoft love linux" 那一刻起，過去的幾年裡，微軟積極擁抱開源的舉動我們有目共睹，即便有過"Linux is a cancer"這種真香警告的 flag，但不得不承認的是，微軟一系列“擁抱開源”的舉措為 Linux 社區乃至整個開源社區都帶來了很多巨大的驚喜。
:og_image: https://static.oschina.net/uploads/space/2018/0604/214736_WD4N_2720166.jpg

從微軟公開宣佈 "Microsoft love linux" 那一刻起，過去的幾年裡，微軟積極擁抱開源的舉動我們有目共睹，即便有過"Linux is a cancer"這種真香警告的 flag，但不得不承認的是，微軟一系列“擁抱開源”的舉措為 Linux 社區乃至整個開源社區都帶來了很多巨大的驚喜。

2015年，微軟宣佈 `支持開源視頻編解碼器 VP9`_ ，對自家的更多項目進行開源，並開始接受 LLVM/Clang；而在2016年 `微軟收購了跨平台移動開發公司 Xamarin`_ ， `推出了 Linux 版的 SQL Server`_ ，並表示會繼續開源。2017年也是非常有趣的一年，這一年，微軟加入了 OSI 組織（開源計畫，Open Source Initiative），繼續積極發展 Windows 的 Linux 子系統(WSL)，並在 Linux 平台上為 .NET 開發提供更多的便利特性。但今年，2018年，可以說是微軟最令人感到驚訝的一年。 

.. _支持開源視頻編解碼器 VP9: https://www.oschina.net/news/66015/announcing-vp9-support-coming-to-microsoft-edge
.. _微軟收購了跨平台移動開發公司 Xamarin: https://www.oschina.net/news/71002/welcoming-the-xamarin-team-to-microsoft
.. _推出了 Linux 版的 SQL Server: https://www.oschina.net/news/79145/ms-release-preview-linux-version-of-sql-server

下面不妨回顧一下 2018 年微軟與開源有關的“里程碑意義”事件：


1. `微軟收購 GitHub`_
+++++++++++++++++++++

.. _微軟收購 GitHub: https://www.oschina.net/news/96750/microsoft-to-acquire-github-for-7-5-billion

2018年6月4日，微軟宣佈以 75 億美元的股票收購代碼託管平台 GitHub，並由微軟副總裁、Xamarin 創始人 Nat Friedman 擔任 CEO 一職。消息公佈後，微軟表示 GitHub 將仍然是一個開放平台，並保留其開發者優先的風格，獨立運營。而這一事件，無論是對於微軟，還是整個開源界，都是一件可以載入史冊的大事。

.. image:: https://static.oschina.net/uploads/space/2018/0604/214736_WD4N_2720166.jpg
   :alt: 微軟 2018 開源大事記
   :align: center


2. `微軟採用 Chromium 內核開發桌面版 Edge 瀏覽器`_
++++++++++++++++++++++++++++++++++++++++++++++++++

.. _微軟採用 Chromium 內核開發桌面版 Edge 瀏覽器: https://www.oschina.net/news/102458/goodbye-edgehtml

2018年12月，微軟通過其博客 `官方宣佈`_ ：未來將採用 Chromium 內核開發桌面版 Edge 瀏覽器，以便為用戶帶來更好的 Web 兼容性，並為所有 Web 開發者減少 Web 碎片化。

.. _官方宣佈: https://blogs.windows.com/windowsexperience/2018/12/06/microsoft-edge-making-the-web-better-through-more-open-source-collaboration/

.. image:: https://static.oschina.net/uploads/space/2018/1207/061238_YW06_2720166.jpg
   :alt: 微軟 2018 開源大事記
   :align: center

微軟還計畫將 Edge 瀏覽器引入 macOS 平台。此外，新版 Edge 可運行在所有受支持的 Windows 版本上，包括 Windows 7 和 Windows 8 系列 —— 直到微軟終止支持，而不僅僅是 Windows 10。


3. `微軟加入 OIN 社區，並開放 6 萬多項專利`_
++++++++++++++++++++++++++++++++++++++++++++

.. _微軟加入 OIN 社區，並開放 6 萬多項專利: https://www.oschina.net/news/100738/microsoft-joins-open-invention-network

2018年10月11日，微軟宣佈正式 `加入 Open Invention Network (“OIN”) 社區`_ 。可以把 Open Invention Network 理解為是一個共享的防禦性專利池，它成立於2015年，旨在保護 Linux 及其相關的開源項目。而為了提升 Linux 和其他開源技術的優勢，微軟給 OIN 帶去了超過6萬項已發佈的專利。

在同一個月， `微軟宣佈加入 LOT Network`_ 。LOT Network 是一個不斷發展的非營利性社區，由谷歌、Dropbox 等數家科技公司聯合創辦，是一個專門用於對抗專利流氓的組織。微軟的加入意味著它正在就這一主題與其他行業領導者保持一致，並承諾在未來採取更多措施來應對知識產權風險。通過加入 LOT Network，微軟承諾會將自家的專利免費許可給其他成員使用，而 LOT Network 大約覆蓋了 135 萬項專利。

.. _加入 Open Invention Network (“OIN”) 社區: https://azure.microsoft.com/en-us/blog/microsoft-joins-open-invention-network-to-help-protect-linux-and-open-source/
.. _微軟宣佈加入 LOT Network: https://azure.microsoft.com/en-us/blog/microsoft-joins-lot-network-helping-protect-developers-against-patent-assertions/


4. `微軟開源三個主要的 Windows UX 技術：WPF, Windows Forms 和 WinUI 框架`_
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

.. _微軟開源三個主要的 Windows UX 技術：WPF, Windows Forms 和 WinUI 框架: https://www.oschina.net/news/102393/ms-connect-2018

在微軟舉辦的 `Microsoft Connect(); 2018 開發者大會`_ 上，微軟為 .NET 開發者帶來了重磅的信息。在大會上，微軟宣佈開源三種主要的 Windows UX 技術，分別是 Windows Presentation Foundation (WPF), Windows Forms 和 Windows UI XAML 庫 (WinUI)。

WPF 是用於構建桌面客戶端應用程序的 UI 框架，具有應用程序模型、控件、圖形、佈局、數據綁定和安全性等功能。WPF 的核心是一個利用現代圖形硬件的渲染引擎。

Windows Forms 用於支持開發“智能客戶端(smart clients)”，微軟對其的描述為易於部署和更新的應用程序。無論這些應用程序是否連接到 Internet，它們都可以正常工作。Windows Forms 中的表單是用於向用戶顯示信息的可視界面。WinUI 具有 Windows 10 默認 UWP XAML UI 平台的向後兼容實現。

.. _Microsoft Connect(); 2018 開發者大會: https://aka.ms/connectevent


5. `微軟重新開源已誕生 36 年的 MS-DOS 1.25/2.0`_
++++++++++++++++++++++++++++++++++++++++++++++++

.. _微軟重新開源已誕生 36 年的 MS-DOS 1.25/2.0: https://www.oschina.net/news/100460/microsoft-open-source-ms-dos

2018年9月，微軟宣佈在 GitHub_ 上重新開源 MS-DOS 1.25、2.0，相比原來的可下載壓縮文件更容易查找、閱讀和引用。MS-DOS 1.25、2.0 的所有源代碼都是用 8086 彙編代碼寫的，其中 86-DOS 的代碼最初完成於1980年12月29日。

.. _GitHub: https://github.com/microsoft/ms-dos

.. image:: https://oscimg.oschina.net/oscnet/8c534e4c74deaadf0190156745e4e0173ce.jpg
   :alt: 微軟 2018 開源大事記
   :align: center

而微軟開源放出的除了源文件和目標文件，還有一些有趣的 .txt、.doc 文件，十分值得一讀，其中包括很多代碼註釋。


6. `微軟的分佈式系統平台 Service Fabric 正式開源`_
++++++++++++++++++++++++++++++++++++++++++++++++++

.. _微軟的分佈式系統平台 Service Fabric 正式開源: https://www.oschina.net/news/94237/service-fabric-opensource

2018年3月14日，微軟宣佈正式開源 Service Fabric，採用 MIT 開源許可證。Windows 內部為 Service Fabric 開發了將近十年的內部服務，其中大部分時間都是微軟內部平台，比如 Office365，Azure Stack 平台等。

其中包括 `Reliable Services`_ ， `Reliable Actors`_ 和 `ASP.NET Core 集成庫`_ ，Azure 基礎架構服務以及 Azure SQL DB，Azure Cosmos DB 和 Cortana 等大型解決方案都在使用 Service Fabric 構建。

.. _Reliable Services: https://github.com/Azure/service-fabric-services-and-actors-dotnet
.. _Reliable Actors: https://github.com/Azure/service-fabric-services-and-actors-dotnet
.. _ASP.NET Core 集成庫: https://github.com/Azure/service-fabric-aspnetcore


7. `微軟正式開源 Blazor ，將 .NET 帶回到瀏覽器`_
++++++++++++++++++++++++++++++++++++++++++++++++

.. _微軟正式開源 Blazor ，將 .NET 帶回到瀏覽器: https://www.oschina.net/news/93475/microsoft-opensource-blazor

2018年2月，微軟  ASP.Net 團隊正式開源 Blazor，這是一個 Web UI 框架，可通過 WebAssembly 在任意瀏覽器中運行 .NET。

Blazor 旨在簡化快速的單頁面 .NET 瀏覽器應用的構建過程，它雖然使用了諸如 CSS 和 HTML 之類的 Web 技術，但它使用 C＃語言和 Razor 語法代替 JavaScript 來構建可組合的 Web UI 。通過提供用於編譯到 Web 的大小和高效加載的格式，WebAssembly 可讓 .NET 在瀏覽器中運行。


8. `微軟代碼託管平台 CodePlex 正式關閉，進入封存狀態`_
++++++++++++++++++++++++++++++++++++++++++++++++++++++

.. _微軟代碼託管平台 CodePlex 正式關閉，進入封存狀態: https://www.oschina.net/news/93043/codeplex-has-been-archived

隨著 CodePlex 的不斷沒落，微軟於2018年1月 30 日發文 `宣佈`_ ，CodePlex.com 網站正式退役，現在打開該網址會跳轉到一個 `封存歸檔頁面`_ 。該存檔包含在2017年下半年進入只讀模式前託管到 CodePlex 上的所有項目。

.. _宣佈: https://blogs.msdn.microsoft.com/codeplex/2018/01/30/codeplex-has-been-archived/
.. _封存歸檔頁面: https://archive.codeplex.com/

.. image:: https://static.oschina.net/uploads/space/2018/0202/164141_iYI2_2896879.png
   :alt: 微軟 2018 開源大事記
   :align: center

CodePlex 是微軟於2006年推出的一個開源軟件託管平台，在提供服務11年之後，微軟於2017年4月1日宣佈將關閉該平台，並給出了結束時間：2017年10月切換到只讀模式，12月徹底關閉（最終的關閉日期是2018年1月29日）。之後，CodePlex 將為之前的開源項目存檔，人們可以瀏覽並下載這些項目。


9. `微軟力挺 Go，宣佈參與 Athens 項目和 GopherSource`_
++++++++++++++++++++++++++++++++++++++++++++++++++++++

.. _微軟力挺 Go，宣佈參與 Athens 項目和 GopherSource: https://www.oschina.net/news/99544/ms-announces-project-athens-gophersource

在2018年的 Go 開發者大會 GopherCon 上，微軟宣佈要為 Athens 項目貢獻代碼，以及與 GopherSource 進行合作。

Athens 是一個開源項目，旨在為 Go 模塊(Go modules)創建首個代理服務器。微軟表示該項目目前仍處於 alpha 階段，並將聯合 Athens 社區繼續致力於改善模塊體驗，重點是確保 Go 模塊與所有代理服務器能無縫協作，並努力建立一個聯合的、組織多樣化的代理網絡。

而所謂的 GopherSource 其實是一項新計畫，旨在通過在社區內為上游項目和關鍵 Go 項目（如 Athens）帶來更多用戶和貢獻者，以增強和實現 Go 生態系統的多樣化。

對於微軟 2018 開源大事記的回顧到此為止，如有紕漏，歡迎在評論區指正。

..
  .. image:: 
   :alt: 
   :align: center

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：微軟 2018 開源大事記
  | 本文地址：https://www.oschina.net/news/103045/microsoft-2018-open-source-events-recap

