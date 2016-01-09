# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


test: js
	@# -v means verbose, can see logs of t.Log
	@#go test -v
	chromium-browser index.html

help:
	@#go help
	gopherjs help

js:
	gopherjs build dom.go -o dom.js

minifyjs:
	gopherjs build -m dom.go

install:
	go get -u github.com/gopherjs/gopherjs
	go get -u honnef.co/go/js/dom

clean:
	-rm -rf bin pkg src
	-rm dom.js
	-rm dom.js.map
