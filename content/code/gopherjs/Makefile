# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../../../go)
export GOPATH=$(realpath ../../../../pali/go)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

DEMO_DIR=history-api

devserver: fmt js
	@echo "\033[92mDevelopment Server Running ...\033[0m"
	@go run server.go -dir=$(DEMO_DIR)/

js:
	@echo "\033[92mGenerating JavaScript ...\033[0m"
	@gopherjs build $(DEMO_DIR)/*.go -m -o $(DEMO_DIR)/app.js

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt server.go
	@go fmt $(DEMO_DIR)/*.go

install:
	go get -u github.com/gopherjs/gopherjs
