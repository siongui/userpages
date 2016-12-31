# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../../../go)
export GOPATH=$(realpath ../../../../pali/go)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

DEMO_DIR=src/arrow-keys
new_default: fmt devserver
GO_PKG_DIR=keyevent


default: kill
	gopherjs serve -w --http ":8000" &
	gopherjs build $(GO_PKG_DIR) -w -o src/$(GO_PKG_DIR)/demo.js &

kill:
	ps aux | grep gopherjs
	-killall gopherjs
	ps aux | grep gopherjs

test: js
	@# -v means verbose, can see logs of t.Log
	@#go test -v
	chromium-browser index.html

help:
	@#go help
	gopherjs help
	gopherjs serve --help
	gopherjs build --help

js:
	@#gopherjs build dom.go -o dom.js
	gopherjs build demo -w -o src/demo/dom.js

serve:
	gopherjs serve --http ":8000"

minifyjs:
	gopherjs build -m dom.go

install:
	go get -u github.com/gopherjs/gopherjs
	go get -u honnef.co/go/js/dom

clean:
	-rm -rf bin pkg
	-rm -rf src/github.com
	-rm -rf src/golang.org
	-rm -rf src/gopkg.in
	-rm -rf src/honnef.co

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt server.go
	@go fmt $(DEMO_DIR)/*.go

devserver:
	@echo "\033[92mGenerating JavaScript ...\033[0m"
	@gopherjs build $(DEMO_DIR)/*.go -m -o $(DEMO_DIR)/app.js
	@echo "\033[92mDevelopment Server Running ...\033[0m"
	@go run server.go -dir=$(DEMO_DIR)/
