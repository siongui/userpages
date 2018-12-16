[Go語言] 本地儲存被import的套件
###############################

:date: 2018-12-01T19:25+08:00
:modified: 2018-12-12T19:18+08:00
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

----

另外有一個 Go Modules 的快速上手指南，請看 [15]_

.. adsu:: 2

----

參考：

.. [1] `Dependency management in Go : golang <https://old.reddit.com/r/golang/comments/a1ycyk/dependency_management_in_go/>`_
.. [2] `Utterly confused with Golang and its setup. : golang <https://old.reddit.com/r/golang/comments/a2b7w8/utterly_confused_with_golang_and_its_setup/>`_
.. [3] `Is this performance normal? : golang <https://old.reddit.com/r/golang/comments/a2214v/is_this_performance_normal/>`_
.. [4] `OAuth server library with Password Credentials Support (like IdentityServer for .NET) : golang <https://old.reddit.com/r/golang/comments/a24saw/oauth_server_library_with_password_credentials/>`_
.. [5] `Emulating a Nintendo DS in Go (video) : golang <https://old.reddit.com/r/golang/comments/a2hiu5/emulating_a_nintendo_ds_in_go_video/>`_
.. [6] `Tool to show type hierarchy for Go pkg : golang <https://old.reddit.com/r/golang/comments/a2csig/tool_to_show_type_hierarchy_for_go_pkg/>`_
.. [7] `Building Web Servers in Go : golang <https://old.reddit.com/r/golang/comments/a2iics/building_web_servers_in_go/>`_
.. [8] `GOPATH will not set? : golang <https://old.reddit.com/r/golang/comments/a3il4k/gopath_will_not_set/>`_
.. [9] `Reimport: a simple go files import match/replace tool using go/parser : golang <https://old.reddit.com/r/golang/comments/a35c1a/reimport_a_simple_go_files_import_matchreplace/>`_
.. [10] `Examples of well architected, large webapps : golang <https://old.reddit.com/r/golang/comments/a2siv8/examples_of_well_architected_large_webapps/>`_
.. [11] `Is MVC a good pattern in Go? : golang <https://old.reddit.com/r/golang/comments/a3lojm/is_mvc_a_good_pattern_in_go/>`_
.. [12] `Neural Net in Golang... And how : golang <https://old.reddit.com/r/golang/comments/a3t4vf/neural_net_in_golang_and_how/>`_
.. [13] `Is there a central and standard place where all the Go packages can be found? : golang <https://old.reddit.com/r/golang/comments/a44wpq/is_there_a_central_and_standard_place_where_all/>`_
.. [14] `go is pass-by-value, does that mean each function call makes a copy of all params? : golang <https://old.reddit.com/r/golang/comments/a410gl/go_is_passbyvalue_does_that_mean_each_function/>`_
.. [15] `Just tell me how to use Go Modules : golang <https://old.reddit.com/r/golang/comments/a539h6/just_tell_me_how_to_use_go_modules/>`_
.. [16] `Advent 2018: Bringing Sanity to your Dependencies with Go Modules and Athens : golang <https://old.reddit.com/r/golang/comments/a5vc16/advent_2018_bringing_sanity_to_your_dependencies/>`_
.. [17] `Go and vim in harmony – Joe Meli – Medium : golang <https://old.reddit.com/r/golang/comments/a5mf92/go_and_vim_in_harmony_joe_meli_medium/>`_
.. [18] `My first program with Go - A program that takes snapshots of subreddits and stores them in a database. : golang <https://old.reddit.com/r/golang/comments/a6hco1/my_first_program_with_go_a_program_that_takes/>`_
.. [19] `A smarter way of Iterating a map? : golang <https://old.reddit.com/r/golang/comments/a6hju8/a_smarter_way_of_iterating_a_map/>`_
.. [20] `Vue.js and Go example project : golang <https://old.reddit.com/r/golang/comments/a6pkcg/vuejs_and_go_example_project/>`_

