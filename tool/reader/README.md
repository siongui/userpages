# RSS Reader

my RSS Reader written in [Go](http://golang.org/) (server side) and [Dart](https://www.dartlang.org/) (client side).

## Set up Development Environment

Tested under:

  - [Ubuntu Linux](http://www.ubuntu.com/) 13.10
  - [Go](http://golang.org/) 1.2
  - [Dart](https://www.dartlang.org/) 1.1

Download Golang and Dart SDK, after extracting the SDK, edit the path in Makefile to point to the path of your extracted SDK.

Need [sqlite](http://www.sqlite.org/) to store data on server side. Please install sqlite3 driver first:

```bash
$ make sqlite3
```

If everything goes fine, start the server by:

```bash
$ make production
```

Then open your browser at [localhost:8080](http://localhost:8080/) to use this reader.


## References

- [Golang 中国博客](http://blog.go-china.org/)
- [《Go Web 编程》](https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/preface.md)
- [【CSDN云计算俱乐部】RSS与爬虫——从如何搜集数据开始](http://www.csdn.net/article/2013-12-30/2817969-RSS-big-data)
- [Using Dart with JSON Web Services](https://www.dartlang.org/articles/json-web-service/)
- [Handling JSON Post Request in Go](http://stackoverflow.com/questions/15672556/handling-json-post-request-in-go)
- [Golang: Is it possible to capture a Ctrl+C signal and run a cleanup function, in a “defer” fashion?](http://stackoverflow.com/questions/11268943/golang-is-it-possible-to-capture-a-ctrlc-signal-and-run-a-cleanup-function-in)
- [Go 语言中实现优雅的停止程序](http://www.sudochina.com/archives/1/4/90.html)
- [A Tour of Go - Concurrency](http://tour.golang.org/#64)
- [sql - The Go Programming Language](http://golang.org/pkg/database/sql/)
- [INSERT IF NOT EXISTS ELSE UPDATE?](http://stackoverflow.com/questions/3634984/insert-if-not-exists-else-update)
- [SQLite “INSERT OR REPLACE INTO” vs. “UPDATE … WHERE”](http://stackoverflow.com/questions/2251699/sqlite-insert-or-replace-into-vs-update-where)
- [How do I check in SQLite whether a table exists?](http://stackoverflow.com/questions/1601151/how-do-i-check-in-sqlite-whether-a-table-exists)
- [mattn / go-sqlite3](https://github.com/mattn/go-sqlite3)
- [How large RSS reader works (netvibes, Google reader…)](http://stackoverflow.com/questions/3949688/how-large-rss-reader-works-netvibes-google-reader)
- [Go Read: Open-Source Google Reader Clone](http://mattjibson.com/blog/2013/06/26/go-read-open-source-google-reader-clone/)
