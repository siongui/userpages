export GOROOT=$(realpath ../../../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

URL=https://siongui.github.io/index.html

default:
	@echo "\033[92mProcessing ${URL}...\033[0m"
	@go run download.go -url=${URL}

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go
