[Go語言] 本地儲存被import的套件
###############################

:date: 2018-12-01T19:25+08:00
:tags: Go語言
:category: Go語言
:summary: Go語言：防止網路import的套件被刪除或更改，以致無法使用。
:og_image: http://cdn.rottmann.net/wp-content/uploads/2014/02/gopher-410x300.png
:adsu: yes

Reddit看到的問題 [1]_ ，有人問到：

問： *如果有人刪除了他們的GitHub repository該怎麼辦？*

答：利用 *go mod -vendor* 可以把套件存到本地 */vendor* 目錄下。
然後在最後建構時， *go build -vendor* 利用本地 */vendor* 目錄下的套件來建構，
而不是透過網路去抓取套件建構。

問： *如果有人更新了他們在GitHub上的程式碼以致無法建構該怎麼辦？*

答： *go mod* 會固定套件版本，除非你告訴它去更改套件版本。

.. adsu:: 2

----

參考：

.. [1] `Dependency management in Go : golang <https://old.reddit.com/r/golang/comments/a1ycyk/dependency_management_in_go/>`_
.. [2] `Utterly confused with Golang and its setup. : golang <https://old.reddit.com/r/golang/comments/a2b7w8/utterly_confused_with_golang_and_its_setup/>`_
.. [3] `Is this performance normal? : golang <https://old.reddit.com/r/golang/comments/a2214v/is_this_performance_normal/>`_
.. [4] `OAuth server library with Password Credentials Support (like IdentityServer for .NET) : golang <https://old.reddit.com/r/golang/comments/a24saw/oauth_server_library_with_password_credentials/>`_

