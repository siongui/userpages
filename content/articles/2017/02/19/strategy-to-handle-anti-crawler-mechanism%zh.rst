普通反爬蟲機制的應對策略
########################

:date: 2017-02-19T18:50
:author: Melwood
:tags: Python, Web開發, 轉錄
:category: Python
:summary: 這篇文章主要討論使用Scrapy框架時，如何應對普通的反爬機制。
:og_image: http://www.treselle.com/wp-content/uploads/freshizer/1baed853c082fe88880a8215d4cd0bb2_phython-with-Scarpy-863-430-c.jpg

轉載自： `普通反爬虫机制的应对策略 - Melwood <http://jiayi.space/post/fan-pa-chong-de-ying-dui-ce-lue>`_

爬蟲與反爬蟲，這相愛相殺的一對，簡直可以寫出一部壯觀的鬥爭史。而在大數據時代，數據就是金錢，很多企業都為自己的網站運用了反爬蟲機制，防止網頁上的數據被爬蟲爬走。然而，如果反爬機制過於嚴格，可能會誤傷到真正的用戶請求；如果既要和爬蟲死磕，又要保證很低的誤傷率，那麼又會加大研發的成本。

簡單低級的爬蟲速度快，偽裝度低，如果沒有反爬機制，它們可以很快的抓取大量數據，甚至因為請求過多，造成服務器不能正常工作。而偽裝度高的爬蟲爬取速度慢，對服務器造成的負擔也相對較小。所以，網站反爬的重點也是那種簡單粗暴的爬蟲，反爬機制也會允許偽裝度高的爬蟲，獲得數據。畢竟偽裝度很高的爬蟲與真實用戶也就沒有太大差別了。

這篇文章主要討論使用Scrapy框架時，如何應對普通的反爬機制。


header檢驗
++++++++++

最簡單的反爬機制，就是檢查HTTP請求的Headers信息，包括User-Agent, Referer、Cookies等。

User-Agent
==========

User-Agent是檢查用戶所用客戶端的種類和版本，在Scrapy中，通常是在下載器中間件中進行處理。比如在setting.py中建立一個包含很多瀏覽器User-Agent的列表，然後新建一個random_user_agent文件：

.. code-block:: python

  class RandomUserAgentMiddleware(object):
      @classmethod
      def process_request(cls, request, spider):
          ua = random.choice(spider.settings['USER_AGENT_LIST'])
          if ua:
              request.headers.setdefault('User-Agent', ua)

這樣就可以在每次請求中，隨機選取一個真實瀏覽器的User-Agent。

Referer
=======

Referer是檢查此請求由哪裡來，通常可以做圖片的盜鏈判斷。在Scrapy中，如果某個頁面url是通過之前爬取的頁面提取到，Scrapy會自動把之前爬取的頁面url作為Referfer。也可以通過上面的方式自己定義Referfer字段。

Cookies
=======

網站可能會檢測Cookie中session_id的使用次數，如果超過限制，就觸發反爬策略。所以可以在Scrapy中設置`COOKIES_ENABLED = False`讓請求不帶Cookies。

也有網站強制開啟Cookis，這時就要麻煩一點了。可以另寫一個簡單的爬蟲，定時向目標網站發送不帶Cookies的請求，提取響應中Set-cookie字段信息並保存。爬取網頁時，把存儲起來的Cookies帶入Headers中。

X-Forwarded-For
===============

在請求頭中添加X-Forwarded-For字段，將自己申明為一個透明的代理服務器，一些網站對代理服務器會手軟一些。

X-Forwarded-For頭一般格式如下

.. code-block:: txt

  X-Forwarded-For: client1, proxy1, proxy2

這裡將client1，proxy1設置為隨機IP地址，把自己的請求偽裝成代理的隨機IP產生的請求。然而由於X-Forwarded-For可以隨意篡改，很多網站並不會信任這個值。


限制IP的請求數量
++++++++++++++++

如果某一IP的請求速度過快，就觸發反爬機制。當然可以通過放慢爬取速度繞過，這要以爬取時間大大增長為代價。另一種方法就是添加代理。

很簡單，在下載器中間件中添加:

.. code-block:: python

  request.meta['proxy'] = 'http://' + 'proxy_host' +  ':' + proxy_port

然後再每次請求時使用不同的代理IP。然而問題是如何獲取大量的代理IP？

可以自己寫一個IP代理獲取和維護系統，定時從各種披露免費代理IP的網站爬取免費IP代理，然後定時掃瞄這些IP和端口是否可用，將不可用的代理IP及時清理。這樣就有一個動態的代理庫，每次請求再從庫中隨機選擇一個代理。然而這個方案的缺點也很明顯，開發代理獲取和維護系統本身就很費時費力，並且這種免費代理的數量並不多，而且穩定性都比較差。如果必須要用到代理，也可以去買一些穩定的代理服務。這些服務大多會用到帶認證的代理。

在requests庫中添加帶認證的代理很簡單，

.. code-block:: python

  proxies = {
      "http": "http://user:pass@10.10.1.10:3128/",
  }

然而Scrapy不支持這種認證方式，需要將認證信息base64編碼後，加入Headers的Proxy-Authorization字段：

.. code-block:: python

  import base64

  # Set the location of the proxy
  proxy_string = choice(self._get_proxies_from_file('proxies.txt')) # user:pass@ip:port
  proxy_items = proxy_string.split('@')
  request.meta['proxy'] = "http://%s" % proxy_items[1]

  # setup basic authentication for the proxy
  user_pass=base64.encodestring(proxy_items[0])
  request.headers['Proxy-Authorization'] = 'Basic ' + user_pass


動態加載
++++++++

現在越來越多的網站使用ajax動態加載內容，這時候可以先截取ajax請求分析一下，有可能根據ajax請求構造出相應的API請求的URL就可以直接獲取想要的內容，通常是json格式，反而還不用去解析HTML。

然而，很多時候ajax請求都會經過後端鑑權，不能直接構造URL獲取。這時就可以通過PhantomJS+Selenium模擬瀏覽器行為，抓取經過js渲染後的頁面。具體可以參考： `Scrapy+PhantomJS+Selenium動態爬蟲`_

需要注意的是，使用Selenium後，請求不再由Scrapy的Downloader執行，所以之前添加的請求頭等信息都會失效，需要在Selenium中重新添加

.. code-block:: python

  headers = {...}
  for key, value in headers.iteritems():
      webdriver.DesiredCapabilities.PHANTOMJS['phantomjs.page.customHeaders.{}'.format(key)] = value

另外，調用PhantomJs需要指定PhantomJs的可執行文件路徑，通常是將該路徑添加到系統的path路徑，讓程序執行時自動去path中尋找。我們的爬蟲經常會放到crontab中定時執行，而crontab中的環境變量和系統的環境變量不同，所以就加載不到PhamtonJs需要的路徑，所以最好是在申明時指定路徑：

.. code-block:: python

  driver = webdriver.PhantomJS(executable_path='/usr/local/bin/phantomjs')

----

- `普通反爬虫机制的应对策略 - Python - 伯乐在线 <http://python.jobbole.com/87669/>`_

.. _Scrapy+PhantomJS+Selenium動態爬蟲: http://jiayi.space/post/scrapy-phantomjs-seleniumdong-tai-pa-chong
